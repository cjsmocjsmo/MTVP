package setup

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"strings"
	"time"
)

// NowFunc and SinceFunc allow main.go to time setup without importing time directly
var NowFunc = func() interface{} {
	return time.Now()
}
var SinceFunc = func(start interface{}) string {
	if t, ok := start.(time.Time); ok {
		return time.Since(t).String()
	}
	return "unknown"
}

func OpenDB(dbPath string) (*sql.DB, error) {
	if dbPath == "" {
		return nil, fmt.Errorf("MTVGO_DB_PATH not set in environment")
	}

	dsn := dbPath
	if !strings.HasPrefix(dsn, "file:") {
		dsn = "file:" + dsn
	}
	dsn += "?_busy_timeout=5000&_journal_mode=WAL"

	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(0)

	if _, err := db.Exec(`PRAGMA busy_timeout = 5000`); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to set sqlite busy timeout: %w", err)
	}

	return db, nil
}

// Run executes the setup process: creates tables and populates the DB
func Run() error {
	dbPath := os.Getenv("MTVGO_DB_PATH")
	log.Println("[SETUP] Opening database at:", dbPath)
	db, err := OpenDB(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()

	log.Println("[SETUP] Creating tables...")
	if err := createTables(db); err != nil {
		return fmt.Errorf("failed to create tables: %w", err)
	}
	log.Println("[SETUP] Tables created.")

	log.Println("[SETUP] Populating movies, TV shows, and videos (serial)...")
	if err := populateMovies(db); err != nil {
		return fmt.Errorf("failed to populate movies: %w", err)
	}
	log.Println("[SETUP] Movies populated.")
	if err := populateTVShows(db); err != nil {
		return fmt.Errorf("failed to populate tvshows: %w", err)
	}
	log.Println("[SETUP] TV shows populated.")
	if err := populateVideos(db); err != nil {
		return fmt.Errorf("failed to populate videos: %w", err)
	}
	log.Println("[SETUP] Videos populated.")

	log.Println("[SETUP] Database setup complete.")
	return nil
}
