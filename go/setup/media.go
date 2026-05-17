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
	if strings.Contains(filename, "Action") {
		catagory = "Action"
	} else if strings.Contains(filename, "Arnold") {
		catagory = "Arnold"
	} else if strings.Contains(filename, "Avatar") {
		catagory = "Avatar"
	} else if strings.Contains(filename, "BruceLee") {
		catagory = "BruceLee"
	} else if strings.Contains(filename, "BruceWillis") {
		catagory = "BruceWillis"
	} else if strings.Contains(filename, "Buzz") {
		catagory = "Buzz"
	} else if strings.Contains(filename, "Cartoons") {
		catagory = "Cartoons"
	} else if strings.Contains(filename, "CharlieBrown") {
		catagory = "CharlieBrown"
	} else if strings.Contains(filename, "CheechAndChong") {
		catagory = "CheechAndChong"
	} else if strings.Contains(filename, "ChuckNorris") {
		catagory = "ChuckNorris"
	} else if strings.Contains(filename, "ClintEastwood") {
		catagory = "ClintEastwood"
	} else if strings.Contains(filename, "Comedy") {
		catagory = "Comedy"
	} else if strings.Contains(filename, "Documentary") {
		catagory = "Documentary"
	} else if strings.Contains(filename, "Drama") {
		catagory = "Drama"
	} else if strings.Contains(filename, "Fantasy") {
		catagory = "Fantasy"
	} else if strings.Contains(filename, "GhostBusters") {
		catagory = "GhostBusters"
	} else if strings.Contains(filename, "Godzilla") {
		catagory = "Godzilla"
	} else if strings.Contains(filename, "HarrisonFord") {
		catagory = "HarrisonFord"
	} else if strings.Contains(filename, "HarryPotter") {
		catagory = "HarryPotter"
	} else if strings.Contains(filename, "HellBoy") {
		catagory = "HellBoy"
	} else if strings.Contains(filename, "HomeVids") {
		catagory = "HomeVids"
	} else if strings.Contains(filename, "IndianaJones") {
		catagory = "IndianaJones"
	} else if strings.Contains(filename, "JamesBond") {
		catagory = "JamesBond"
	} else if strings.Contains(filename, "JohnWayne") {
		catagory = "JohnWayne"
	} else if strings.Contains(filename, "JohnWick") {
		catagory = "JohnWick"
	} else if strings.Contains(filename, "JurassicPark") {
		catagory = "JurassicPark"
	} else if strings.Contains(filename, "KevinCostner") {
		catagory = "KevinCostner"
	} else if strings.Contains(filename, "KingsMan") {
		catagory = "KingsMan"
	} else if strings.Contains(filename, "Lego") {
		catagory = "Lego"
	} else if strings.Contains(filename, "MenInBlack") {
		catagory = "MenInBlack"
	} else if strings.Contains(filename, "Minions") {
		catagory = "Minions"
	} else if strings.Contains(filename, "Misc") {
		catagory = "Misc"
	} else if strings.Contains(filename, "Mummy") {
		catagory = "Mummy"
	} else if strings.Contains(filename, "MusicVids") {
		catagory = "MusicVids"
	} else if strings.Contains(filename, "Nature") {
		catagory = "Nature"
	} else if strings.Contains(filename, "NicolasCage") {
		catagory = "NicolasCage"
	} else if strings.Contains(filename, "Oldies") {
		catagory = "Oldies"
	} else if strings.Contains(filename, "Pandas") {
		catagory = "Pandas"
	} else if strings.Contains(filename, "Pirates") {
		catagory = "Pirates"
	} else if strings.Contains(filename, "Predator") {
		catagory = "Predator"
	} else if strings.Contains(filename, "Riddick") {
		catagory = "Riddick"
	} else if strings.Contains(filename, "SciFi") {
		catagory = "SciFi"
	} else if strings.Contains(filename, "Science") {
		catagory = "Science"
	} else if strings.Contains(filename, "StarWars") {
		catagory = "StarWars"
	} else if strings.Contains(filename, "Stalone") {
		catagory = "Stalone"
	} else if strings.Contains(filename, "StarTrek") {
		catagory = "StarTrek"
	} else if strings.Contains(filename, "Stooges") {
		catagory = "Stooges"
	} else if strings.Contains(filename, "SuperHeros") {
		catagory = "SuperHeros"
	} else if strings.Contains(filename, "Superman") {
		catagory = "Superman"
	} else if strings.Contains(filename, "TheRock") {
		catagory = "TheRock"
	} else if strings.Contains(filename, "TinkerBell") {
		catagory = "TinkerBell"
	} else if strings.Contains(filename, "TomCruise") {
		catagory = "TomCruise"
	} else if strings.Contains(filename, "Tremors") {
		catagory = "Tremors"
	} else if strings.Contains(filename, "Transformers") {
		catagory = "Transformers"
	} else if strings.Contains(filename, "Trolls") {
		catagory = "Trolls"
	} else if strings.Contains(filename, "VanDam") {
		catagory = "VanDam"
	} else if strings.Contains(filename, "XMen") {
		catagory = "XMen"
	}
	return catagory
}

func GetHttpThumbPath(filename string) string {
	servAddr := os.Getenv("MTVGO_SERVER_ADDR")
	port := os.Getenv("MTVGO_SERVER_PORT")
	fn := filepath.Base(filename)
	httpThumbPath := servAddr + ":" + port + "/thumbnails/" + fn + ".jpg"
	return httpThumbPath
}

// InsertMovies inserts movie metadata into the database
func InsertMovies(db *sql.DB, moviePaths []string, idxStart int) error {
	for idx, path := range moviePaths {
		fileInfo, err := os.Stat(path)
		if err != nil {
			return err
		}
		fmt.Println("path:", path)
		filename := filepath.Base(path)
		movId := GenerateMD5(path)
		name := GetMovieName(filename)
		year := GetMovieYear(filename)
		size := fileInfo.Size()
		poster := strings.TrimSuffix(path, filepath.Ext(path)) + ".jpg"
		fmt.Println("Poster path:", poster)
		category := GetMovieCategory(filename)
		thumbPath := GetHttpThumbPath(filename)
		_, err = db.Exec(`INSERT OR IGNORE INTO movies (Name, Year, PosterAddr, Size, Path, Idx, MovId, Category, HttpThumbPath) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			name, year, poster, size, path, idx+idxStart+1, movId, category, thumbPath)
		if err != nil {
			return fmt.Errorf("failed to insert movie %s: %w", path, err)
		}
	}
	return nil
}
