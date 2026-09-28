package setup

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// Dummy data population for demonstration. Replace with real directory walking and parsing logic.

func ensureMediaScanStateTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS media_scan_state (
		Source TEXT NOT NULL,
		Path TEXT NOT NULL,
		Size INTEGER NOT NULL,
		ModTime INTEGER NOT NULL,
		PRIMARY KEY (Source, Path)
	)`)
	return err
}

func pathsNotInDB(db *sql.DB, table, pathColumn string, paths []string) ([]string, int, error) {
	rows, err := db.Query(fmt.Sprintf("SELECT %s FROM %s", pathColumn, table))
	if err != nil {
		return nil, 0, err
	}
	existing := make(map[string]bool)
	for rows.Next() {
		var path string
		if err := rows.Scan(&path); err != nil {
			rows.Close()
			return nil, 0, err
		}
		existing[path] = true
	}
	if err := rows.Err(); err != nil {
		rows.Close()
		return nil, 0, err
	}
	rows.Close()

	var maxIdx int
	if err := db.QueryRow(fmt.Sprintf("SELECT COALESCE(MAX(Idx), 0) FROM %s", table)).Scan(&maxIdx); err != nil {
		return nil, 0, err
	}
	newPaths := make([]string, 0, len(paths))
	for _, path := range paths {
		if !existing[path] {
			newPaths = append(newPaths, path)
		}
	}
	return newPaths, maxIdx, nil
}

type mediaFileState struct {
	size     int64
	modified int64
}

func loadMediaFileStates(db *sql.DB, source string) (map[string]mediaFileState, error) {
	rows, err := db.Query(`SELECT Path, Size, ModTime FROM media_scan_state WHERE Source = ?`, source)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	states := make(map[string]mediaFileState)
	for rows.Next() {
		var path string
		var state mediaFileState
		if err := rows.Scan(&path, &state.size, &state.modified); err != nil {
			return nil, err
		}
		states[path] = state
	}
	return states, rows.Err()
}

func saveMediaFileStates(db *sql.DB, source string, states map[string]mediaFileState) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`INSERT INTO media_scan_state (Source, Path, Size, ModTime)
		VALUES (?, ?, ?, ?)
		ON CONFLICT(Source, Path) DO UPDATE SET Size = excluded.Size, ModTime = excluded.ModTime`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	for path, state := range states {
		if _, err := stmt.Exec(source, path, state.size, state.modified); err != nil {
			return err
		}
	}
	return tx.Commit()
}

func populateMovies(db *sql.DB) error {
	movieDir := os.Getenv("MTVGO_MOVIES_PATH")
	log.Println("[SETUP] [Movies] Scanning movie directory:", movieDir)
	if movieDir == "" {
		log.Println("[SETUP] [Movies] MTVGO_MOVIES_PATH not set, skipping movie population.")
	} else {
		exts := []string{".mp4", ".mkv", ".avi", ".mpg"}
		moviePaths, err := WalkMediaDirs(movieDir, exts)
		if err != nil {
			log.Println("[SETUP] [Movies] Error walking movie dirs:", err)
			return err
		}
		newMoviePaths, idxStart, err := pathsNotInDB(db, "movies", "Path", moviePaths)
		if err != nil {
			return err
		}
		log.Printf("[SETUP] [Movies] Found %d movie files; adding %d new files.", len(moviePaths), len(newMoviePaths))
		if err := InsertMovies(db, newMoviePaths, idxStart); err != nil {
			log.Println("[SETUP] [Movies] Error inserting movies:", err)
			return err
		}
	}

	// Images for movies
	imgDir := os.Getenv("MTVGO_POSTER_PATH")
	thumbDir := os.Getenv("MTVGO_THUMBNAIL_PATH")
	serverAddr := os.Getenv("MTVGO_SERVER_ADDR")
	log.Println("[SETUP] [Movies] Scanning poster directory:", imgDir)
	if imgDir != "" && thumbDir != "" && serverAddr != "" {
		imgExts := []string{".jpg"}
		imgPaths, err := WalkMediaDirs(imgDir, imgExts)
		if err == nil {
			idxStart, err := maxIndex(db, "images")
			if err != nil {
				return err
			}
			log.Printf("[SETUP] [Movies] Found %d poster images. Synchronizing...", len(imgPaths))
			if err := InsertImages(db, imgPaths, idxStart, thumbDir, serverAddr); err != nil {
				log.Println("[SETUP] [Movies] Error inserting images:", err)
				return err
			}
		} else {
			log.Println("[SETUP] [Movies] Error walking poster dirs:", err)
			return err
		}
	}
	return nil
}

func populateTVShows(db *sql.DB) error {
	tvDir := os.Getenv("MTVGO_TV_PATH")
	log.Println("[SETUP] [TVShows] Scanning TV shows directory:", tvDir)
	if tvDir == "" {
		log.Println("[SETUP] [TVShows] MTVGO_TV_PATH not set, skipping TV show population.")
	} else {
		exts := []string{".mp4", ".mkv", ".avi"}
		tvPaths, err := WalkMediaDirs(tvDir, exts)
		if err != nil {
			log.Println("[SETUP] [TVShows] Error walking TV show dirs:", err)
			return err
		}
		newTVPaths, idxStart, err := pathsNotInDB(db, "tvshows", "Path", tvPaths)
		if err != nil {
			return err
		}
		log.Printf("[SETUP] [TVShows] Found %d TV show files; adding %d new files.", len(tvPaths), len(newTVPaths))
		if err := InsertTVShows(db, newTVPaths, idxStart); err != nil {
			log.Println("[SETUP] [TVShows] Error inserting TV shows:", err)
			return err
		}
	}

	// TV show images
	imgDir := os.Getenv("MTVGO_TV_POSTER_PATH")
	thumbDir := os.Getenv("MTVGO_TV_THUMBNAIL_PATH")
	serverAddr := os.Getenv("MTVGO_SERVER_ADDR")
	log.Println("[SETUP] [TVShows] Scanning TV poster directory:", imgDir)
	if imgDir != "" && thumbDir != "" && serverAddr != "" {
		imgExts := []string{".jpg"}
		imgPaths, err := WalkMediaDirs(imgDir, imgExts)
		if err == nil {
			idxStart, err := maxIndex(db, "images")
			if err != nil {
				return err
			}
			log.Printf("[SETUP] [TVShows] Found %d TV poster images. Synchronizing...", len(imgPaths))
			if err := InsertImages(db, imgPaths, idxStart, thumbDir, serverAddr); err != nil {
				log.Println("[SETUP] [TVShows] Error inserting images:", err)
				return err
			}
		} else {
			log.Println("[SETUP] [TVShows] Error walking TV poster dirs:", err)
			return err
		}
	}
	return nil
}

func populateVideos(db *sql.DB) error {
	vidDir := os.Getenv("MTVGO_VIDEOS_PATH")
	log.Println("[SETUP] [Videos] Scanning videos directory:", vidDir)
	if vidDir == "" {
		log.Println("[SETUP] [Videos] MTVGO_VIDEOS_PATH not set, skipping video population.")
		return nil // or error if required
	}
	exts := []string{".mp4", ".mkv", ".avi"}
	vidPaths, err := WalkMediaDirs(vidDir, exts)
	if err != nil {
		log.Println("[SETUP] [Videos] Error walking video dirs:", err)
		return err
	}
	newVideoPaths, states, idxStart, err := videosNeedingSync(db, vidPaths)
	if err != nil {
		return err
	}
	log.Printf("[SETUP] [Videos] Found %d video files; hashing %d new or changed files.", len(vidPaths), len(newVideoPaths))
	if err := InsertVideos(db, newVideoPaths, idxStart); err != nil {
		return err
	}
	return saveMediaFileStates(db, "videos", states)
}

func maxIndex(db *sql.DB, table string) (int, error) {
	var idx int
	err := db.QueryRow(fmt.Sprintf("SELECT COALESCE(MAX(Idx), 0) FROM %s", table)).Scan(&idx)
	return idx, err
}

func videosNeedingSync(db *sql.DB, paths []string) ([]string, map[string]mediaFileState, int, error) {
	existingSizes := make(map[string]int64)
	rows, err := db.Query(`SELECT VidPath, Size FROM videos`)
	if err != nil {
		return nil, nil, 0, err
	}
	for rows.Next() {
		var path string
		var size int64
		if err := rows.Scan(&path, &size); err != nil {
			rows.Close()
			return nil, nil, 0, err
		}
		existingSizes[path] = size
	}
	if err := rows.Err(); err != nil {
		rows.Close()
		return nil, nil, 0, err
	}
	rows.Close()

	states, err := loadMediaFileStates(db, "videos")
	if err != nil {
		return nil, nil, 0, err
	}
	idxStart, err := maxIndex(db, "videos")
	if err != nil {
		return nil, nil, 0, err
	}
	currentStates := make(map[string]mediaFileState, len(paths))
	changedPaths := make([]string, 0, len(paths))
	for _, path := range paths {
		info, err := os.Stat(path)
		if err != nil {
			return nil, nil, 0, err
		}
		current := mediaFileState{size: info.Size(), modified: info.ModTime().UnixNano()}
		currentStates[path] = current

		previous, hasState := states[path]
		if hasState && previous == current {
			continue
		}
		if existingSize, exists := existingSizes[path]; !hasState && exists && existingSize == current.size {
			continue
		}
		changedPaths = append(changedPaths, path)
	}
	return changedPaths, currentStates, idxStart, nil
}
