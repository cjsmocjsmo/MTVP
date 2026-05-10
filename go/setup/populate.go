package setup

import (
	"database/sql"
	"os"
)

// Dummy data population for demonstration. Replace with real directory walking and parsing logic.

func populateMovies(db *sql.DB) error {
	movieDir := os.Getenv("MTVGO_MOVIES_PATH")
	if movieDir == "" {
		return nil // or error if required
	}
	exts := []string{".mp4", ".mkv", ".avi", ".mpg"}
	moviePaths, err := WalkMediaDirs(movieDir, exts)
	if err != nil {
		return err
	}
	if err := InsertMovies(db, moviePaths, 0); err != nil {
		return err
	}

	// Images for movies
	imgDir := os.Getenv("MTVGO_POSTER_PATH")
	thumbDir := os.Getenv("MTVGO_THUMBNAIL_PATH")
	serverAddr := os.Getenv("MTVGO_SERVER_ADDR")
	if imgDir != "" && thumbDir != "" && serverAddr != "" {
		imgExts := []string{".jpg"}
		imgPaths, err := WalkMediaDirs(imgDir, imgExts)
		if err == nil {
			if err := InsertImages(db, imgPaths, 0, thumbDir, serverAddr); err != nil {
				return err
			}
		}
	}
	return nil
}

func populateTVShows(db *sql.DB) error {
	tvDir := os.Getenv("MTVGO_TV_PATH")
	if tvDir == "" {
		return nil // or error if required
	}
	exts := []string{".mp4", ".mkv", ".avi"}
	tvPaths, err := WalkMediaDirs(tvDir, exts)
	if err != nil {
		return err
	}
	if err := InsertTVShows(db, tvPaths, 0); err != nil {
		return err
	}

	// TV show images
	imgDir := os.Getenv("MTVGO_TV_POSTER_PATH")
	thumbDir := os.Getenv("MTVGO_TV_THUMBNAIL_PATH")
	serverAddr := os.Getenv("MTVGO_SERVER_ADDR")
	if imgDir != "" && thumbDir != "" && serverAddr != "" {
		imgExts := []string{".jpg"}
		imgPaths, err := WalkMediaDirs(imgDir, imgExts)
		if err == nil {
			if err := InsertImages(db, imgPaths, 0, thumbDir, serverAddr); err != nil {
				return err
			}
		}
	}
	return nil
}

func populateVideos(db *sql.DB) error {
	vidDir := os.Getenv("MTVGO_VIDEOS_PATH")
	if vidDir == "" {
		return nil // or error if required
	}
	exts := []string{".mp4", ".mkv", ".avi"}
	vidPaths, err := WalkMediaDirs(vidDir, exts)
	if err != nil {
		return err
	}
	return InsertVideos(db, vidPaths, 0)
}
