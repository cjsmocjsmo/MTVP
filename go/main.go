package main


import (
	"log"
	"mtvp/setup"
	"mtvp/server"
)

func main() {
	// Load environment variables
	err := setup.LoadEnv()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Run setup to populate the database
	err = setup.Run()
	if err != nil {
		log.Fatalf("Setup failed: %v", err)
	}

	log.Println("Setup completed successfully.")

	// Start the server (WebSocket and static files)
	server.StartServer()
}
