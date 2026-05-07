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

	_, err = db.Exec(`CREATE TABLE movies (Title TEXT); CREATE TABLE tvshows (Name TEXT);`)
	assert.NoError(t, err)
	_, err = db.Exec(`INSERT INTO movies (Title) VALUES ("Integration Movie")`)
	assert.NoError(t, err)

	results := handlers.SearchMedia(db, "Integration")
	assert.NotEmpty(t, results)
	found := false
	for _, r := range results {
		if r["Title"] == "Integration Movie" {
			found = true
		}
	}
	assert.True(t, found)
}
