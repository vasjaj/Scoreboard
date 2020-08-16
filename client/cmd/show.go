package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	pb "github.com/vasjaj/Scoreboard/proto"
	"google.golang.org/grpc"
)

func init() {
	rootCmd.AddCommand(showCmd)
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Shows table results",
	Run:   show,
}

func show(cmd *cobra.Command, args []string) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewScoreboardClient(conn)

	req := pb.LeaderboardRequest{
		Name:     "John",
		PageSize: 3,
		Page:     1,
		Period:   1,
	}

	res, err := c.GetLeaderboard(context.Background(), &req)
	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Printf("Response %+v \n", res.GetScore())
	// doUnary(c)
	// doServerStreaming(c)
	// doClientStreaming(c)
	// doBidrectional(c)
}
