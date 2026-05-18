package server

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHomePageHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rw := httptest.NewRecorder()

	handler := HomePageHandler()
	handler(rw, req)
	resp := rw.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestWsHandler(t *testing.T) {
	// TODO: Mock http.ResponseWriter, http.Request, and websocket.Upgrader
	assert.True(t, true)
}
