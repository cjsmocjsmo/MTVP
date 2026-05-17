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
	"runtime"
	"sync"
	"strconv"

	_ "image/png"
)

func GenerateImgId(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func CreateThumbnail(srcPath, dstDir string) (string, error) {
	       // Ensure thumbnail directory exists
	       if err := os.MkdirAll(dstDir, 0755); err != nil {
		       return "", fmt.Errorf("failed to create thumbnail directory %s: %w", dstDir, err)
	       }

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
	type job struct {
		idx  int
		path string
	}
	type result struct {
		idx         int
		path        string
		imgId       string
		name        string
		thumbPath   string
		size        int64
		httpThumbPath string
		err         error
	}

	numWorkers := runtime.NumCPU()
	if env := os.Getenv("MTVGO_IMAGE_WORKERS"); env != "" {
		if n, err := strconv.Atoi(env); err == nil && n > 0 {
			numWorkers = n
		}
	}

	jobs := make(chan job, len(imgPaths))
	results := make(chan result, len(imgPaths))
	var wg sync.WaitGroup

	worker := func() {
		defer wg.Done()
		for j := range jobs {
			imgId := GenerateImgId(j.path)
			name := strings.TrimSuffix(filepath.Base(j.path), filepath.Ext(j.path))
			thumbPath, err := CreateThumbnail(j.path, thumbDir)
			if err != nil {
				results <- result{idx: j.idx, path: j.path, err: fmt.Errorf("failed to create thumbnail for %s: %w", j.path, err)}
				continue
			}
			fileInfo, err := os.Stat(thumbPath)
			if err != nil {
				results <- result{idx: j.idx, path: j.path, err: err}
				continue
			}
			size := fileInfo.Size()
			httpThumbPath := fmt.Sprintf("%s:8080/thumbnails/%s", serverAddr, filepath.Base(thumbPath))
			results <- result{
				idx: j.idx, path: j.path, imgId: imgId, name: name, thumbPath: thumbPath, size: size, httpThumbPath: httpThumbPath, err: nil,
			}
		}
	}

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker()
	}

	for idx, path := range imgPaths {
		jobs <- job{idx, path}
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	var firstErr error
	inserts := make([]result, len(imgPaths))
	for res := range results {
		if res.err != nil && firstErr == nil {
			firstErr = res.err
		}
		inserts[res.idx] = res
	}
	for _, res := range inserts {
		if res.err != nil {
			continue
		}
		_, err := db.Exec(`INSERT OR IGNORE INTO images (ImgId, Path, ImgPath, Size, Name, ThumbPath, Idx, HttpThumbPath) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
			res.imgId, res.path, res.path, res.size, res.name, res.thumbPath, res.idx+idxStart+1, res.httpThumbPath)
		if err != nil && firstErr == nil {
			firstErr = fmt.Errorf("failed to insert image %s: %w", res.path, err)
		}
	}
	return firstErr
}
