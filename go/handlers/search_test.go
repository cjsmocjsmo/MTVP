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

	_, err = db.Exec(`CREATE TABLE movies (Title TEXT); CREATE TABLE tvshows (Name TEXT);`)
	assert.NoError(t, err)
	_, err = db.Exec(`INSERT INTO movies (Title) VALUES ("Test Movie")`)
	assert.NoError(t, err)

	results := SearchMedia(db, "Test")
	assert.NotEmpty(t, results)
	foundMovie := false
	for _, r := range results {
		if r["Title"] == "Test Movie" {
			foundMovie = true
		}
	}
	assert.True(t, foundMovie)
}

func TestSearchMedia_TVShows(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	assert.NoError(t, err)
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE movies (Title TEXT); CREATE TABLE tvshows (Name TEXT);`)
	assert.NoError(t, err)
	_, err = db.Exec(`INSERT INTO tvshows (Name) VALUES ("Test Show")`)
	assert.NoError(t, err)

	results := SearchMedia(db, "Test")
	assert.NotEmpty(t, results)
	foundShow := false
	for _, r := range results {
		if r["Title"] == "Test Show" {
			foundShow = true
		}
	}
	assert.True(t, foundShow)
}
