package cmd

import (
	"log"

	"github.com/spf13/cobra"
	pb "github.com/vasjaj/Scoreboard/client/proto"
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
		log.Fatalf("Error: %v", err)
	}
	defer conn.Close()

	c := pb.NewScoreboardClient(conn)

	_, err = c.Seed(contextWithAuth(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Println("Successfully created new records")
}
