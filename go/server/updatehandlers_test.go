package server

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func setupUpdateTestDB(t *testing.T) *sql.DB {
	t.Helper()
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("open db: %v", err)
	}

	schema := []string{
		`CREATE TABLE movies (
			Name TEXT,
			Year TEXT,
			PosterAddr TEXT,
			Size INTEGER,
			Path TEXT UNIQUE,
			Idx INTEGER,
			MovId TEXT,
			Catagory TEXT,
			HttpThumbPath TEXT
		);`,
		`CREATE TABLE tvshows (
			TvId TEXT,
			Size INTEGER,
			Catagory TEXT,
			Name TEXT,
			Season TEXT,
			Episode TEXT,
			Path TEXT UNIQUE,
			Idx INTEGER
		);`,
		`CREATE TABLE videos (
			VidId TEXT,
			VidPath TEXT UNIQUE,
			Size INTEGER,
			Name TEXT,
			Idx INTEGER
		);`,
	}

	for _, stmt := range schema {
		if _, err := db.Exec(stmt); err != nil {
			t.Fatalf("create schema: %v", err)
		}
	}

	t.Cleanup(func() {
		_ = db.Close()
	})
	return db
}

func drainUpdateSemaphore() {
	for {
		select {
		case <-updateSem:
		default:
			return
		}
	}
}

func mustWriteFile(t *testing.T, path string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	if err := os.WriteFile(path, []byte("test"), 0o644); err != nil {
		t.Fatalf("write file: %v", err)
	}
}

func countRows(t *testing.T, db *sql.DB, query string) int {
	t.Helper()
	var count int
	if err := db.QueryRow(query).Scan(&count); err != nil {
		t.Fatalf("count rows failed: %v", err)
	}
	return count
}

func TestUpdateHandler_MethodNotAllowed(t *testing.T) {
	drainUpdateSemaphore()
	t.Cleanup(drainUpdateSemaphore)

	db := setupUpdateTestDB(t)
	h := UpdateHandler(db)

	req := httptest.NewRequest(http.MethodDelete, "/update", nil)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected 405, got %d", rr.Code)
	}

	var payload map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if payload["status"] != "error" {
		t.Fatalf("expected error status, got %v", payload["status"])
	}
}

func TestUpdateHandler_UnauthorizedMissingToken(t *testing.T) {
	drainUpdateSemaphore()
	t.Cleanup(drainUpdateSemaphore)
	t.Setenv("MTVGO_UPDATE_TOKEN", "secret-token")

	db := setupUpdateTestDB(t)
	h := UpdateHandler(db)

	req := httptest.NewRequest(http.MethodGet, "/update", nil)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", rr.Code)
	}

	var payload map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	msg, _ := payload["message"].(string)
	if msg == "" {
		t.Fatalf("expected non-empty unauthorized message")
	}
}

func TestUpdateHandler_UnauthorizedInvalidToken(t *testing.T) {
	drainUpdateSemaphore()
	t.Cleanup(drainUpdateSemaphore)
	t.Setenv("MTVGO_UPDATE_TOKEN", "secret-token")

	db := setupUpdateTestDB(t)
	h := UpdateHandler(db)

	req := httptest.NewRequest(http.MethodGet, "/update", nil)
	req.Header.Set("X-Update-Token", "wrong-token")
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", rr.Code)
	}

	var payload map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if payload["message"] != "invalid update token" {
		t.Fatalf("unexpected message: %v", payload["message"])
	}
}

func TestUpdateHandler_ConflictWhenUpdateInProgress(t *testing.T) {
	drainUpdateSemaphore()
	t.Cleanup(drainUpdateSemaphore)

	db := setupUpdateTestDB(t)
	h := UpdateHandler(db)

	updateSem <- struct{}{}

	req := httptest.NewRequest(http.MethodGet, "/update", nil)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	if rr.Code != http.StatusConflict {
		t.Fatalf("expected 409, got %d", rr.Code)
	}

	var payload map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if payload["message"] != "update already in progress" {
		t.Fatalf("unexpected message: %v", payload["message"])
	}
}

