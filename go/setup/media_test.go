package setup

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"database/sql"
)

func TestWalkMediaDirs(t *testing.T) {
	tmpDir := t.TempDir()
	file1 := filepath.Join(tmpDir, "movie1.mkv")
	file2 := filepath.Join(tmpDir, "movie2.avi")
	file3 := filepath.Join(tmpDir, "ignore.txt")
	os.WriteFile(file1, []byte("test"), 0644)
	os.WriteFile(file2, []byte("test"), 0644)
	os.WriteFile(file3, []byte("test"), 0644)

	files, err := WalkMediaDirs(tmpDir, []string{".mkv", ".avi"})
	assert.NoError(t, err)
	assert.ElementsMatch(t, []string{file1, file2}, files)
}

func TestGenerateMD5(t *testing.T) {
	hash1 := GenerateMD5("test")
	hash2 := GenerateMD5("test")
	hash3 := GenerateMD5("other")
	assert.NotEmpty(t, hash1)
	assert.Equal(t, hash1, hash2)
	assert.NotEqual(t, hash1, hash3)
}

func TestGetMovieName(t *testing.T) {
	assert.Equal(t, "Movie Title", GetMovieName("Movie Title (2020).mkv"))
	assert.Equal(t, "NoYearTitle.mkv", GetMovieName("NoYearTitle.mkv"))
}

func TestGetMovieYear(t *testing.T) {
	assert.Equal(t, "2020", GetMovieYear("Movie Title (2020).mkv"))
	assert.Equal(t, "", GetMovieYear("NoYearTitle.mkv"))
}

func TestInsertMovies(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	assert.NoError(t, err)
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE movies (Name TEXT, Year TEXT, PosterAddr TEXT, Size INTEGER, Path TEXT, Idx INTEGER, MovId TEXT, Catagory TEXT, HttpThumbPath TEXT)`)
	assert.NoError(t, err)

	tmpDir := t.TempDir()
	moviePath := filepath.Join(tmpDir, "Test Movie (2022).mkv")
	os.WriteFile(moviePath, []byte("test"), 0644)

	err = InsertMovies(db, []string{moviePath}, 0)
	assert.NoError(t, err)

	row := db.QueryRow("SELECT Name, Year FROM movies WHERE Path=?", moviePath)
	var name, year string
	assert.NoError(t, row.Scan(&name, &year))
	assert.Equal(t, "Test Movie", name)
	assert.Equal(t, "2022", year)
}
