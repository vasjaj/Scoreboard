package main

import (
	"fmt"
	"log"
	"net"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func main() {
	var err error
	db, err = setupDB()
	if err != nil {
		log.Fatalln("Failed to open DB connection")
	}
	defer db.Close()

	server := setupServer()

	fmt.Println("Starting server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("Exit")
}
