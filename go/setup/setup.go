package setup

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
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

func databaseFileExists(dbPath string) (bool, error) {
	if dbPath == ":memory:" {
		return false, nil
	}

	path := dbPath
	if strings.HasPrefix(dbPath, "file:") {
		parsed, err := url.Parse(dbPath)
		if err != nil {
			return false, fmt.Errorf("invalid database path: %w", err)
		}
		if parsed.Opaque != "" {
			path = parsed.Opaque
		} else {
			path = parsed.Path
		}
	}
	if path == "" {
		return false, fmt.Errorf("database path does not refer to a file")
	}

	_, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("failed to inspect database path: %w", err)
	}
	return true, nil
}

func validateSchema(db *sql.DB) error {
	required := []struct {
		table   string
		columns []string
	}{
		{"movies", []string{"Name", "Year", "PosterAddr", "Size", "Path", "Idx", "MovId", "Catagory", "HttpThumbPath"}},
		{"tvshows", []string{"TvId", "Size", "Catagory", "Name", "Season", "Episode", "Path", "Idx"}},
		{"images", []string{"ImgId", "Path", "ImgPath", "Size", "Name", "ThumbPath", "Idx", "HttpThumbPath"}},
		{"videos", []string{"VidId", "VidPath", "Size", "Name", "Idx"}},
		{"weather", []string{"FetchedAt", "Location", "Temperature", "TemperatureUnit", "Conditions", "WindDirection", "WindSpeed", "Humidity", "Idx"}},
		{"nasa", []string{"Date", "Explanation", "HDURL", "MediaType", "ServiceVersion", "Title", "URL", "ThumbnailURL", "Copyright", "Idx"}},
	}

	for _, table := range required {
		rows, err := db.Query("PRAGMA table_info(" + table.table + ")")
		if err != nil {
			return fmt.Errorf("inspect %s table: %w", table.table, err)
		}
		found := make(map[string]bool, len(table.columns))
		for rows.Next() {
			var cid, notnull, primaryKey int
			var name, dataType string
			var defaultValue interface{}
			if err := rows.Scan(&cid, &name, &dataType, &notnull, &defaultValue, &primaryKey); err != nil {
				rows.Close()
				return fmt.Errorf("inspect %s table columns: %w", table.table, err)
			}
			found[name] = true
		}
		if err := rows.Err(); err != nil {
			rows.Close()
			return fmt.Errorf("inspect %s table columns: %w", table.table, err)
		}
		rows.Close()

		for _, column := range table.columns {
			if !found[column] {
				return fmt.Errorf("required table %q is missing or incomplete (missing column %q)", table.table, column)
			}
		}
	}
	return nil
}

// Run initializes a new database or validates and synchronizes an existing one.
func Run() error {
	dbPath := os.Getenv("MTVGO_DB_PATH")
	log.Println("[SETUP] Opening database at:", dbPath)
	exists, err := databaseFileExists(dbPath)
	if err != nil {
		return err
	}
	db, err := OpenDB(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()

	if !exists {
		log.Println("[SETUP] New database detected; creating tables...")
		if err := createTables(db); err != nil {
			return fmt.Errorf("failed to create tables: %w", err)
		}
	} else {
		log.Println("[SETUP] Existing database detected; validating schema...")
		if err := validateSchema(db); err != nil {
			return fmt.Errorf("existing database schema is incomplete: %w", err)
		}
	}
	if err := ensureMediaScanStateTable(db); err != nil {
		return fmt.Errorf("failed to prepare media scan state: %w", err)
	}
	log.Println("[SETUP] Database schema is ready.")

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
