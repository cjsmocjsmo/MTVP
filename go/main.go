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

	// Time the setup process
	start := setup.NowFunc()
	err = setup.Run()
	if err != nil {
		log.Fatalf("Setup failed: %v", err)
	}
	elapsed := setup.SinceFunc(start)
	log.Printf("Setup completed successfully. Elapsed time: %s", elapsed)

	// Start the server (WebSocket and static files)
	server.StartServer()
}
