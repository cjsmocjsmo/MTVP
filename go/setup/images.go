package setup

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
	"strings"

	_ "image/png"
)

func GenerateImgId(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func CreateThumbnail(srcPath, dstDir string) (string, error) {
	file, err := os.Open(srcPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return "", err
	}

	thumb := resizeImage(img, 300, 300)
	fname := filepath.Base(srcPath)
	thumbPath := filepath.Join(dstDir, fname)
	thumbFile, err := os.Create(thumbPath)
	if err != nil {
		return "", err
	}
	defer thumbFile.Close()

	if err := jpeg.Encode(thumbFile, thumb, &jpeg.Options{Quality: 85}); err != nil {
		return "", err
	}
	return thumbPath, nil
}

func resizeImage(img image.Image, maxW, maxH int) image.Image {
	// Simple nearest-neighbor resize (replace with a better algorithm if needed)
	bounds := img.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	if w <= maxW && h <= maxH {
		return img
	}
	ratio := float64(w) / float64(h)
	var newW, newH int
	if ratio > 1 {
		newW = maxW
		newH = int(float64(maxW) / ratio)
	} else {
		newH = maxH
		newW = int(float64(maxH) * ratio)
	}
	thumb := image.NewRGBA(image.Rect(0, 0, newW, newH))
	for y := 0; y < newH; y++ {
		for x := 0; x < newW; x++ {
			srcX := x * w / newW
			srcY := y * h / newH
			thumb.Set(x, y, img.At(srcX, srcY))
		}
	}
	return thumb
}

func InsertImages(db *sql.DB, imgPaths []string, idxStart int, thumbDir string, serverAddr string) error {
	for idx, path := range imgPaths {
		imgId := GenerateImgId(path)
		name := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
		thumbPath, err := CreateThumbnail(path, thumbDir)
		if err != nil {
			return fmt.Errorf("failed to create thumbnail for %s: %w", path, err)
		}
		fileInfo, err := os.Stat(thumbPath)
		if err != nil {
			return err
		}
		size := fileInfo.Size()
		httpThumbPath := fmt.Sprintf("%s:8080/thumbnails/%s", serverAddr, filepath.Base(thumbPath))
		_, err = db.Exec(`INSERT OR IGNORE INTO images (ImgId, Path, ImgPath, Size, Name, ThumbPath, Idx, HttpThumbPath) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
			imgId, path, path, size, name, thumbPath, idx+idxStart+1, httpThumbPath)
		if err != nil {
			return fmt.Errorf("failed to insert image %s: %w", path, err)
		}
	}
	return nil
}
