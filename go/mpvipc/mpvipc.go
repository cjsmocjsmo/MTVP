package mpvipc

import (
	"encoding/json"
	"net"
)

type MPVIPC struct {
	conn net.Conn
}

func New(socketPath string) (*MPVIPC, error) {
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		return nil, err
	}
	return &MPVIPC{conn: conn}, nil
}

// Command sends a command to MPV and returns the response
func (m *MPVIPC) Command(args ...interface{}) (map[string]interface{}, error) {
	cmd := map[string]interface{}{
		"command": args,
	}
	data, err := json.Marshal(cmd)
	if err != nil {
		return nil, err
	}
	_, err = m.conn.Write(append(data, '\n'))
	if err != nil {
		return nil, err
	}
	buf := make([]byte, 4096)
	n, err := m.conn.Read(buf)
	if err != nil {
		return nil, err
	}
	var resp map[string]interface{}
	err = json.Unmarshal(buf[:n], &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Play resumes playback
func (m *MPVIPC) Play() error {
	_, err := m.Command("set_property", "pause", false)
	return err
}

// Pause pauses playback
func (m *MPVIPC) Pause() error {
	_, err := m.Command("set_property", "pause", true)
	return err
}

// Stop stops playback
func (m *MPVIPC) Stop() error {
	_, err := m.Command("stop")
	return err
}

// Seek seeks by seconds (positive or negative)
func (m *MPVIPC) Seek(seconds int) error {
	_, err := m.Command("seek", seconds, "relative")
	return err
}

// LoadFile loads a file for playback
func (m *MPVIPC) LoadFile(path string) error {
	_, err := m.Command("loadfile", path, "replace")
	return err
}
