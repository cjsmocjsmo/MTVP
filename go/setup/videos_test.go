package setup

import (
	"database/sql"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInsertVideos_TableMissing(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	assert.NoError(t, err)
	defer db.Close()

	// Do NOT create the videos table
	tmpDir := t.TempDir()
	imgPath := filepath.Join(tmpDir, "testvid.mp4")
	// Create a dummy file to simulate a video
	f, err := os.Create(imgPath)
	assert.NoError(t, err)
	defer f.Close()
	img := image.NewRGBA(image.Rect(0, 0, 10, 10))
	jpeg.Encode(f, img, nil) // Just to have some content

	err = InsertVideos(db, []string{imgPath}, 0)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no such table: videos")
}

func TestInsertVideos_TableExists(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	assert.NoError(t, err)
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE videos (VidId TEXT PRIMARY KEY, VidPath TEXT, Size INTEGER, Name TEXT, Idx INTEGER)`)
	assert.NoError(t, err)

	tmpDir := t.TempDir()
	imgPath := filepath.Join(tmpDir, "testvid.mp4")
	f, err := os.Create(imgPath)
	assert.NoError(t, err)
	defer f.Close()
	img := image.NewRGBA(image.Rect(0, 0, 10, 10))
	jpeg.Encode(f, img, nil)

	err = InsertVideos(db, []string{imgPath}, 0)
	assert.NoError(t, err)
}

func TestPopulateVideosSkipsUnchangedFilesAndUpdatesChangedFiles(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	require.NoError(t, err)
	defer db.Close()
	db.SetMaxOpenConns(1)
	require.NoError(t, createTables(db))
	require.NoError(t, ensureMediaScanStateTable(db))

	videoDir := t.TempDir()
	videoPath := filepath.Join(videoDir, "testvid.mp4")
	require.NoError(t, os.WriteFile(videoPath, []byte("first"), 0600))
	t.Setenv("MTVGO_VIDEOS_PATH", videoDir)
	require.NoError(t, populateVideos(db))

	var firstID string
	require.NoError(t, db.QueryRow(`SELECT VidId FROM videos WHERE VidPath = ?`, videoPath).Scan(&firstID))
	require.NoError(t, populateVideos(db))
	var unchangedID string
	require.NoError(t, db.QueryRow(`SELECT VidId FROM videos WHERE VidPath = ?`, videoPath).Scan(&unchangedID))
	assert.Equal(t, firstID, unchangedID)

	require.NoError(t, os.WriteFile(videoPath, []byte("other"), 0600))
	changedTime := time.Now().Add(time.Second)
	require.NoError(t, os.Chtimes(videoPath, changedTime, changedTime))
	require.NoError(t, populateVideos(db))
	var changedID string
	require.NoError(t, db.QueryRow(`SELECT VidId FROM videos WHERE VidPath = ?`, videoPath).Scan(&changedID))
	assert.NotEqual(t, firstID, changedID)
	var videoCount int
	require.NoError(t, db.QueryRow(`SELECT COUNT(*) FROM videos`).Scan(&videoCount))
	assert.Equal(t, 1, videoCount)
}
