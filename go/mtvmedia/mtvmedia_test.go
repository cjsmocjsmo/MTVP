package mtvmedia

import (
	"database/sql"
	"testing"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	assert.NoError(t, err)
	return db
}

func TestCategoryMovies(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()
	_, err := db.Exec(`CREATE TABLE movies (catagory TEXT, year INTEGER, Title TEXT)`)
	assert.NoError(t, err)
	_, err = db.Exec(`INSERT INTO movies (catagory, year, Title) VALUES (?, ?, ?)`, "action", 2022, "Action Movie")
	assert.NoError(t, err)

	media := NewMTVMedia(db)
	results, err := media.CategoryMovies("action")
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, "Action Movie", results[0]["Title"])
}

func TestCategoryTVShows(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()
	_, err := db.Exec(`CREATE TABLE tvshows (catagory TEXT, season TEXT, Episode INTEGER, Name TEXT)`)
	assert.NoError(t, err)
	_, err = db.Exec(`INSERT INTO tvshows (catagory, season, Episode, Name) VALUES (?, ?, ?, ?)`, "comedy", "1", 1, "Funny Show")
	assert.NoError(t, err)

	media := NewMTVMedia(db)
	results, err := media.CategoryTVShows("comedy", "1")
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, "Funny Show", results[0]["Name"])
}
