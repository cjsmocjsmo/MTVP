package mpvipc

import (
	"net"
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

type mockConn struct {
	writeData []byte
	readData  []byte
	closed    bool
}

func (m *mockConn) Read(b []byte) (n int, err error) {
	copy(b, m.readData)
	return len(m.readData), nil
}
func (m *mockConn) Write(b []byte) (n int, err error) {
	m.writeData = append(m.writeData, b...)
	return len(b), nil
}
func (m *mockConn) Close() error { m.closed = true; return nil }
func (m *mockConn) LocalAddr() net.Addr { return nil }
func (m *mockConn) RemoteAddr() net.Addr { return nil }
func (m *mockConn) SetDeadline(t time.Time) error { return nil }
func (m *mockConn) SetReadDeadline(t time.Time) error { return nil }
func (m *mockConn) SetWriteDeadline(t time.Time) error { return nil }

func TestNew_Error(t *testing.T) {
	// Try to connect to a non-existent socket
	_, err := New("/tmp/nonexistent.sock")
	assert.Error(t, err)
}

func TestCommand(t *testing.T) {
	m := &MPVIPC{conn: &mockConn{readData: []byte(`{"data":123}`)}}
	resp, err := m.Command("get_property", "pause")
	assert.NoError(t, err)
	assert.Equal(t, float64(123), resp["data"])
}
