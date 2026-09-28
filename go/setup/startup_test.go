package setup

import (
	"database/sql"
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"path/filepath"
	"strings"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRunCreatesSchemaForNewDatabase(t *testing.T) {
	dbPath := filepath.Join(t.TempDir(), "mtvp.db")
	t.Setenv("MTVGO_DB_PATH", dbPath)
	t.Setenv("MTVGO_MOVIES_PATH", "")
	t.Setenv("MTVGO_TV_PATH", "")
	t.Setenv("MTVGO_VIDEOS_PATH", "")
	t.Setenv("MTVGO_POSTER_PATH", "")
	t.Setenv("MTVGO_TV_POSTER_PATH", "")

	require.NoError(t, Run())

	db, err := sql.Open("sqlite3", dbPath)
	require.NoError(t, err)
	defer db.Close()
	assert.NoError(t, validateSchema(db))
}

func TestRunRejectsExistingIncompleteDatabase(t *testing.T) {
	dbPath := filepath.Join(t.TempDir(), "partial.db")
	require.NoError(t, os.WriteFile(dbPath, nil, 0600))
	t.Setenv("MTVGO_DB_PATH", dbPath)

	err := Run()
	require.Error(t, err)
	assert.True(t, strings.Contains(err.Error(), "existing database schema is incomplete"))

	db, err := sql.Open("sqlite3", dbPath)
	require.NoError(t, err)
	defer db.Close()
	var tableCount int
	require.NoError(t, db.QueryRow(`SELECT COUNT(*) FROM sqlite_master WHERE type='table'`).Scan(&tableCount))
	assert.Zero(t, tableCount)
}

func TestRunAddsNewMoviesWithoutDuplicatingExistingRows(t *testing.T) {
	root := t.TempDir()
	dbPath := filepath.Join(root, "mtvp.db")
	moviesPath := filepath.Join(root, "movies")
	require.NoError(t, os.MkdirAll(moviesPath, 0755))
	firstMovie := filepath.Join(moviesPath, "First (2020).mkv")
	require.NoError(t, os.WriteFile(firstMovie, []byte("first"), 0600))

	t.Setenv("MTVGO_DB_PATH", dbPath)
	t.Setenv("MTVGO_MOVIES_PATH", moviesPath)
	t.Setenv("MTVGO_TV_PATH", "")
	t.Setenv("MTVGO_VIDEOS_PATH", "")
	t.Setenv("MTVGO_POSTER_PATH", "")
	t.Setenv("MTVGO_TV_POSTER_PATH", "")

	require.NoError(t, Run())
	secondMovie := filepath.Join(moviesPath, "Second (2021).mkv")
	require.NoError(t, os.WriteFile(secondMovie, []byte("second"), 0600))
	require.NoError(t, Run())
	require.NoError(t, Run())

	db, err := sql.Open("sqlite3", dbPath)
	require.NoError(t, err)
	defer db.Close()
	var movieCount int
	require.NoError(t, db.QueryRow(`SELECT COUNT(*) FROM movies`).Scan(&movieCount))
	assert.Equal(t, 2, movieCount)
	var distinctIndexes int
	require.NoError(t, db.QueryRow(`SELECT COUNT(DISTINCT Idx) FROM movies`).Scan(&distinctIndexes))
	assert.Equal(t, 2, distinctIndexes)
}

func TestPopulateTVShowsAddsOnlyNewEpisodes(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	require.NoError(t, err)
	defer db.Close()
	db.SetMaxOpenConns(1)
	require.NoError(t, createTables(db))

	tvDir := t.TempDir()
	showDir := filepath.Join(tvDir, "Example")
	require.NoError(t, os.MkdirAll(showDir, 0755))
	firstEpisode := filepath.Join(showDir, "Example S01E01.mkv")
	require.NoError(t, os.WriteFile(firstEpisode, []byte("first"), 0600))
	t.Setenv("MTVGO_TV_PATH", tvDir)
	t.Setenv("MTVGO_TV_POSTER_PATH", "")

	require.NoError(t, populateTVShows(db))
	secondEpisode := filepath.Join(showDir, "Example S01E02.mkv")
	require.NoError(t, os.WriteFile(secondEpisode, []byte("second"), 0600))
	require.NoError(t, populateTVShows(db))
	require.NoError(t, populateTVShows(db))

	var episodeCount, distinctIndexes int
	require.NoError(t, db.QueryRow(`SELECT COUNT(*) FROM tvshows`).Scan(&episodeCount))
	require.NoError(t, db.QueryRow(`SELECT COUNT(DISTINCT Idx) FROM tvshows`).Scan(&distinctIndexes))
	assert.Equal(t, 2, episodeCount)
	assert.Equal(t, 2, distinctIndexes)
}

func TestPopulateMoviesScansPostersWithoutMovieDirectory(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	require.NoError(t, err)
	defer db.Close()
	db.SetMaxOpenConns(1)
	require.NoError(t, createTables(db))

	root := t.TempDir()
	posterDir := filepath.Join(root, "posters")
	require.NoError(t, os.MkdirAll(posterDir, 0755))
	posterPath := filepath.Join(posterDir, "poster.jpg")
	poster, err := os.Create(posterPath)
	require.NoError(t, err)
	img := image.NewRGBA(image.Rect(0, 0, 10, 10))
	img.Set(0, 0, color.RGBA{R: 255, A: 255})
	require.NoError(t, jpeg.Encode(poster, img, nil))
	require.NoError(t, poster.Close())

	t.Setenv("MTVGO_MOVIES_PATH", "")
	t.Setenv("MTVGO_POSTER_PATH", posterDir)
	t.Setenv("MTVGO_THUMBNAIL_PATH", filepath.Join(root, "thumbnails"))
	t.Setenv("MTVGO_SERVER_ADDR", "127.0.0.1")
	t.Setenv("MTVGO_SERVER_PORT", "8090")

	require.NoError(t, populateMovies(db))
	var imageCount int
	require.NoError(t, db.QueryRow(`SELECT COUNT(*) FROM images WHERE Path = ?`, posterPath).Scan(&imageCount))
	assert.Equal(t, 1, imageCount)
}
