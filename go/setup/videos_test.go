package setup

import (
	"database/sql"
	"os"
	"path/filepath"
	"testing"
	"image"
	"image/jpeg"
	"github.com/stretchr/testify/assert"
	_ "github.com/mattn/go-sqlite3"
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
