package server

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net"
	"net/http"
	"os/exec"
	"sync"
	"time"
)

func sendJSON(conn *websocket.Conn, v interface{}) {
	msg, err := json.Marshal(v)
	if err != nil {
		log.Println("JSON marshal error:", err)
		return
	}
	if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
		log.Println("WebSocket write error:", err)
	}
}

func WSHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     func(r *http.Request) bool { return true },
		}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("WebSocket upgrade error:", err)
			return
		}
		HandleWS(conn, db)
	}
}

// PlayerManager manages the media player process and state
type PlayerManager struct {
	mu      sync.Mutex
	cmd     *exec.Cmd
	playing bool
	paused  bool
	ipcSock string
}

var player = &PlayerManager{ipcSock: "/tmp/mpvsocket"}

func (p *PlayerManager) StartMPV(path string) error {
	log.Printf("[PlayerManager] StartMPV called with path: %s", path)
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.cmd != nil && p.playing {
		log.Printf("[PlayerManager] Killing existing MPV process")
		p.cmd.Process.Kill()
		time.Sleep(500 * time.Millisecond)
	}
	log.Printf("[PlayerManager] Removing IPC socket: %s", p.ipcSock)
	_ = exec.Command("rm", "-f", p.ipcSock).Run()
	p.cmd = exec.Command("mpv", "--fs", "--volume=100", "--input-ipc-server="+p.ipcSock, path)
	err := p.cmd.Start()
	if err == nil {
		log.Printf("[PlayerManager] MPV started successfully for path: %s", path)
		p.playing = true
		p.paused = false
	} else {
		log.Printf("[PlayerManager] Failed to start MPV for path %s: %v", path, err)
	}
	return err
}

func (p *PlayerManager) StopMPV() {
	log.Printf("[PlayerManager] StopMPV called")
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.cmd != nil && p.playing {
		log.Printf("[PlayerManager] Killing MPV process")
		p.cmd.Process.Kill()
		p.playing = false
		p.paused = false
	}
}

func (p *PlayerManager) sendMPVCommand(cmd interface{}) error {
	log.Printf("[PlayerManager] sendMPVCommand: %v", cmd)
	conn, err := net.Dial("unix", p.ipcSock)
	if err != nil {
		log.Printf("[PlayerManager] Failed to dial IPC socket %s: %v", p.ipcSock, err)
		return err
	}
	defer conn.Close()
	b, _ := json.Marshal(cmd)
	_, err = conn.Write(append(b, '\n'))
	if err != nil {
		log.Printf("[PlayerManager] Failed to write command to IPC: %v", err)
	}
	return err
}

func (p *PlayerManager) Pause() error {
	log.Printf("[PlayerManager] Pause called")
	err := p.sendMPVCommand(map[string]interface{}{
		"command": []interface{}{"set_property", "pause", true},
	})
	if err != nil {
		log.Printf("[PlayerManager] Pause error: %v", err)
	}
	return err
}

func (p *PlayerManager) Play() error {
	log.Printf("[PlayerManager] Play called")
	err := p.sendMPVCommand(map[string]interface{}{
		"command": []interface{}{"set_property", "pause", false},
	})
	if err != nil {
		log.Printf("[PlayerManager] Play error: %v", err)
	}
	return err
}

func (p *PlayerManager) Next() error {
	log.Printf("[PlayerManager] Next called")
	err := p.sendMPVCommand(map[string]interface{}{
		"command": []interface{}{"seek", 35, "relative"},
	})
	if err != nil {
		log.Printf("[PlayerManager] Next error: %v", err)
	}
	return err
}

func (p *PlayerManager) Previous() error {
	log.Printf("[PlayerManager] Previous called")
	err := p.sendMPVCommand(map[string]interface{}{
		"command": []interface{}{"seek", -35, "relative"},
	})
	if err != nil {
		log.Printf("[PlayerManager] Previous error: %v", err)
	}
	return err
}

