package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	pb "github.com/vasjaj/Scoreboard/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func init() {
	rootCmd.AddCommand(seedCmd)
}

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Creates new records",
	Run:   seed,
}

func seed(cmd *cobra.Command, args []string) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewScoreboardClient(conn)

	_, err = c.Seed(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Println("Successfully created new records")
}
