package setup

import (
	"database/sql"
	"log"
)

func createTables(db *sql.DB) error {
	movieTable := `CREATE TABLE IF NOT EXISTS movies (
		Name TEXT,
		Year TEXT,
		PosterAddr TEXT,
		Size INTEGER,
		Path TEXT,
		Idx INTEGER,
		MovId TEXT PRIMARY KEY,
		Catagory TEXT,
		HttpThumbPath TEXT
	);`
	tvTable := `CREATE TABLE IF NOT EXISTS tvshows (
		TvId TEXT PRIMARY KEY,
		Size INTEGER,
		Catagory TEXT,
		Name TEXT,
		Season TEXT,
		Episode TEXT,
		Path TEXT,
		Idx INTEGER
	);`
		       log.Println("[SETUP] Creating movies table...")
		       _, err := db.Exec(movieTable)
		       if err != nil {
			       return err
		       }
		       log.Println("[SETUP] Creating tvshows table...")
		       _, err = db.Exec(tvTable)
		       if err != nil {
			       return err
		       }
		       log.Println("[SETUP] Creating images table...")
		       imagesTable := `CREATE TABLE IF NOT EXISTS images (
			       ImgId TEXT,
			       Path TEXT,
			       ImgPath TEXT,
			       Size INTEGER,
			       Name TEXT,
			       ThumbPath TEXT,
			       Idx INTEGER,
			       HttpThumbPath TEXT
		       );`
		       _, err = db.Exec(imagesTable)
		       if err != nil {
			       return err
		       }
		       log.Println("[SETUP] Creating videos table...")
		       videosTable := `CREATE TABLE IF NOT EXISTS videos (
			       VidId TEXT PRIMARY KEY,
			       VidPath TEXT,
			       Size INTEGER,
			       Name TEXT,
			       Idx INTEGER
		       );`
		       _, err = db.Exec(videosTable)
		       if err != nil {
			       return err
		       }
		       return nil
}
