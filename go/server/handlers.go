package server

import (
	"database/sql"
	"encoding/json"
	"log"
	"github.com/gorilla/websocket"
)

type WSCommand struct {
	Command   string      `json:"command"`
	MediaID   string      `json:"media_id,omitempty"`
	MediaTVID string      `json:"media_tv_id,omitempty"`
	VideoID   string      `json:"video_id,omitempty"`
	Phrase    string      `json:"phrase,omitempty"`
}

type WSResponse map[string]interface{}

func HandleWS(conn *websocket.Conn, db *sql.DB) {
	defer conn.Close()
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket read error:", err)
			break
		}
		var cmd WSCommand
		if err := json.Unmarshal(msg, &cmd); err != nil {
			log.Println("Invalid WS command:", err)
			conn.WriteJSON(WSResponse{"status": "error", "message": "invalid command"})
			continue
		}
		resp := processCommand(cmd, db)
		if err := conn.WriteJSON(resp); err != nil {
			log.Println("WebSocket write error:", err)
			break
		}
	}
}

func processCommand(cmd WSCommand, db *sql.DB) WSResponse {
	switch cmd.Command {
	case "movcount":
		return WSResponse{"count": getMovieCount(db)}
	case "tvcount":
		return WSResponse{"count": getTVShowCount(db)}
	// Add more command cases as needed
	default:
		return WSResponse{"status": "unknown command"}
	}
}

func getMovieCount(db *sql.DB) int {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM movies").Scan(&count)
	if err != nil {
		log.Println("DB error (movcount):", err)
		return 0
	}
	return count
}

func getTVShowCount(db *sql.DB) int {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM tvshows").Scan(&count)
	if err != nil {
		log.Println("DB error (tvcount):", err)
		return 0
	}
	return count
}
