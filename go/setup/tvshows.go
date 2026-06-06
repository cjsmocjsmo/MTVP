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

type TVShowFile struct {
	TvId     string
	Size     int64
	Category string
	Name     string
	Season   string
	Episode  string
	Path     string
	Idx      int
}

func GenerateTVId(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func GetTVShowName(filename string) string {
	// Extract name before SxxExx
	re := regexp.MustCompile(`(?i)(.*?)\sS\d{2}E\d{2}`)
	match := re.FindStringSubmatch(filename)
	if len(match) > 1 {
		return strings.TrimSpace(match[1])
	}
	return filename
}

func GetTVShowSeason(filename string) string {
	re := regexp.MustCompile(`(?i)S(\d{2})E\d{2}`)
	match := re.FindStringSubmatch(filename)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func GetTVShowEpisode(filename string) string {
	re := regexp.MustCompile(`(?i)S\d{2}E(\d{2})`)
	match := re.FindStringSubmatch(filename)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func GetTVShowCategory(name string) string {
	fn2 := filepath.Dir(name)
	cleaned := filepath.Clean(fn2)
	parts := strings.Split(cleaned, string(os.PathSeparator))
	if len(parts) > 0 {
		return parts[len(parts)-2]
	} else {
		return "Unknown"
	}
}

func InsertTVShows(db *sql.DB, tvPaths []string, idxStart int) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin tvshow insert transaction: %w", err)
	}
	defer func() {
		if tx != nil {
			_ = tx.Rollback()
		}
	}()

	stmt, err := tx.Prepare(`INSERT OR IGNORE INTO tvshows (TvId, Size, Catagory, Name, Season, Episode, Path, Idx) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return fmt.Errorf("failed to prepare tvshow insert statement: %w", err)
	}
	defer stmt.Close()

	for idx, path := range tvPaths {
		fileInfo, err := os.Stat(path)
		if err != nil {
			return err
		}
		filename := filepath.Base(path)
		tvId := GenerateTVId(path)
		name := GetTVShowName(filename)
		season := GetTVShowSeason(filename)
		episode := GetTVShowEpisode(filename)
		size := fileInfo.Size()
		category := GetTVShowCategory(path)
		_, err = stmt.Exec(
			tvId, size, category, name, season, episode, path, idx+idxStart+1)
		if err != nil {
			return fmt.Errorf("failed to insert tvshow %s: %w", path, err)
		}
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit tvshow insert transaction: %w", err)
	}
	tx = nil

	return nil
}
