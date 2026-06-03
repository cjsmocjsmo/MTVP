package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"mtvp/setup"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

var (
	movieExts = map[string]struct{}{".mp4": {}, ".mkv": {}, ".avi": {}, ".mpg": {}}
	tvExts    = map[string]struct{}{".mp4": {}, ".mkv": {}, ".avi": {}}
	videoExts = map[string]struct{}{".mp4": {}, ".mkv": {}, ".avi": {}}
	updateSem = make(chan struct{}, 1)
)

const maxErrorSamples = 10

func writeJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func writeJSONError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]interface{}{
		"status":  "error",
		"message": message,
	})
}

func authorizeUpdateRequest(r *http.Request) (bool, string) {
	requiredToken := strings.TrimSpace(os.Getenv("MTVGO_UPDATE_TOKEN"))
	if requiredToken == "" {
		return true, ""
	}

	requestToken := strings.TrimSpace(r.Header.Get("X-Update-Token"))
	if requestToken == "" {
		requestToken = strings.TrimSpace(r.URL.Query().Get("token"))
	}

	if requestToken == "" {
		return false, "missing update token; provide X-Update-Token header or token query parameter"
	}

	if requestToken != requiredToken {
		return false, "invalid update token"
	}

	return true, ""
}

type updateBatchResult struct {
	Scanned  int      `json:"scanned"`
	Inserted int      `json:"inserted"`
	Skipped  int      `json:"skipped"`
	Failed   int      `json:"failed"`
	Errors   []string `json:"errors,omitempty"`
}

type insertFn func(tx *sql.Tx, path string) (bool, error)

func appendErrorSample(errors []string, message string) []string {
	if len(errors) >= maxErrorSamples {
		return errors
	}
	return append(errors, message)
}

func processUpdatesInTx(db *sql.DB, paths []string, fn insertFn) (updateBatchResult, error) {
	result := updateBatchResult{Scanned: len(paths)}
	tx, err := db.Begin()
	if err != nil {
		return result, fmt.Errorf("failed to start transaction: %w", err)
	}

	for _, path := range paths {
		inserted, insertErr := fn(tx, path)
		if insertErr != nil {
			result.Failed++
			result.Errors = appendErrorSample(result.Errors, fmt.Sprintf("%s: %v", path, insertErr))
			continue
		}

		if inserted {
			result.Inserted++
		} else {
			result.Skipped++
		}
	}

	if err := tx.Commit(); err != nil {
		_ = tx.Rollback()
		return result, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return result, nil
}

func hasAllowedExt(path string, allowed map[string]struct{}) bool {
	ext := strings.ToLower(filepath.Ext(path))
	_, ok := allowed[ext]
	return ok
}

func UpdateHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[UpdateHandler] update request started method=%s remote=%s", r.Method, r.RemoteAddr)

		if r.Method != http.MethodGet && r.Method != http.MethodPost {
			writeJSONError(w, http.StatusMethodNotAllowed, "method not allowed")
			return
		}

		if ok, reason := authorizeUpdateRequest(r); !ok {
			writeJSONError(w, http.StatusUnauthorized, reason)
			return
		}

		select {
		case updateSem <- struct{}{}:
			defer func() { <-updateSem }()
		default:
			writeJSONError(w, http.StatusConflict, "update already in progress")
			return
		}

		newmovfiles, err := movScanCompare(db)
		if err != nil {
			writeJSONError(w, http.StatusInternalServerError, fmt.Sprintf("failed to scan movies: %v", err))
			return
		}
		newtvfiles, err := tvshowScanCompare(db)
		if err != nil {
			writeJSONError(w, http.StatusInternalServerError, fmt.Sprintf("failed to scan tvshows: %v", err))
			return
		}
		newvideofiles, err := videoScanCompare(db)
		if err != nil {
			writeJSONError(w, http.StatusInternalServerError, fmt.Sprintf("failed to scan videos: %v", err))
			return
		}

		var moviePaths []string
		if err := json.Unmarshal([]byte(newmovfiles), &moviePaths); err != nil {
			writeJSONError(w, http.StatusInternalServerError, fmt.Sprintf("failed to decode movie paths: %v", err))
			return
		}

		var tvshowPaths []string
		if err := json.Unmarshal([]byte(newtvfiles), &tvshowPaths); err != nil {
			writeJSONError(w, http.StatusInternalServerError, fmt.Sprintf("failed to decode tvshow paths: %v", err))
			return
		}

		var videoPaths []string
		if err := json.Unmarshal([]byte(newvideofiles), &videoPaths); err != nil {
			writeJSONError(w, http.StatusInternalServerError, fmt.Sprintf("failed to decode video paths: %v", err))
			return
		}

		sort.Strings(moviePaths)
		sort.Strings(tvshowPaths)
		sort.Strings(videoPaths)

		movieResult, err := processUpdatesInTx(db, moviePaths, movUpdateInsert)
		if err != nil {
			writeJSONError(w, http.StatusInternalServerError, fmt.Sprintf("failed to process movie updates: %v", err))
			return
		}

		tvResult, err := processUpdatesInTx(db, tvshowPaths, tvshowUpdateInsert)
		if err != nil {
			writeJSONError(w, http.StatusInternalServerError, fmt.Sprintf("failed to process tvshow updates: %v", err))
			return
		}

		videoResult, err := processUpdatesInTx(db, videoPaths, videoUpdateInsert)
		if err != nil {
			writeJSONError(w, http.StatusInternalServerError, fmt.Sprintf("failed to process video updates: %v", err))
			return
		}

		totalFailed := movieResult.Failed + tvResult.Failed + videoResult.Failed
		status := "ok"
		if totalFailed > 0 {
			status = "partial"
		}

		duration := time.Since(start).String()
		log.Printf("[UpdateHandler] update completed status=%s duration=%s movies=%+v tvshows=%+v videos=%+v",
			status, duration, movieResult, tvResult, videoResult)

		writeJSON(w, http.StatusOK, map[string]interface{}{
			"status":   status,
			"method":   r.Method,
			"duration": duration,
			"summary": map[string]interface{}{
				"movies":  movieResult,
				"tvshows": tvResult,
				"videos":  videoResult,
			},
			"moviesInserted":  movieResult.Inserted,
			"tvshowsInserted": tvResult.Inserted,
			"videosInserted":  videoResult.Inserted,
		})

	}
}

