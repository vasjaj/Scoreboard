package cmd

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/spf13/cobra"
	pb "github.com/vasjaj/Scoreboard/client/proto"
	"google.golang.org/grpc"
)

func init() {
	rootCmd.AddCommand(saveCmd)
}

var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Starts steream for score saving",
	Run:   save,
}

func save(cmd *cobra.Command, args []string) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewScoreboardClient(conn)

	stream, err := c.StoreScore(contextWithAuth())

	var name string
	var points int64
	var choice string

	for {
		fmt.Println("Enter name")
		fmt.Scanln(&name)
		fmt.Println("Entered name: ", name)

		fmt.Println("Enter points")
		fmt.Scanln(&points)
		fmt.Println("Entered points: ", points)

		req := pb.PlayerScoreRequest{
			Name:   name,
			Points: int64(points),
		}

		stream.Send(&req)

		res, err := stream.Recv()

		if err != nil {
			if err == io.EOF {
				log.Println("Stream closed")
			}

			if authenticationError(err) {
				log.Println("Authentication error")
			}

			log.Fatalf("Error: %v", err)
		}

		fmt.Println("Your rank: ", res.GetRank())

		fmt.Println("Do you want to continue saving scores? y/n")
		fmt.Scanln(&choice)

		if choice == "n" || choice == "no" {
			log.Println("Closing stream ")

			if err := stream.CloseSend(); err != nil {
				log.Fatalf("Error: %v", err)
			}

			break
		}

	}

	log.Println("Waiting 5 seconds to close stream properly")
	time.Sleep(time.Second * 5)

	log.Println("Done")
}

// func doBidrectional(c calculatorpb.CalculatorServiceClient) {
// 	requests := []*calculatorpb.FindMaximumRequest{
// 		&calculatorpb.FindMaximumRequest{Number: 21},
// 		&calculatorpb.FindMaximumRequest{Number: 23},
// 		&calculatorpb.FindMaximumRequest{Number: 3},
// 		&calculatorpb.FindMaximumRequest{Number: 44},
// 	}

// 	stream, err := c.FindMaximum(context.Background())
// 	if err != nil {
// 		log.Fatalf("error while calling Average: %v", err)
// 	}

// 	waitc := make(chan struct{})

// 	go func() {
// 		for _, req := range requests {
// 			stream.Send(req)
// 			time.Sleep(1 * time.Second)
// 		}

// 		stream.CloseSend()
// 	}()

// 	go func() {
// 		for {
// 			res, err := stream.Recv()
// 			if err == io.EOF {
// 				close(waitc)
// 				break
// 			}
// 			if err != nil {
// 				log.Fatalf("Error while receveing %v", err)
// 			}

// 			fmt.Printf("Received: %v", res)
// 		}

// 	}()

// 	<-waitc
// }
