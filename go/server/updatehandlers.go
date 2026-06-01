package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"mtvp/setup"
)

func UpdateHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		newmovfiles, err := movScanCompare(db)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to scan movies: %v", err), http.StatusInternalServerError)
			return
		}
		newtvfiles, err := tvshowScanCompare(db)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to scan tvshows: %v", err), http.StatusInternalServerError)
			return
		}
		newvideofiles, err := videoScanCompare(db)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to scan videos: %v", err), http.StatusInternalServerError)
			return
		}

		var moviePaths []string
		if err := json.Unmarshal([]byte(newmovfiles), &moviePaths); err != nil {
			http.Error(w, fmt.Sprintf("failed to decode movie paths: %v", err), http.StatusInternalServerError)
			return
		}

		var tvshowPaths []string
		if err := json.Unmarshal([]byte(newtvfiles), &tvshowPaths); err != nil {
			http.Error(w, fmt.Sprintf("failed to decode tvshow paths: %v", err), http.StatusInternalServerError)
			return
		}

		var videoPaths []string
		if err := json.Unmarshal([]byte(newvideofiles), &videoPaths); err != nil {
			http.Error(w, fmt.Sprintf("failed to decode video paths: %v", err), http.StatusInternalServerError)
			return
		}

		movieInserted := 0
		for _, path := range moviePaths {
			if err := movUpdateInsert(db, path); err != nil {
				http.Error(w, fmt.Sprintf("failed to insert movie %s: %v", path, err), http.StatusInternalServerError)
				return
			}
			movieInserted++
		}

		tvshowInserted := 0
		for _, path := range tvshowPaths {
			if err := tvshowUpdateInsert(db, path); err != nil {
				http.Error(w, fmt.Sprintf("failed to insert tvshow %s: %v", path, err), http.StatusInternalServerError)
				return
			}
			tvshowInserted++
		}

		videoInserted := 0
		for _, path := range videoPaths {
			if err := videoUpdateInsert(db, path); err != nil {
				http.Error(w, fmt.Sprintf("failed to insert video %s: %v", path, err), http.StatusInternalServerError)
				return
			}
			videoInserted++
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]int{
			"moviesInserted":  movieInserted,
			"tvshowsInserted": tvshowInserted,
			"videosInserted":  videoInserted,
		})

	}
}

func movUpdateInsert(db *sql.DB, path string) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("failed to stat movie %s: %w", path, err)
	}

	filename := filepath.Base(path)
	movId := setup.GenerateMD5(path)
	name := setup.GetMovieName(filename)
	year := setup.GetMovieYear(filename)
	poster := strings.TrimSuffix(path, filepath.Ext(path)) + ".jpg"
	catagory := setup.GetMovieCategory(path)
	thumbPath := setup.GetHttpThumbPath(filename)

	var nextIdx int
	if err := db.QueryRow(`SELECT COALESCE(MAX(Idx), 0) + 1 FROM movies`).Scan(&nextIdx); err != nil {
		return fmt.Errorf("failed to determine next movie index for %s: %w", path, err)
	}

	_, err = db.Exec(`INSERT OR IGNORE INTO movies (Name, Year, PosterAddr, Size, Path, Idx, MovId, Catagory, HttpThumbPath) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		name, year, poster, fileInfo.Size(), path, nextIdx, movId, catagory, thumbPath)
	if err != nil {
		return fmt.Errorf("failed to insert movie %s: %w", path, err)
	}
	return nil
}

func tvshowUpdateInsert(db *sql.DB, path string) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("failed to stat tvshow %s: %w", path, err)
	}

	filename := filepath.Base(path)
	tvId := setup.GenerateTVId(path)
	name := setup.GetTVShowName(filename)
	season := setup.GetTVShowSeason(filename)
	episode := setup.GetTVShowEpisode(filename)
	category := setup.GetTVShowCategory(path)

	var nextIdx int
	if err := db.QueryRow(`SELECT COALESCE(MAX(Idx), 0) + 1 FROM tvshows`).Scan(&nextIdx); err != nil {
		return fmt.Errorf("failed to determine next tvshow index for %s: %w", path, err)
	}

	_, err = db.Exec(`INSERT OR IGNORE INTO tvshows (TvId, Size, Catagory, Name, Season, Episode, Path, Idx) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		tvId, fileInfo.Size(), category, name, season, episode, path, nextIdx)
	if err != nil {
		return fmt.Errorf("failed to insert tvshow %s: %w", path, err)
	}
	return nil
}

