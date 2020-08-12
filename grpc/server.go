package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"errors"

	auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/jinzhu/gorm"
	"github.com/vasjaj/Scoreboard/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (*server) StoreScore(stream pb.Scoreboard_StoreScoreServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		rank, err := processStoreScoreRequest(req)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		err = stream.Send(&pb.PlayerScoreResponse{Rank: rank})

		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}
}

func processStoreScoreRequest(req *pb.PlayerScoreRequest) (int64, error) {
	var rank int64 = 0
	message := req.GetScore()
	name := message.GetName()
	points := message.GetPoints()

	score, err := findScore(name)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			saveScore(name, points)
		} else {
			return rank, err
		}
	}

	if score.Points < points {
		err := updateScore(score, points)

		if err != nil {
			return rank, err
		}
	}

	getRanks(name)

	return rank, nil
}

func (*server) GetLeaderboard(ctx context.Context, req *pb.LeaderboardRequest) (*pb.LeaderboardResponse, error) {
	return nil, nil
}

func authFunc(ctx context.Context) (context.Context, error) {
	// fmt.Println("Authenticating")
	// token, err := auth.AuthFromMD(ctx, "basic")
	// fmt.Println("Token:", token)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, err
	// }
	// authToken := os.Getenv("AUTH_TOKEN")
	authToken := "correct_token"

	if authToken == "" {
		return ctx, errors.New("auth token not found")
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Fatalln("Error while reading metadata")
	}

	if md["token"][0] != authToken {
		return ctx, errors.New("incorrect auth token")
	}
	fmt.Printf("Token: %v \n", md["token"])
	fmt.Printf("Metadata: %v \n", md)

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
	// s := grpc.NewServer()
	pb.RegisterScoreboardServer(s, &server{})
	reflection.Register(s)

	return s
}

// func setupServer() (*grpc.Server, error) {
// 	lis, err := net.Listen("tcp", "0.0.0.0:50051")
// 	if err != nil {
// 		log.Fatalf("Failed to listen: %v", err)
// 	}

// 	s := grpc.NewServer()
// 	pb.RegisterScoreboardServer(s, &server{})
// 	reflection.Register(s)

// 	if err := s.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }
