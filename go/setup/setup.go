package setup

import (
	"database/sql"
	"fmt"
	"os"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

// Run executes the setup process: creates tables and populates the DB
func Run() error {
	dbPath := os.Getenv("MTVGO_DB_PATH")
	if dbPath == "" {
	    return fmt.Errorf("MTVGO_DB_PATH not set in environment")
	}
	log.Println("[SETUP] Opening database at:", dbPath)
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
	    return fmt.Errorf("failed to open database: %w", err)
	}
	defer db.Close()

	log.Println("[SETUP] Creating tables...")
	if err := createTables(db); err != nil {
	    return fmt.Errorf("failed to create tables: %w", err)
	}
	log.Println("[SETUP] Tables created.")

	log.Println("[SETUP] Populating movies...")
	if err := populateMovies(db); err != nil {
	    return fmt.Errorf("failed to populate movies: %w", err)
	}
	log.Println("[SETUP] Movies populated.")

	log.Println("[SETUP] Populating TV shows...")
	if err := populateTVShows(db); err != nil {
	    return fmt.Errorf("failed to populate tvshows: %w", err)
	}
	log.Println("[SETUP] TV shows populated.")

	log.Println("[SETUP] Populating videos...")
	if err := populateVideos(db); err != nil {
	    return fmt.Errorf("failed to populate videos: %w", err)
	}
	log.Println("[SETUP] Videos populated.")

	log.Println("[SETUP] Database setup complete.")
	return nil
}
