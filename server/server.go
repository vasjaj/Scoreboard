package main

import (
	"context"
	"io"
	"log"
	"os"

	"errors"

	"github.com/golang/protobuf/ptypes/empty"
	auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/jinzhu/gorm"
	pb "github.com/vasjaj/Scoreboard/server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct{}

func (*server) StoreScore(stream pb.Scoreboard_StoreScoreServer) error {
	for {
		log.Println("Got StoreScore request")

		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("Stream closed")

			return nil
		}
		if err != nil {
			log.Println("Error while stoping stream")
			log.Fatalf("Error: %v", err)
		}

		position, err := processStoreScoreRequest(req)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		err = stream.Send(&pb.PlayerScoreResponse{Rank: position})

		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}
}

func (*server) GetLeaderboard(ctx context.Context, req *pb.LeaderboardRequest) (*pb.LeaderboardResponse, error) {
	log.Println("Got GetLeaderboard request")
	log.Println("Name: ", req.GetName())
	log.Println("Page: ", req.GetPage())
	log.Println("Page size: ", req.GetPageSize())
	log.Println("Monthly: ", req.GetMonthly())

	scores, nextPage, includesName, err := findScoresWithPositionsByPage(req.GetName(), req.GetPage(), req.GetPageSize(), req.GetMonthly())
	if err != nil {
		return nil, err
	}

	log.Printf("Found %v scores \n", len(scores))

	var scoresRes []*pb.LeaderboardScore
	for i := range scores {
		scoresRes = append(scoresRes, &pb.LeaderboardScore{Name: scores[i].Name, Position: scores[i].Position, Points: scores[i].Points})
	}

	var aroundMeRes []*pb.LeaderboardScore
	if req.GetName() != "" && !includesName {
		log.Println("Searching for players around")

		aroundScores, err := findScoresAround(req.GetName(), req.GetMonthly())
		if err != nil {
			return nil, err
		}

		log.Printf("Found %v scores around \n", len(aroundScores))

		for i := range aroundScores {
			aroundMeRes = append(aroundMeRes, &pb.LeaderboardScore{Name: aroundScores[i].Name, Position: aroundScores[i].Position, Points: aroundScores[i].Points})
		}
	}

	return &pb.LeaderboardResponse{Score: scoresRes, NextPage: nextPage, AroundMe: aroundMeRes}, nil
}

func (*server) Seed(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
	seed()

	return &emptypb.Empty{}, nil
}

func processStoreScoreRequest(req *pb.PlayerScoreRequest) (int64, error) {
	var position int64
	name := req.GetName()
	points := req.GetPoints()

	score, err := findScore(name)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			saveScore(name, points)

			score, err = findScoreWithPosition(name)
			if err != nil {
				log.Fatalf("Error: %v", err)
			}

			return score.Position, nil
		} else {
			return position, err
		}
	}

	if score.Points < points {
		err := updateScore(name, points)

		if err != nil {
			return position, err
		}
	}

	score, err = findScoreWithPosition(name)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return score.Position, nil
}

func authFunc(ctx context.Context) (context.Context, error) {
	authToken := os.Getenv("AUTH_TOKEN")

	if authToken == "" {
		return ctx, errors.New("auth token not found")
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Fatalln("Error while reading metadata")
	}

	log.Println("Incoming token:", md["token"])

	if md["token"][0] != authToken {
		return ctx, status.Error(codes.Unauthenticated, "invalid authentication token")
	}

	return ctx, nil
}

func setupServer() *grpc.Server {

	s := grpc.NewServer(
		grpc.StreamInterceptor(
			auth.StreamServerInterceptor(authFunc),
		),
		grpc.UnaryInterceptor(
			auth.UnaryServerInterceptor(authFunc),
		),
	)

	pb.RegisterScoreboardServer(s, &server{})
	reflection.Register(s)

	return s
}
