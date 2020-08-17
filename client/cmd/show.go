package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	pb "github.com/vasjaj/Scoreboard/client/proto"
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
	conn, err := startConn()
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewScoreboardClient(conn)

	var name string
	var pageSize int64
	var page int64
	var monthlyStr string
	var monthly bool
	var defPageSize int64 = 5

	for {
		fmt.Println("Enter name (can be empty)")
		fmt.Scanln(&name)

		fmt.Printf("Enter page size (default: %v) \n", defPageSize)
		fmt.Scanln(&pageSize)
		if pageSize == 0 {
			pageSize = defPageSize
		} else {
			defPageSize = pageSize
		}

		fmt.Println("Enter page")
		fmt.Scanln(&page)

		fmt.Println("Monthly? y/n")
		fmt.Scanln(&monthlyStr)

		if monthlyStr == "y" || monthlyStr == "yes" {
			monthly = true
		} else {
			monthly = false
		}

		req := pb.LeaderboardRequest{
			Name:     name,
			PageSize: pageSize,
			Page:     page,
			Monthly:  monthly,
		}

		res, err := c.GetLeaderboard(contextWithAuth(), &req)
		if err != nil {
			if authenticationError(err) {
				fmt.Println("Authentication error")
			}

			if invalidArgumentError(err) {
				fmt.Println("Invalid page number")
			}

			log.Fatalf("Error: %v", err)
		}

		printRes(res)
	}
}

func printRes(res *pb.LeaderboardResponse) {

	printSeparator()
	fmt.Println("Scoreboard")
	printSeparator()

	scores := res.GetScore()
	fmt.Printf("Got %v scores \n", len(scores))

	for i := range scores {
		fmt.Println("--------------------------------------------------------")
		fmt.Printf("%v has %v points and takes %v place \n", scores[i].GetName(), scores[i].GetPoints(), scores[i].GetPosition())
	}
	printSeparator()

	aroundScores := res.GetAroundMe()

	if len(aroundScores) > 0 {
		fmt.Println("Around me")
		printSeparator()
		fmt.Printf("Got %v scores \n", len(aroundScores))

		for i := range aroundScores {
			fmt.Println("--------------------------------------------------------")
			fmt.Printf("%v has %v points and takes %v place \n", aroundScores[i].GetName(), aroundScores[i].GetPoints(), aroundScores[i].GetPosition())
		}
	}

	fmt.Println("Next page: ", res.GetNextPage())
}

func printSeparator() {
	fmt.Println("========================================================")
}
