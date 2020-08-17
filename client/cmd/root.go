package cmd

import (
	"context"
	"log"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	rootCmd    = &cobra.Command{}
	tokenInput string
	urlInput   string
)

func init() {
	rootCmd.PersistentFlags().StringVar(&tokenInput, "token", "correct_token", "Authentication token")
	rootCmd.PersistentFlags().StringVar(&urlInput, "url", "localhost:50051", "Service url with port")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func startConn() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(urlInput, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return conn, err
}

func contextWithAuth() context.Context {
	header := metadata.New(map[string]string{"token": tokenInput})

	return metadata.NewOutgoingContext(context.Background(), header)
}

func authenticationError(err error) bool {
	errStatus, _ := status.FromError(err)

	if codes.Unauthenticated == errStatus.Code() {
		return true
	} else {
		return false
	}
}

func invalidArgumentError(err error) bool {
	errStatus, _ := status.FromError(err)

	if codes.InvalidArgument == errStatus.Code() {
		return true
	} else {
		return false
	}
}
