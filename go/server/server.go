package server

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"github.com/gorilla/websocket"
	_ "github.com/mattn/go-sqlite3"
)
// indexHandler serves the main page using Go's html/template
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("go", "templates", "index.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
	}
}

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

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ws", wsHandler(db))
	http.Handle("/thumbnails/", staticFileHandler("/thumbnails/", "/usr/share/MTV/thumbnails"))
	http.Handle("/tvthumbnails/", staticFileHandler("/tvthumbnails/", "/usr/share/MTV/tvthumbnails"))

	go func() {
		if err := http.ListenAndServe(wsAddr+":8765", nil); err != nil {
			log.Fatal("WebSocket server error:", err)
		}
	}()

	if err := http.ListenAndServe(wsAddr+":8080", nil); err != nil {
		log.Fatal("Static file server error:", err)
	}
}
