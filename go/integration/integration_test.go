package integration

import (
	"database/sql"
	"testing"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"mtvp/handlers"
)

func TestEndToEnd_MovieInsertAndSearch(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	assert.NoError(t, err)
	defer db.Close()

	       _, err = db.Exec(`CREATE TABLE movies (
		       Name TEXT,
		       Year TEXT,
		       PosterAddr TEXT,
		       Size INTEGER,
		       Path TEXT,
		       Idx INTEGER,
		       MovId TEXT PRIMARY KEY,
		       Catagory TEXT,
		       HttpThumbPath TEXT
	       );
	       CREATE TABLE tvshows (
		       TvId TEXT PRIMARY KEY,
		       Size INTEGER,
		       Catagory TEXT,
		       Name TEXT,
		       Season TEXT,
		       Episode TEXT,
		       Path TEXT,
		       Idx INTEGER
	       );`)
	       assert.NoError(t, err)
	       _, err = db.Exec(`INSERT INTO movies (Name, Year, PosterAddr, Size, Path, Idx, MovId, Catagory, HttpThumbPath) VALUES ("Integration Movie", "2022", "poster.jpg", 12345, "/path/to/movie", 1, "id1", "action", "thumb.jpg")`)
	       assert.NoError(t, err)

	results := handlers.SearchMedia(db, "Integration")
	assert.NotEmpty(t, results)
	found := false
	       for _, r := range results {
		       if r["Name"] == "Integration Movie" {
			       found = true
		       }
	       }
	       assert.True(t, found)
}
