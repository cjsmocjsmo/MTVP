package setup

import (
	"database/sql"
	"os"
)

// Dummy data population for demonstration. Replace with real directory walking and parsing logic.

func populateMovies(db *sql.DB) error {
	       movieDir := os.Getenv("MTVGO_MOVIES_PATH")
	       log.Println("[SETUP] [Movies] Scanning movie directory:", movieDir)
	       if movieDir == "" {
		       log.Println("[SETUP] [Movies] MTVGO_MOVIES_PATH not set, skipping movie population.")
		       return nil // or error if required
	       }
	       exts := []string{".mp4", ".mkv", ".avi", ".mpg"}
	       moviePaths, err := WalkMediaDirs(movieDir, exts)
	       if err != nil {
		       log.Println("[SETUP] [Movies] Error walking movie dirs:", err)
		       return err
	       }
	       log.Printf("[SETUP] [Movies] Found %d movie files. Inserting...", len(moviePaths))
	       if err := InsertMovies(db, moviePaths, 0); err != nil {
		       log.Println("[SETUP] [Movies] Error inserting movies:", err)
		       return err
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
			       log.Printf("[SETUP] [Movies] Found %d poster images. Inserting...", len(imgPaths))
			       if err := InsertImages(db, imgPaths, 0, thumbDir, serverAddr); err != nil {
				       log.Println("[SETUP] [Movies] Error inserting images:", err)
				       return err
			       }
		       } else {
			       log.Println("[SETUP] [Movies] Error walking poster dirs:", err)
		       }
	       }
	       return nil
}

func populateTVShows(db *sql.DB) error {
	       tvDir := os.Getenv("MTVGO_TV_PATH")
	       log.Println("[SETUP] [TVShows] Scanning TV shows directory:", tvDir)
	       if tvDir == "" {
		       log.Println("[SETUP] [TVShows] MTVGO_TV_PATH not set, skipping TV show population.")
		       return nil // or error if required
	       }
	       exts := []string{".mp4", ".mkv", ".avi"}
	       tvPaths, err := WalkMediaDirs(tvDir, exts)
	       if err != nil {
		       log.Println("[SETUP] [TVShows] Error walking TV show dirs:", err)
		       return err
	       }
	       log.Printf("[SETUP] [TVShows] Found %d TV show files. Inserting...", len(tvPaths))
	       if err := InsertTVShows(db, tvPaths, 0); err != nil {
		       log.Println("[SETUP] [TVShows] Error inserting TV shows:", err)
		       return err
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
			       log.Printf("[SETUP] [TVShows] Found %d TV poster images. Inserting...", len(imgPaths))
			       if err := InsertImages(db, imgPaths, 0, thumbDir, serverAddr); err != nil {
				       log.Println("[SETUP] [TVShows] Error inserting images:", err)
				       return err
			       }
		       } else {
			       log.Println("[SETUP] [TVShows] Error walking TV poster dirs:", err)
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
	       log.Printf("[SETUP] [Videos] Found %d video files. Inserting...", len(vidPaths))
	       return InsertVideos(db, vidPaths, 0)
}
