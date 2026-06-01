
package server

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "net/http"
    "os"
    "path/filepath"
)

func UpdateHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		newmovfiles := movScanCompare(db)

	}
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
    tv_dir_path := os.Getenv("MTVGO_TVSHOWS_PATH")
    if tv_dir_path == "" {
        return "", fmt.Errorf("MTVGO_TVSHOWS_PATH not set")
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
    rows, err := db.Query("SELECT Path FROM videos")
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