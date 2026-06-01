package server

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHomePageHandler(t *testing.T) {
	t.Skip("requires live external APIs and template/runtime setup")
}

func TestWsHandler(t *testing.T) {
	// TODO: Mock http.ResponseWriter, http.Request, and websocket.Upgrader
	assert.True(t, true)
}
