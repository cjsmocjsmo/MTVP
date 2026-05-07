package setup

import (
	"database/sql"
	"testing"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestCreateTables(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	assert.NoError(t, err)
	defer db.Close()

	err = createTables(db)
	assert.NoError(t, err)

	// Check if movies table exists
	row := db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='movies'")
	var name string
	assert.NoError(t, row.Scan(&name))
	assert.Equal(t, "movies", name)

	// Check if tvshows table exists
	row = db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='tvshows'")
	assert.NoError(t, row.Scan(&name))
	assert.Equal(t, "tvshows", name)
}
