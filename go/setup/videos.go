package setup

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
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
	for idx, path := range vidPaths {
		vidId, err := GenerateVidId(path)
		if err != nil {
			return fmt.Errorf("failed to hash video %s: %w", path, err)
		}
		name := filepath.Base(path)
		fileInfo, err := os.Stat(path)
		if err != nil {
			return err
		}
		size := fileInfo.Size()
		_, err = db.Exec(`INSERT OR IGNORE INTO videos (VidId, VidPath, Size, Name, Idx) VALUES (?, ?, ?, ?, ?)`,
			vidId, path, size, name, idx+idxStart+1)
		if err != nil {
			return fmt.Errorf("failed to insert video %s: %w", path, err)
		}
	}
	return nil
}
