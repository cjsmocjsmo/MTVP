package handlers

import (
	"database/sql"
	"testing"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestSearchMedia_Movies(t *testing.T) {
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
	       _, err = db.Exec(`INSERT INTO movies (Name, Year, PosterAddr, Size, Path, Idx, MovId, Catagory, HttpThumbPath) VALUES ("Test Movie", "2022", "poster.jpg", 12345, "/path/to/movie", 1, "id1", "action", "thumb.jpg")`)
	       assert.NoError(t, err)

	results := SearchMedia(db, "Test")
	assert.NotEmpty(t, results)
	foundMovie := false
	       for _, r := range results {
		       if r["Name"] == "Test Movie" {
			       foundMovie = true
		       }
	       }
	       assert.True(t, foundMovie)
}

func TestSearchMedia_TVShows(t *testing.T) {
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
	       _, err = db.Exec(`INSERT INTO tvshows (TvId, Size, Catagory, Name, Season, Episode, Path, Idx) VALUES ("tid1", 54321, "comedy", "Test Show", "01", "01", "/path/to/show", 1)`)
	       assert.NoError(t, err)

	results := SearchMedia(db, "Test")
	assert.NotEmpty(t, results)
	foundShow := false
	       for _, r := range results {
		       if r["Name"] == "Test Show" {
			       foundShow = true
		       }
	       }
	       assert.True(t, foundShow)
}
