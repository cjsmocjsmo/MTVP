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
	dbPath := os.Getenv("MTV_DB_PATH")
	if dbPath == "" {
		return fmt.Errorf("MTV_DB_PATH not set in environment")
	}
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	defer db.Close()

	if err := createTables(db); err != nil {
		return fmt.Errorf("failed to create tables: %w", err)
	}

	if err := populateMovies(db); err != nil {
		return fmt.Errorf("failed to populate movies: %w", err)
	}

	if err := populateTVShows(db); err != nil {
		return fmt.Errorf("failed to populate tvshows: %w", err)
	}

	if err := populateVideos(db); err != nil {
		return fmt.Errorf("failed to populate videos: %w", err)
	}

	log.Println("Database setup complete.")
	return nil
}
