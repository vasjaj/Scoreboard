package main

import (
	"log"
	"net"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func main() {
	var err error

	// setup db connection
	db, err = setupDB()
	if err != nil {
		log.Fatalln("Failed to open DB connection")
	}

	log.Println("Successfully opened DB connection")

	defer db.Close()

	server := setupServer()

	log.Println("Starting server")

	// start server
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Println("Server started")

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Error: %v", err)
	}

}
