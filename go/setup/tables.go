package setup

import (
	"database/sql"
)

func createTables(db *sql.DB) error {
	movieTable := `CREATE TABLE IF NOT EXISTS movies (
		MovId INTEGER PRIMARY KEY,
		Title TEXT,
		Path TEXT
	);`
	tvTable := `CREATE TABLE IF NOT EXISTS tvshows (
		TvId INTEGER PRIMARY KEY,
		Title TEXT,
		Path TEXT
	);`
	_, err := db.Exec(movieTable)
	if err != nil {
		return err
	}
	_, err = db.Exec(tvTable)
	if err != nil {
		return err
	}
	return nil
}