func TestUpdateHandler_SuccessSummaryWithExtensionFiltering(t *testing.T) {
	drainUpdateSemaphore()
	t.Cleanup(drainUpdateSemaphore)
	// Keep update endpoint open (no token), and set env required by update scans + movie thumb path build.
	t.Setenv("MTVGO_UPDATE_TOKEN", "")
	t.Setenv("MTVGO_SERVER_ADDR", "http://127.0.0.1")
	t.Setenv("MTVGO_SERVER_PORT", "8090")

	base := t.TempDir()
	moviesDir := filepath.Join(base, "movies")
	tvDir := filepath.Join(base, "tv")
	videosDir := filepath.Join(base, "videos")
	t.Setenv("MTVGO_MOVIES_PATH", moviesDir)
	t.Setenv("MTVGO_TV_PATH", tvDir)
	t.Setenv("MTVGO_VIDEOS_PATH", videosDir)

	mustWriteFile(t, filepath.Join(moviesDir, "Action", "Alpha (2025).mp4"))
	mustWriteFile(t, filepath.Join(moviesDir, "Action", "ignore.txt"))
	mustWriteFile(t, filepath.Join(tvDir, "SciFi", "Show Name S01E01.mkv"))
	mustWriteFile(t, filepath.Join(tvDir, "SciFi", "note.doc"))
	mustWriteFile(t, filepath.Join(videosDir, "HomeVids", "clip.avi"))
	mustWriteFile(t, filepath.Join(videosDir, "HomeVids", "readme.md"))

	db := setupUpdateTestDB(t)
	h := UpdateHandler(db)

	req := httptest.NewRequest(http.MethodGet, "/update", nil)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d body=%s", rr.Code, rr.Body.String())
	}

	var payload map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if payload["status"] != "ok" {
		t.Fatalf("expected ok status, got %v", payload["status"])
	}

	if int(payload["moviesInserted"].(float64)) != 1 {
		t.Fatalf("expected 1 movie inserted, got %v", payload["moviesInserted"])
	}
	if int(payload["tvshowsInserted"].(float64)) != 1 {
		t.Fatalf("expected 1 tvshow inserted, got %v", payload["tvshowsInserted"])
	}
	if int(payload["videosInserted"].(float64)) != 1 {
		t.Fatalf("expected 1 video inserted, got %v", payload["videosInserted"])
	}

	summary, ok := payload["summary"].(map[string]interface{})
	if !ok {
		t.Fatalf("missing summary object")
	}
	movies, ok := summary["movies"].(map[string]interface{})
	if !ok {
		t.Fatalf("missing movies summary")
	}
	if int(movies["scanned"].(float64)) != 1 {
		t.Fatalf("expected movies scanned=1 after extension filtering, got %v", movies["scanned"])
	}
}

func TestMovUpdateHandler_UpdatesMoviesOnly(t *testing.T) {
	drainUpdateSemaphore()
	t.Cleanup(drainUpdateSemaphore)
	t.Setenv("MTVGO_UPDATE_TOKEN", "")
	t.Setenv("MTVGO_SERVER_ADDR", "http://127.0.0.1")
	t.Setenv("MTVGO_SERVER_PORT", "8090")

	base := t.TempDir()
	moviesDir := filepath.Join(base, "movies")
	tvDir := filepath.Join(base, "tv")
	t.Setenv("MTVGO_MOVIES_PATH", moviesDir)
	t.Setenv("MTVGO_TV_PATH", tvDir)

	mustWriteFile(t, filepath.Join(moviesDir, "Action", "Alpha (2025).mp4"))
	mustWriteFile(t, filepath.Join(tvDir, "SciFi", "Show Name S01E01.mkv"))

	db := setupUpdateTestDB(t)
	h := MovUpdateHandler(db)

	req := httptest.NewRequest(http.MethodGet, "/movupdate", nil)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d body=%s", rr.Code, rr.Body.String())
	}

	var payload map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if payload["status"] != "ok" {
		t.Fatalf("expected ok status, got %v", payload["status"])
	}
	if int(payload["moviesInserted"].(float64)) != 1 {
		t.Fatalf("expected 1 movie inserted, got %v", payload["moviesInserted"])
	}
	if _, exists := payload["tvshowsInserted"]; exists {
		t.Fatalf("did not expect tvshowsInserted in movupdate response")
	}

	moviesCount := countRows(t, db, "SELECT COUNT(*) FROM movies")
	tvshowsCount := countRows(t, db, "SELECT COUNT(*) FROM tvshows")
	if moviesCount != 1 {
		t.Fatalf("expected movies count=1, got %d", moviesCount)
	}
	if tvshowsCount != 0 {
		t.Fatalf("expected tvshows count=0, got %d", tvshowsCount)
	}
}

func TestTVUpdateHandler_UpdatesTVShowsOnly(t *testing.T) {
	drainUpdateSemaphore()
	t.Cleanup(drainUpdateSemaphore)
	t.Setenv("MTVGO_UPDATE_TOKEN", "")

	base := t.TempDir()
	moviesDir := filepath.Join(base, "movies")
	tvDir := filepath.Join(base, "tv")
	t.Setenv("MTVGO_MOVIES_PATH", moviesDir)
	t.Setenv("MTVGO_TV_PATH", tvDir)

	mustWriteFile(t, filepath.Join(moviesDir, "Action", "Alpha (2025).mp4"))
	mustWriteFile(t, filepath.Join(tvDir, "SciFi", "Show Name S01E01.mkv"))

	db := setupUpdateTestDB(t)
	h := TVUpdateHandler(db)

	req := httptest.NewRequest(http.MethodGet, "/tvupdate", nil)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d body=%s", rr.Code, rr.Body.String())
	}

	var payload map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if payload["status"] != "ok" {
		t.Fatalf("expected ok status, got %v", payload["status"])
	}
	if int(payload["tvshowsInserted"].(float64)) != 1 {
		t.Fatalf("expected 1 tvshow inserted, got %v", payload["tvshowsInserted"])
	}
	if _, exists := payload["moviesInserted"]; exists {
		t.Fatalf("did not expect moviesInserted in tvupdate response")
	}

	moviesCount := countRows(t, db, "SELECT COUNT(*) FROM movies")
	tvshowsCount := countRows(t, db, "SELECT COUNT(*) FROM tvshows")
	if moviesCount != 0 {
		t.Fatalf("expected movies count=0, got %d", moviesCount)
	}
	if tvshowsCount != 1 {
		t.Fatalf("expected tvshows count=1, got %d", tvshowsCount)
	}
}
