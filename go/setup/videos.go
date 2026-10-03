package setup

import (
	"crypto/sha256"
	"database/sql"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"
)

const vidSampleSize int64 = 1 << 20

// GenerateVidId fingerprints a video from its size plus up to three 1 MiB
// samples (start, middle, end) instead of reading the whole file.
func GenerateVidId(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return "", err
	}
	size := info.Size()

	hash := sha256.New()
	var sizeBuf [8]byte
	binary.LittleEndian.PutUint64(sizeBuf[:], uint64(size))
	hash.Write(sizeBuf[:])

	if size <= 3*vidSampleSize {
		if _, err := io.Copy(hash, file); err != nil {
			return "", err
		}
		return hex.EncodeToString(hash.Sum(nil)), nil
	}

	offsets := []int64{0, size/2 - vidSampleSize/2, size - vidSampleSize}
	for _, off := range offsets {
		if _, err := io.Copy(hash, io.NewSectionReader(file, off, vidSampleSize)); err != nil {
			return "", err
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
	// Hashing is I/O-bound, so oversubscribe CPUs
	numWorkers := runtime.NumCPU() * 2
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

	tx, err := db.Begin()
	if err != nil {
		if firstErr != nil {
			return firstErr
		}
		return fmt.Errorf("failed to begin video insert transaction: %w", err)
	}
	defer func() {
		if tx != nil {
			_ = tx.Rollback()
		}
	}()

	stmt, err := tx.Prepare(`INSERT OR IGNORE INTO videos (VidId, VidPath, Size, Name, Idx) VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		if firstErr != nil {
			return firstErr
		}
		return fmt.Errorf("failed to prepare video insert statement: %w", err)
	}
	defer stmt.Close()

	for _, res := range inserts {
		if res.err != nil {
			continue
		}
		_, err := stmt.Exec(
			res.vidId, res.path, res.size, res.name, res.idx+idxStart+1)
		if err != nil && firstErr == nil {
			firstErr = fmt.Errorf("failed to insert video %s: %w", res.path, err)
		}
	}

	if err := tx.Commit(); err != nil && firstErr == nil {
		firstErr = fmt.Errorf("failed to commit video insert transaction: %w", err)
	}
	tx = nil

	return firstErr
}
