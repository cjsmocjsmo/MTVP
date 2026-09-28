package main

import (
	"log"
	"mtvp/server"
	"mtvp/setup"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load("../env/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Time database initialization and media synchronization.
	start := setup.NowFunc()
	err = setup.Run()
	if err != nil {
		log.Fatalf("Startup initialization failed: %v", err)
	}
	elapsed := setup.SinceFunc(start)
	log.Printf("Startup initialization completed successfully. Elapsed time: %s", elapsed)

	// Start the server (WebSocket and static files)
	server.StartServer()
}