func movUpdateInsert(tx *sql.Tx, path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, fmt.Errorf("failed to stat movie %s: %w", path, err)
	}

	filename := filepath.Base(path)
	movId := setup.GenerateMD5(path)
	name := setup.GetMovieName(filename)
	year := setup.GetMovieYear(filename)
	poster := strings.TrimSuffix(path, filepath.Ext(path)) + ".jpg"
	catagory := setup.GetMovieCategory(path)
	thumbPath := setup.GetHttpThumbPath(filename)

	var nextIdx int
	if err := tx.QueryRow(`SELECT COALESCE(MAX(Idx), 0) + 1 FROM movies`).Scan(&nextIdx); err != nil {
		return false, fmt.Errorf("failed to determine next movie index for %s: %w", path, err)
	}

	res, err := tx.Exec(`INSERT OR IGNORE INTO movies (Name, Year, PosterAddr, Size, Path, Idx, MovId, Catagory, HttpThumbPath) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		name, year, poster, fileInfo.Size(), path, nextIdx, movId, catagory, thumbPath)
	if err != nil {
		return false, fmt.Errorf("failed to insert movie %s: %w", path, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("failed to determine movie insert result for %s: %w", path, err)
	}
	return rowsAffected > 0, nil
}

func tvshowUpdateInsert(tx *sql.Tx, path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, fmt.Errorf("failed to stat tvshow %s: %w", path, err)
	}

	filename := filepath.Base(path)
	tvId := setup.GenerateTVId(path)
	name := setup.GetTVShowName(filename)
	season := setup.GetTVShowSeason(filename)
	episode := setup.GetTVShowEpisode(filename)
	category := setup.GetTVShowCategory(path)

	var nextIdx int
	if err := tx.QueryRow(`SELECT COALESCE(MAX(Idx), 0) + 1 FROM tvshows`).Scan(&nextIdx); err != nil {
		return false, fmt.Errorf("failed to determine next tvshow index for %s: %w", path, err)
	}

	res, err := tx.Exec(`INSERT OR IGNORE INTO tvshows (TvId, Size, Catagory, Name, Season, Episode, Path, Idx) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		tvId, fileInfo.Size(), category, name, season, episode, path, nextIdx)
	if err != nil {
		return false, fmt.Errorf("failed to insert tvshow %s: %w", path, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("failed to determine tvshow insert result for %s: %w", path, err)
	}
	return rowsAffected > 0, nil
}

func videoUpdateInsert(tx *sql.Tx, path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, fmt.Errorf("failed to stat video %s: %w", path, err)
	}

	vidId, err := setup.GenerateVidId(path)
	if err != nil {
		return false, fmt.Errorf("failed to generate video id for %s: %w", path, err)
	}
	name := filepath.Base(path)

	var nextIdx int
	if err := tx.QueryRow(`SELECT COALESCE(MAX(Idx), 0) + 1 FROM videos`).Scan(&nextIdx); err != nil {
		return false, fmt.Errorf("failed to determine next video index for %s: %w", path, err)
	}

	res, err := tx.Exec(`INSERT OR IGNORE INTO videos (VidId, VidPath, Size, Name, Idx) VALUES (?, ?, ?, ?, ?)`,
		vidId, path, fileInfo.Size(), name, nextIdx)
	if err != nil {
		return false, fmt.Errorf("failed to insert video %s: %w", path, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("failed to determine video insert result for %s: %w", path, err)
	}
	return rowsAffected > 0, nil
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
		if !d.IsDir() && hasAllowedExt(path, movieExts) {
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
		if !d.IsDir() && hasAllowedExt(path, tvExts) {
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
		if !d.IsDir() && hasAllowedExt(path, videoExts) {
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
