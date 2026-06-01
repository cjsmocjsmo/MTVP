package setup

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"
)

func GenerateVidId(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	hash := sha256.New()
	buf := make([]byte, 4096)
	for {
		n, err := file.Read(buf)
		if n > 0 {
			hash.Write(buf[:n])
		}
		if err != nil {
			break
		}
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func InsertVideos(db *sql.DB, vidPaths []string, idxStart int) error {
	type job struct {
		idx  int
		path string
	}
	type result struct {
		idx   int
		path  string
		vidId string
		size  int64
		name  string
		err   error
	}

	// Get number of workers from env, fallback to NumCPU if not set or invalid
	numWorkers := runtime.NumCPU()
	if env := os.Getenv("MTVGO_VIDEO_WORKERS"); env != "" {
		if n, err := strconv.Atoi(env); err == nil && n > 0 {
			numWorkers = n
		}
	}

	jobs := make(chan job, len(vidPaths))
	results := make(chan result, len(vidPaths))
	var wg sync.WaitGroup

	// Worker: hash and stat only (no DB)
	worker := func() {
		defer wg.Done()
		for j := range jobs {
			vidId, err := GenerateVidId(j.path)
			if err != nil {
				results <- result{idx: j.idx, path: j.path, err: fmt.Errorf("failed to hash video %s: %w", j.path, err)}
				continue
			}
			name := filepath.Base(j.path)
			fileInfo, err := os.Stat(j.path)
			if err != nil {
				results <- result{idx: j.idx, path: j.path, err: err}
				continue
			}
			size := fileInfo.Size()
			results <- result{idx: j.idx, path: j.path, vidId: vidId, size: size, name: name, err: nil}
		}
	}

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker()
	}

	// Send jobs
	for idx, path := range vidPaths {
		jobs <- job{idx, path}
	}
	close(jobs)

	// Wait for workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results and do DB inserts serially
	var firstErr error
	inserts := make([]result, len(vidPaths))
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
		_, err := db.Exec(`INSERT OR IGNORE INTO videos (VidId, VidPath, Size, Name, Idx) VALUES (?, ?, ?, ?, ?)`,
			res.vidId, res.path, res.size, res.name, res.idx+idxStart+1)
		if err != nil && firstErr == nil {
			firstErr = fmt.Errorf("failed to insert video %s: %w", res.path, err)
		}
	}
	return firstErr
}
