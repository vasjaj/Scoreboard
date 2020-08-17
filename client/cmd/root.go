package cmd

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var rootCmd = &cobra.Command{}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func contextWithAuth() context.Context {
	header := metadata.New(map[string]string{"token": "correct_token"})

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