func videoUpdateInsert(db *sql.DB, path string) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("failed to stat video %s: %w", path, err)
	}

	vidId, err := setup.GenerateVidId(path)
	if err != nil {
		return fmt.Errorf("failed to generate video id for %s: %w", path, err)
	}
	name := filepath.Base(path)

	var nextIdx int
	if err := db.QueryRow(`SELECT COALESCE(MAX(Idx), 0) + 1 FROM videos`).Scan(&nextIdx); err != nil {
		return fmt.Errorf("failed to determine next video index for %s: %w", path, err)
	}

	_, err = db.Exec(`INSERT OR IGNORE INTO videos (VidId, VidPath, Size, Name, Idx) VALUES (?, ?, ?, ?, ?)`,
		vidId, path, fileInfo.Size(), name, nextIdx)
	if err != nil {
		return fmt.Errorf("failed to insert video %s: %w", path, err)
	}
	return nil
}

func movScanCompare(db *sql.DB) (string, error) {
	mov_dir_path := os.Getenv("MTVGO_MOVIES_PATH")
	if mov_dir_path == "" {
		return "", fmt.Errorf("MTVGO_MOVIES_PATH not set")
	}

	// 1. Walk the directory and collect all file paths
	fsPaths := []string{}
	err := filepath.WalkDir(mov_dir_path, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			fsPaths = append(fsPaths, path)
		}
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("error walking movie directory: %w", err)
	}

	// 2. Get all movie paths from the database
	rows, err := db.Query("SELECT Path FROM movies")
	if err != nil {
		return "", fmt.Errorf("database query error: %w", err)
	}
	defer rows.Close()
	dbPaths := map[string]struct{}{}
	for rows.Next() {
		var dbPath string
		if err := rows.Scan(&dbPath); err == nil {
			dbPaths[dbPath] = struct{}{}
		}
	}
	if err := rows.Err(); err != nil {
		return "", fmt.Errorf("database rows error: %w", err)
	}

	// 3. Compare and collect new paths
	newPaths := []string{}
	for _, path := range fsPaths {
		if _, exists := dbPaths[path]; !exists {
			newPaths = append(newPaths, path)
		}
	}

	// 4. Return new paths as a JSON array string
	result, err := json.Marshal(newPaths)
	if err != nil {
		return "", fmt.Errorf("json marshal error: %w", err)
	}
	return string(result), nil
}

func tvshowScanCompare(db *sql.DB) (string, error) {
	tv_dir_path := os.Getenv("MTVGO_TV_PATH")
	if tv_dir_path == "" {
		return "", fmt.Errorf("MTVGO_TV_PATH not set")
	}

	// 1. Walk the directory and collect all file paths
	fsPaths := []string{}
	err := filepath.WalkDir(tv_dir_path, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			fsPaths = append(fsPaths, path)
		}
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("error walking TV show directory: %w", err)
	}

	// 2. Get all TV show paths from the database
	rows, err := db.Query("SELECT Path FROM tvshows")
	if err != nil {
		return "", fmt.Errorf("database query error: %w", err)
	}
	defer rows.Close()
	dbPaths := map[string]struct{}{}
	for rows.Next() {
		var dbPath string
		if err := rows.Scan(&dbPath); err == nil {
			dbPaths[dbPath] = struct{}{}
		}
	}
	if err := rows.Err(); err != nil {
		return "", fmt.Errorf("database rows error: %w", err)
	}

	// 3. Compare and collect new paths
	newPaths := []string{}
	for _, path := range fsPaths {
		if _, exists := dbPaths[path]; !exists {
			newPaths = append(newPaths, path)
		}
	}

	// 4. Return new paths as a JSON array string
	result, err := json.Marshal(newPaths)
	if err != nil {
		return "", fmt.Errorf("json marshal error: %w", err)
	}
	return string(result), nil
}

func videoScanCompare(db *sql.DB) (string, error) {
	video_dir_path := os.Getenv("MTVGO_VIDEOS_PATH")
	if video_dir_path == "" {
		return "", fmt.Errorf("MTVGO_VIDEOS_PATH not set")
	}

	// 1. Walk the directory and collect all file paths
	fsPaths := []string{}
	err := filepath.WalkDir(video_dir_path, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			fsPaths = append(fsPaths, path)
		}
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("error walking video directory: %w", err)
	}

	// 2. Get all video paths from the database
	rows, err := db.Query("SELECT VidPath FROM videos")
	if err != nil {
		return "", fmt.Errorf("database query error: %w", err)
	}
	defer rows.Close()
	dbPaths := map[string]struct{}{}
	for rows.Next() {
		var dbPath string
		if err := rows.Scan(&dbPath); err == nil {
			dbPaths[dbPath] = struct{}{}
		}
	}
	if err := rows.Err(); err != nil {
		return "", fmt.Errorf("database rows error: %w", err)
	}

	// 3. Compare and collect new paths
	newPaths := []string{}
	for _, path := range fsPaths {
		if _, exists := dbPaths[path]; !exists {
			newPaths = append(newPaths, path)
		}
	}

	// 4. Return new paths as a JSON array string
	result, err := json.Marshal(newPaths)
	if err != nil {
		return "", fmt.Errorf("json marshal error: %w", err)
	}
	return string(result), nil
}
