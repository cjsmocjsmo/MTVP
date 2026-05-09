package server

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	_ "github.com/mattn/go-sqlite3"
	"github.com/gorilla/websocket"
)

func wsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool { return true },
		}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("WebSocket upgrade error:", err)
			return
		}
		HandleWS(conn, db)
	}
}

func staticFileHandler(prefix, dir string) http.Handler {
	return http.StripPrefix(prefix, http.FileServer(http.Dir(dir)))
}

func StartServer() {
	dbPath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("Failed to open DB:", err)
	}
	defer db.Close()

	wsAddr := os.Getenv("MTV_SERVER_ADDR")
	if wsAddr == "" {
		wsAddr = "0.0.0.0"
	}
	log.Println("Starting WebSocket server on ws://" + wsAddr + ":8765/ws")
	log.Println("Starting static file server on http://" + wsAddr + ":8080/thumbnails/ and /tvthumbnails/")

	http.HandleFunc("/ws", wsHandler(db))
	http.Handle("/thumbnails/", staticFileHandler("/thumbnails/", "/usr/share/MTV/thumbnails"))
	http.Handle("/tvthumbnails/", staticFileHandler("/tvthumbnails/", "/usr/share/MTV/tvthumbnails"))
	http.Handle("/templates/", staticFileHandler("/templates/", "go/templates/"))
	http.Handle("/static/css/", staticFileHandler("/static/css/", "go/static/css/"))
	http.Handle("/static/js/", staticFileHandler("/static/js/", "go/static/js/"))
	// Register /action route for action movies page
	http.HandleFunc("/action", ActionPageHandler(db))
	// Register /arnold route for Arnold movies page
	http.HandleFunc("/arnold", ArnoldPageHandler(db))
	// Register /avatar route for Avatar movies page
	http.HandleFunc("/avatar", AvatarPageHandler(db))

	go func() {
		if err := http.ListenAndServe(wsAddr+":8765", nil); err != nil {
			log.Fatal("WebSocket server error:", err)
		}
	}()

	if err := http.ListenAndServe(wsAddr+":8080", nil); err != nil {
		log.Fatal("Static file server error:", err)
	}
}
