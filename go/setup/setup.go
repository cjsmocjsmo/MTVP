
package setup

import (
	"database/sql"
	"fmt"
	"os"
	"log"
	"sync"
	"time"
	_ "github.com/mattn/go-sqlite3"
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



	   log.Println("[SETUP] Populating movies, TV shows, and videos (parallel)...")
	   var moviesErr, tvErr, videosErr error
	   var wg sync.WaitGroup
	   wg.Add(3)
	   go func() {
		   defer wg.Done()
		   if err := populateMovies(db); err != nil {
			   moviesErr = fmt.Errorf("failed to populate movies: %w", err)
		   }
	   }()
	   go func() {
		   defer wg.Done()
		   if err := populateTVShows(db); err != nil {
			   tvErr = fmt.Errorf("failed to populate tvshows: %w", err)
		   }
	   }()
	   go func() {
		   defer wg.Done()
		   if err := populateVideos(db); err != nil {
			   videosErr = fmt.Errorf("failed to populate videos: %w", err)
		   }
	   }()
	   wg.Wait()
	   if moviesErr != nil {
		   return moviesErr
	   }
	   log.Println("[SETUP] Movies populated.")
	   if tvErr != nil {
		   return tvErr
	   }
	   log.Println("[SETUP] TV shows populated.")
	   if videosErr != nil {
		   return videosErr
	   }
	   log.Println("[SETUP] Videos populated.")

	log.Println("[SETUP] Database setup complete.")
	return nil
}
