package mpvipc

import (
	"encoding/json"
	"net"
	"log"
)

type MPVIPC struct {
	conn net.Conn
}

func New(socketPath string) (*MPVIPC, error) {
	log.Printf("[MPVIPC] Attempting to connect to socket: %s", socketPath)
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		log.Printf("[MPVIPC] Failed to connect to socket %s: %v", socketPath, err)
		return nil, err
	}
	log.Printf("[MPVIPC] Connected to socket: %s", socketPath)
	return &MPVIPC{conn: conn}, nil
}

// Command sends a command to MPV and returns the response
func (m *MPVIPC) Command(args ...interface{}) (map[string]interface{}, error) {
	log.Printf("[MPVIPC] Sending command: %v", args)
	cmd := map[string]interface{}{
		"command": args,
	}
	data, err := json.Marshal(cmd)
	if err != nil {
		log.Printf("[MPVIPC] Failed to marshal command %v: %v", args, err)
		return nil, err
	}
	_, err = m.conn.Write(append(data, '\n'))
	if err != nil {
		log.Printf("[MPVIPC] Failed to write command %v: %v", args, err)
		return nil, err
	}
	buf := make([]byte, 4096)
	n, err := m.conn.Read(buf)
	if err != nil {
		log.Printf("[MPVIPC] Failed to read response for command %v: %v", args, err)
		return nil, err
	}
	var resp map[string]interface{}
	err = json.Unmarshal(buf[:n], &resp)
	if err != nil {
		log.Printf("[MPVIPC] Failed to unmarshal response for command %v: %v", args, err)
		return nil, err
	}
	log.Printf("[MPVIPC] Command response: %v", resp)
	return resp, nil
}

// Play resumes playback
func (m *MPVIPC) Play() error {
	log.Printf("[MPVIPC] Play called")
	_, err := m.Command("set_property", "pause", false)
	if err != nil {
		log.Printf("[MPVIPC] Play error: %v", err)
	}
	return err
}

// Pause pauses playback
func (m *MPVIPC) Pause() error {
	log.Printf("[MPVIPC] Pause called")
	_, err := m.Command("set_property", "pause", true)
	if err != nil {
		log.Printf("[MPVIPC] Pause error: %v", err)
	}
	return err
}

// Stop stops playback
func (m *MPVIPC) Stop() error {
	log.Printf("[MPVIPC] Stop called")
	_, err := m.Command("stop")
	if err != nil {
		log.Printf("[MPVIPC] Stop error: %v", err)
	}
	return err
}

// Seek seeks by seconds (positive or negative)
func (m *MPVIPC) Seek(seconds int) error {
	log.Printf("[MPVIPC] Seek called: %d seconds", seconds)
	_, err := m.Command("seek", seconds, "relative")
	if err != nil {
		log.Printf("[MPVIPC] Seek error: %v", err)
	}
	return err
}

// LoadFile loads a file for playback
func (m *MPVIPC) LoadFile(path string) error {
	log.Printf("[MPVIPC] LoadFile called: %s", path)
	_, err := m.Command("loadfile", path, "replace")
	if err != nil {
		log.Printf("[MPVIPC] LoadFile error: %v", err)
	}
	return err
}
