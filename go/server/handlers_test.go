package server

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIndexHandler(t *testing.T) {
	tmpDir := t.TempDir()
	tmplPath := filepath.Join(tmpDir, "index.html")
	err := os.WriteFile(tmplPath, []byte("<html><body>Test</body></html>"), 0644)
	assert.NoError(t, err)

	req := httptest.NewRequest("GET", "/", nil)
	rw := httptest.NewRecorder()

	handler := indexHandlerWithPath(tmplPath)
	handler(rw, req)
	resp := rw.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestWsHandler(t *testing.T) {
	// TODO: Mock http.ResponseWriter, http.Request, and websocket.Upgrader
	assert.True(t, true)
}
