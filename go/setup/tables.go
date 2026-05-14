package setup

import (
	"database/sql"
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
	       _, err := db.Exec(movieTable)
	       if err != nil {
		       return err
	       }
	       _, err = db.Exec(tvTable)
	       if err != nil {
		       return err
	       }
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
	       return nil
}
