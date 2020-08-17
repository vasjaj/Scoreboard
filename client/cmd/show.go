package cmd

import (
	"fmt"
	"log"
	"strconv"
	"strings"

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
	var choice string

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

		fmt.Println("Do you want to continue viewing table? y/n")
		fmt.Scanln(&choice)

		if choice == "n" || choice == "no" {
			break
		}
	}
}

func printRes(res *pb.LeaderboardResponse) {
	printSeparator("=")
	fmt.Println("Scoreboard")
	printSeparator("=")

	scores := res.GetScore()
	fmt.Printf("Got %v scores \n", len(scores))

	for i := range scores {
		printSeparator("-")
		fmt.Printf("%v|%v| %v\n", scores[i].GetPosition(), scores[i].GetPoints(), scores[i].GetName())
	}
	printSeparator("=")

	aroundScores := res.GetAroundMe()

	if len(aroundScores) > 0 {
		printSeparator(" ")
		fmt.Println("Around me")
		printSeparator("=")
		fmt.Printf("Got %v scores \n", len(aroundScores))

		for i := range aroundScores {
			printSeparator("-")
			fmt.Printf("%v|%v| %v\n", aroundScores[i].GetPosition(), aroundScores[i].GetPoints(), aroundScores[i].GetName())
		}
	}

	fmt.Println("Next page: ", res.GetNextPage())
	printSeparator(" ")

}

func printSeparator(sep string) {
	fmt.Println(strings.Repeat(sep, 15))
}

func normalizeNumber(n int64) string {
	output := strconv.FormatInt(n, 10)

	output = output + strings.Repeat(" ", 10-len(output))
	return output
}
