package setup

import (
	"database/sql"
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateImgId(t *testing.T) {
	id1 := GenerateImgId("test")
	id2 := GenerateImgId("test")
	id3 := GenerateImgId("different")
	assert.NotEmpty(t, id1)
	assert.Equal(t, id1, id2)
	assert.NotEqual(t, id1, id3)
}

func TestCreateThumbnail(t *testing.T) {
	// Create a temporary image file
	tmpDir := t.TempDir()
	tmpImgPath := filepath.Join(tmpDir, "test.jpg")
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	for y := 0; y < 100; y++ {
		for x := 0; x < 100; x++ {
			img.Set(x, y, color.RGBA{255, 0, 0, 255})
		}
	}
	f, err := os.Create(tmpImgPath)
	assert.NoError(t, err)
	defer f.Close()
	err = jpeg.Encode(f, img, nil)
	assert.NoError(t, err)

	thumbPath, err := CreateThumbnail(tmpImgPath, tmpDir)
	assert.NoError(t, err)
	assert.FileExists(t, thumbPath)
}

func TestResizeImage_ShrinksImage(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 600, 400))
	resized := resizeImage(img, 300, 300)
	assert.Equal(t, 300, resized.Bounds().Dx())
	assert.Equal(t, 200, resized.Bounds().Dy())
}

func TestResizeImage_SmallImage(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	resized := resizeImage(img, 300, 300)
	assert.Equal(t, 100, resized.Bounds().Dx())
	assert.Equal(t, 100, resized.Bounds().Dy())
}

func TestInsertImages(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	assert.NoError(t, err)
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE images (ImgId TEXT, Path TEXT, ImgPath TEXT, Size INTEGER, Name TEXT, ThumbPath TEXT, Idx INTEGER, HttpThumbPath TEXT)`)
	assert.NoError(t, err)

	tmpDir := t.TempDir()
	imgPath := filepath.Join(tmpDir, "img1.jpg")
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	f, err := os.Create(imgPath)
	assert.NoError(t, err)
	defer f.Close()
	err = jpeg.Encode(f, img, nil)
	assert.NoError(t, err)

	thumbDir := filepath.Join(tmpDir, "thumbs")
	err = InsertImages(db, []string{imgPath}, 0, thumbDir, "http://localhost")
	assert.NoError(t, err)
	assert.NoError(t, InsertImages(db, []string{imgPath}, 0, thumbDir, "http://localhost"))

	row := db.QueryRow("SELECT Name FROM images WHERE Path=?", imgPath)
	var name string
	assert.NoError(t, row.Scan(&name))
	assert.Equal(t, "img1", name)

	var thumbPath string
	assert.NoError(t, db.QueryRow("SELECT ThumbPath FROM images WHERE Path=?", imgPath).Scan(&thumbPath))
	sourceTime := time.Now().Add(-time.Minute)
	assert.NoError(t, os.Chtimes(imgPath, sourceTime, sourceTime))
	thumbTime := sourceTime.Add(-time.Second)
	assert.NoError(t, os.Chtimes(thumbPath, thumbTime, thumbTime))
	assert.NoError(t, InsertImages(db, []string{imgPath}, 1, thumbDir, "http://localhost"))

	var imageCount int
	assert.NoError(t, db.QueryRow("SELECT COUNT(*) FROM images WHERE Path=?", imgPath).Scan(&imageCount))
	assert.Equal(t, 1, imageCount)
	thumbInfo, err := os.Stat(thumbPath)
	assert.NoError(t, err)
	assert.True(t, thumbInfo.ModTime().After(sourceTime))
}