// use wscat -c ws://10.0.4.41:8090/ws
func HandleWS(conn *websocket.Conn, db *sql.DB) {
	defer conn.Close()
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket read error:", err)
			break
		}
		var data map[string]interface{}
		if err := json.Unmarshal(message, &data); err != nil {
			log.Println("JSON unmarshal error:", err)
			continue
		}
		command, _ := data["command"].(string)
		switch command {
		case "set_media":
			mediaID, _ := data["media_id"].(string)
			log.Printf("[HandleWS] Received 'set_media' command. media_id: %v", mediaID)
			if mediaID != "" {
				var path string
				log.Printf("[HandleWS] Querying DB for media_id: %v", mediaID)
				err := db.QueryRow("SELECT Path FROM movies WHERE MovId = ?", mediaID).Scan(&path)
				if err != nil {
					log.Printf("[HandleWS] Media not found for media_id %v: %v", mediaID, err)
					sendJSON(conn, map[string]interface{}{"status": "error", "message": "media not found"})
				} else {
					log.Printf("[HandleWS] Found media path for media_id %v: %v", mediaID, path)
					log.Printf("[HandleWS] Attempting to start player with path: %v", path)
					if err := player.StartMPV(path); err != nil {
						log.Printf("[HandleWS] Error starting player for path %v: %v", path, err)
						sendJSON(conn, map[string]interface{}{"status": "error", "message": err.Error()})
					} else {
						log.Printf("[HandleWS] Media set successfully for media_id %v", mediaID)
						sendJSON(conn, map[string]interface{}{"status": "media_set"})
					}
				}
			} else {
				log.Printf("[HandleWS] 'set_media' command received with empty media_id")
			}
		case "tv_set_media":
			tvID, _ := data["media_id"].(string)
			log.Printf("[HandleWS] Received 'tv_set_media' command. tv_id: %v", tvID)
			if tvID != "" {
				var path string
				log.Printf("[HandleWS] Querying DB for tv_id: %v", tvID)
				err := db.QueryRow("SELECT Path FROM tvshows WHERE TvId = ?", tvID).Scan(&path)
				if err != nil {
					log.Printf("[HandleWS] Media not found for tv_id %v: %v", tvID, err)
					sendJSON(conn, map[string]interface{}{"status": "error", "message": "media not found"})
				} else {
					log.Printf("[HandleWS] Found media path for tv_id %v: %v", tvID, path)
					log.Printf("[HandleWS] Attempting to start player with path: %v", path)
					if err := player.StartMPV(path); err != nil {
						log.Printf("[HandleWS] Error starting player for path %v: %v", path, err)
						sendJSON(conn, map[string]interface{}{"status": "error", "message": err.Error()})
					} else {
						log.Printf("[HandleWS] Media set successfully for tv_id %v", tvID)
						sendJSON(conn, map[string]interface{}{"status": "media_set"})
					}
				}
			} else {
				log.Printf("[HandleWS] 'tv_set_media' command received with empty tv_id")
			}
		case "home_set_media":
			homevidId, _ := data["media_id"].(string)
			log.Printf("[HandleWS] Received 'home_set_media' command. homevidId: %v", homevidId)
			if homevidId != "" {
				var path string
				log.Printf("[HandleWS] Querying DB for homevidId: %v", homevidId)
				err := db.QueryRow("SELECT VidPath FROM videos WHERE VidId = ?", homevidId).Scan(&path)
				if err != nil {
					log.Printf("[HandleWS] Media not found for homevidId %v: %v", homevidId, err)
					sendJSON(conn, map[string]interface{}{"status": "error", "message": "media not found"})
				} else {
					log.Printf("[HandleWS] Found media path for homevidId %v: %v", homevidId, path)
					log.Printf("[HandleWS] Attempting to start player with path: %v", path)
					if err := player.StartMPV(path); err != nil {
						log.Printf("[HandleWS] Error starting player for path %v: %v", path, err)
						sendJSON(conn, map[string]interface{}{"status": "error", "message": err.Error()})
					} else {
						log.Printf("[HandleWS] Media set successfully for homevidId %v", homevidId)
						sendJSON(conn, map[string]interface{}{"status": "media_set"})
					}
				}
			} else {
				log.Printf("[HandleWS] 'home_set_media' command received with empty homevidId")
			}
		case "stop":
			player.StopMPV()
			sendJSON(conn, map[string]interface{}{"status": "stopped"})
		case "play":
			if err := player.Play(); err != nil {
				sendJSON(conn, map[string]interface{}{"status": "error", "message": err.Error()})
			} else {
				sendJSON(conn, map[string]interface{}{"status": "playing"})
			}
		case "pause":
			if err := player.Pause(); err != nil {
				sendJSON(conn, map[string]interface{}{"status": "error", "message": err.Error()})
			} else {
				sendJSON(conn, map[string]interface{}{"status": "paused"})
			}
		case "next":
			if err := player.Next(); err != nil {
				sendJSON(conn, map[string]interface{}{"status": "error", "message": err.Error()})
			} else {
				sendJSON(conn, map[string]interface{}{"status": "next"})
			}
		case "previous":
			if err := player.Previous(); err != nil {
				sendJSON(conn, map[string]interface{}{"status": "error", "message": err.Error()})
			} else {
				sendJSON(conn, map[string]interface{}{"status": "previous"})
			}
		case "test":
			sendJSON(conn, map[string]interface{}{"status": "It worked"})
		default:
			sendJSON(conn, map[string]interface{}{"status": "unknown command"})
		}
	}
}
