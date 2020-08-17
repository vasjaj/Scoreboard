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
	// endless loop for requests
	for {
		log.Println("Got StoreScore request")

		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("Stream closed")

			return nil
		}
		if err != nil {
			return err
		}

		// process request and get position
		// done in additional function for readability purposes
		position, err := processStoreScoreRequest(req)
		if err != nil {
			return err
		}

		err = stream.Send(&pb.PlayerScoreResponse{Rank: position})
		if err != nil {
			return err
		}
	}
}

func (*server) GetLeaderboard(ctx context.Context, req *pb.LeaderboardRequest) (*pb.LeaderboardResponse, error) {
	log.Println("Got GetLeaderboard request")

	// get score table bases on request params
	// scores - player scores
	// nextPage - number of next page
	// includesName - needed for around_me functionality
	scores, nextPage, includesName, err := findScoresWithPositionsByPage(req.GetName(), req.GetPage(), req.GetPageSize(), req.GetMonthly())
	if err != nil {
		return nil, err
	}

	var scoresRes []*pb.LeaderboardScore
	for i := range scores {
		scoresRes = append(scoresRes, &pb.LeaderboardScore{Name: scores[i].Name, Position: scores[i].Position, Points: scores[i].Points})
	}

	var aroundMeRes []*pb.LeaderboardScore
	// if request has name and this name was not in previous results
	if req.GetName() != "" && !includesName {
		// find scores around player
		aroundScores, err := findScoresAround(req.GetName(), req.GetMonthly())
		if err != nil {
			return nil, err
		}

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

	// find current score for later manipulations
	score, err := findScore(name)
	if err != nil {
		// record is new then save and find current position
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

	// check whether new score is more than previus, update if it is
	if score.Points < points {
		err := updateScore(name, points)

		if err != nil {
			return position, err
		}
	}

	// find current position
	score, err = findScoreWithPosition(name)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return score.Position, nil
}

func authFunc(ctx context.Context) (context.Context, error) {
	authToken := os.Getenv("AUTH_TOKEN")

	// happens if token is not provided as environmant variable
	if authToken == "" {
		return ctx, errors.New("auth token not found")
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Fatalln("Error while reading metadata")
	}

	log.Println("Incoming token:", md["token"])

	// return error if requests's token is invalid
	if md["token"][0] != authToken {
		return ctx, status.Error(codes.Unauthenticated, "invalid authentication token")
	}

	return ctx, nil
}

func setupServer() *grpc.Server {
	// interceptors for both stream and unary
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
