package setup

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// MediaFile holds metadata for a media file
// (expand as needed for TV shows, images, etc.)
type MediaFile struct {
	Name      string
	Year      string
	Poster    string
	Size      int64
	Path      string
	Idx       int
	MovId     string
	Category  string
	ThumbPath string
}

// WalkMediaDirs walks a directory and returns a list of media file paths with allowed extensions
func WalkMediaDirs(root string, exts []string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			ext := strings.ToLower(filepath.Ext(path))
			for _, allowed := range exts {
				if ext == allowed {
					files = append(files, path)
					break
				}
			}
		}
		return nil
	})
	return files, err
}

// GenerateMD5 returns the MD5 hash of a string
func GenerateMD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// GetMovieName extracts the movie name from the filename (before ' (')
func GetMovieName(filename string) string {
	idx := strings.Index(filename, " (")
	if idx > 0 {
		return filename[:idx]
	}
	return filename
}

// GetMovieYear extracts the year from the filename (inside parentheses)
func GetMovieYear(filename string) string {
	re := regexp.MustCompile(`\((\d{4})\)`)
	match := re.FindStringSubmatch(filename)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

// InsertMovies inserts movie metadata into the database
func InsertMovies(db *sql.DB, moviePaths []string, idxStart int) error {
	for idx, path := range moviePaths {
		fileInfo, err := os.Stat(path)
		if err != nil {
			return err
		}
		filename := filepath.Base(path)
		movId := GenerateMD5(path)
		name := GetMovieName(filename)
		year := GetMovieYear(filename)
		size := fileInfo.Size()
		poster := strings.TrimSuffix(path, filepath.Ext(path)) + ".jpg"
		category := "" // TODO: implement category logic as in Python
		thumbPath := "" // TODO: implement HTTP thumb path logic
		_, err = db.Exec(`INSERT OR IGNORE INTO movies (Name, Year, PosterAddr, Size, Path, Idx, MovId, Catagory, HttpThumbPath) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			name, year, poster, size, path, idx+idxStart+1, movId, category, thumbPath)
		if err != nil {
			return fmt.Errorf("failed to insert movie %s: %w", path, err)
		}
	}
	return nil
}
