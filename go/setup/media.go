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

func GetMovieCategory(filename string) string {
	catagory := ""
	fn := filepath.Dir(filename)
	if strings.Contains(fn, "Action") {
		catagory = "Action"
	} else if strings.Contains(fn, "Arnold") {
		catagory = "Arnold"
	} else if strings.Contains(fn, "Avatar") {
		catagory = "Avatar"
	} else if strings.Contains(fn, "BruceLee") {
		catagory = "BruceLee"
	} else if strings.Contains(fn, "BruceWillis") {
		catagory = "BruceWillis"
	} else if strings.Contains(fn, "Buzz") {
		catagory = "Buzz"
	} else if strings.Contains(fn, "Cartoons") {
		catagory = "Cartoons"
	} else if strings.Contains(fn, "CharlieBrown") {
		catagory = "CharlieBrown"
	} else if strings.Contains(fn, "CheechAndChong") {
		catagory = "CheechAndChong"
	} else if strings.Contains(fn, "ChuckNorris") {
		catagory = "ChuckNorris"
	} else if strings.Contains(fn, "ClintEastwood") {
		catagory = "ClintEastwood"
	} else if strings.Contains(fn, "Comedy") {
		catagory = "Comedy"
	} else if strings.Contains(fn, "Documentary") {
		catagory = "Documentary"
	} else if strings.Contains(fn, "Drama") {
		catagory = "Drama"
	} else if strings.Contains(fn, "Fantasy") {
		catagory = "Fantasy"
	} else if strings.Contains(fn, "GhostBusters") {
		catagory = "GhostBusters"
	} else if strings.Contains(fn, "Godzilla") {
		catagory = "Godzilla"
	} else if strings.Contains(fn, "HarrisonFord") {
		catagory = "HarrisonFord"
	} else if strings.Contains(fn, "HarryPotter") {
		catagory = "HarryPotter"
	} else if strings.Contains(fn, "HellBoy") {
		catagory = "HellBoy"
	} else if strings.Contains(fn, "HomeVids") {
		catagory = "HomeVids"
	} else if strings.Contains(fn, "IndianaJones") {
		catagory = "IndianaJones"
	} else if strings.Contains(fn, "JamesBond") {
		catagory = "JamesBond"
	} else if strings.Contains(fn, "JohnWayne") {
		catagory = "JohnWayne"
	} else if strings.Contains(fn, "JohnWick") {
		catagory = "JohnWick"
	} else if strings.Contains(fn, "JurassicPark") {
		catagory = "JurassicPark"
	} else if strings.Contains(fn, "KevinCostner") {
		catagory = "KevinCostner"
	} else if strings.Contains(fn, "KingsMan") {
		catagory = "KingsMan"
	} else if strings.Contains(fn, "Lego") {
		catagory = "Lego"
	} else if strings.Contains(fn, "MenInBlack") {
		catagory = "MenInBlack"
	} else if strings.Contains(fn, "Minions") {
		catagory = "Minions"
	} else if strings.Contains(fn, "Misc") {
		catagory = "Misc"
	} else if strings.Contains(fn, "Mummy") {
		catagory = "Mummy"
	} else if strings.Contains(fn, "MusicVids") {
		catagory = "MusicVids"
	} else if strings.Contains(fn, "Nature") {
		catagory = "Nature"
	} else if strings.Contains(fn, "NicolasCage") {
		catagory = "NicolasCage"
	} else if strings.Contains(fn, "Oldies") {
		catagory = "Oldies"
	} else if strings.Contains(fn, "Panda") {
		catagory = "Pandas"
	} else if strings.Contains(fn, "Pirates") {
		catagory = "Pirates"
	} else if strings.Contains(fn, "Predator") {
		catagory = "Predator"
	} else if strings.Contains(fn, "Riddick") {
		catagory = "Riddick"
	} else if strings.Contains(fn, "SciFi") {
		catagory = "SciFi"
	} else if strings.Contains(fn, "Science") {
		catagory = "Science"
	} else if strings.Contains(fn, "StarWars") {
		catagory = "StarWars"
	} else if strings.Contains(fn, "Stalone") {
		catagory = "Stalone"
	} else if strings.Contains(fn, "StarTrek") {
		catagory = "StarTrek"
	} else if strings.Contains(fn, "Stooges") {
		catagory = "Stooges"
	} else if strings.Contains(fn, "SuperHeros") {
		catagory = "SuperHeros"
	} else if strings.Contains(fn, "SuperMan") {
		catagory = "SuperMan"
	} else if strings.Contains(fn, "TheRock") {
		catagory = "TheRock"
	} else if strings.Contains(fn, "TinkerBell") {
		catagory = "TinkerBell"
	} else if strings.Contains(fn, "TomCruize") {
		catagory = "TomCruize"
	} else if strings.Contains(fn, "Tremors") {
		catagory = "Tremors"
	} else if strings.Contains(fn, "Transformers") {
		catagory = "Transformers"
	} else if strings.Contains(fn, "Trolls") {
		catagory = "Trolls"
	} else if strings.Contains(fn, "VanDam") {
		catagory = "VanDam"
	} else if strings.Contains(fn, "XMen") {
		catagory = "XMen"
	}

	return catagory
}

func GetHttpThumbPath(filename string) string {
	servAddr := os.Getenv("MTVGO_SERVER_ADDR")
	port := os.Getenv("MTVGO_SERVER_PORT")
	fn1 := strings.TrimSuffix(filename, filepath.Ext(filename))
	fn := filepath.Base(fn1)
	httpThumbPath := servAddr + ":" + port + "/thumbnails/" + fn + ".jpg"
	return httpThumbPath
}

// InsertMovies inserts movie metadata into the database
func InsertMovies(db *sql.DB, moviePaths []string, idxStart int) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin movie insert transaction: %w", err)
	}
	defer func() {
		if tx != nil {
			_ = tx.Rollback()
		}
	}()

	stmt, err := tx.Prepare(`INSERT OR IGNORE INTO movies (Name, Year, PosterAddr, Size, Path, Idx, MovId, Catagory, HttpThumbPath) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return fmt.Errorf("failed to prepare movie insert statement: %w", err)
	}
	defer stmt.Close()

	for idx, path := range moviePaths {
		fileInfo, err := os.Stat(path)
		if err != nil {
			return err
		}
		// fmt.Println("path:", path)
		filename := filepath.Base(path)
		movId := GenerateMD5(path)
		name := GetMovieName(filename)
		year := GetMovieYear(filename)
		size := fileInfo.Size()
		poster := strings.TrimSuffix(path, filepath.Ext(path)) + ".jpg"
		// fmt.Println("Poster path:", poster)
		catagory := GetMovieCategory(path)
		thumbPath := GetHttpThumbPath(filename)
		_, err = stmt.Exec(
			name, year, poster, size, path, idx+idxStart+1, movId, catagory, thumbPath)
		if err != nil {
			return fmt.Errorf("failed to insert movie %s: %w", path, err)
		}
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit movie insert transaction: %w", err)
	}
	tx = nil

	return nil
}
