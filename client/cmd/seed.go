package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	pb "github.com/vasjaj/Scoreboard/client/proto"
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
	conn, err := startConn()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer conn.Close()

	c := pb.NewScoreboardClient(conn)

	_, err = c.Seed(contextWithAuth(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("Successfully created new records")
}
