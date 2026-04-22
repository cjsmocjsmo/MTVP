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

func InsertTVShows(db *sql.DB, tvPaths []string, idxStart int) error {
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
		category := "" // TODO: implement category logic as in Python
		_, err = db.Exec(`INSERT OR IGNORE INTO tvshows (TvId, Size, Catagory, Name, Season, Episode, Path, Idx) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
			tvId, size, category, name, season, episode, path, idx+idxStart+1)
		if err != nil {
			return fmt.Errorf("failed to insert tvshow %s: %w", path, err)
		}
	}
	return nil
}
