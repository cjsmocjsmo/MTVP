
package server

import (
    "os/exec"
    "sync"
    "net"
    "encoding/json"
    "time"
    "database/sql"
    "log"
    "fmt"
    "github.com/gorilla/websocket"
    "html/template"
    "net/http"
    "syscall"
)

// HomePageHandler serves the index.html page for the root path
func HomePageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        movCount := getMovieCount(db)
        tvCount := getTVShowCount(db)
        videoCount := getVideoCount(db)
        movsizeondisk := getMoviesSizeOnDisk(db)
        tvsizeondisk := getTVShowsSizeOnDisk(db)
        videosizeondisk := getVideosSizeOnDisk(db)
        freespaceondisk := freeSpaceOnDisk("/")
        type Stats struct {
            TotalMovies    int
            TotalTVShows   int
            TotalVideos    int
            MovieSizeOnDisk string
            TVShowSizeOnDisk string
            VideoSizeOnDisk string
            FreeSpaceOnDisk string
        }
        stats := Stats{
            TotalMovies:      movCount,
            TotalTVShows:     tvCount,
            TotalVideos:      videoCount,
            MovieSizeOnDisk:  movsizeondisk,
            TVShowSizeOnDisk: tvsizeondisk,
            VideoSizeOnDisk:  videosizeondisk,
            FreeSpaceOnDisk:  freespaceondisk,
        }
        tmpl, err := template.ParseFiles("templates/index.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        err = tmpl.Execute(w, stats)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}



func MovSearchHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        query := r.URL.Query().Get("q")
        if query == "" {
            http.Error(w, "Missing search query", http.StatusBadRequest)
            return
        }
        rows, err := db.Query("SELECT * FROM movies WHERE Name LIKE ?", "%"+query+"%")
        if err != nil {
            http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        defer rows.Close()
        cols, _ := rows.Columns()
        results := []map[string]interface{}{}
        for rows.Next() {
            vals := make([]interface{}, len(cols))
            valPtrs := make([]interface{}, len(cols))
            for i := range vals {
                valPtrs[i] = &vals[i]
            }
            if err := rows.Scan(valPtrs...); err == nil {
                row := make(map[string]interface{})
                for i, col := range cols {
                    b, ok := vals[i].([]byte)
                    if ok {
                        row[col] = string(b)
                    } else {
                        row[col] = vals[i]
                    }
                }
                results = append(results, row)
            }
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]interface{}{
            "results": results,
        })
    }
}

// TVSearchHandler returns JSON search results for TV shows
func TVSearchHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        query := r.URL.Query().Get("q")
        if query == "" {
            http.Error(w, "Missing search query", http.StatusBadRequest)
            return
        }
        rows, err := db.Query("SELECT * FROM tvshows WHERE Name LIKE ?", "%"+query+"%")
        if err != nil {
            http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        defer rows.Close()
        cols, _ := rows.Columns()
        results := []map[string]interface{}{}
        for rows.Next() {
            vals := make([]interface{}, len(cols))
            valPtrs := make([]interface{}, len(cols))
            for i := range vals {
                valPtrs[i] = &vals[i]
            }
            if err := rows.Scan(valPtrs...); err == nil {
                row := make(map[string]interface{})
                for i, col := range cols {
                    b, ok := vals[i].([]byte)
                    if ok {
                        row[col] = string(b)
                    } else {
                        row[col] = vals[i]
                    }
                }
                results = append(results, row)
            }
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]interface{}{
            "results": results,
        })
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
        "command": []interface{}{ "set_property", "pause", true },
    })
    if err != nil {
        log.Printf("[PlayerManager] Pause error: %v", err)
    }
    return err
}

func (p *PlayerManager) Play() error {
    log.Printf("[PlayerManager] Play called")
    err := p.sendMPVCommand(map[string]interface{}{
        "command": []interface{}{ "set_property", "pause", false },
    })
    if err != nil {
        log.Printf("[PlayerManager] Play error: %v", err)
    }
    return err
}

func (p *PlayerManager) Next() error {
    log.Printf("[PlayerManager] Next called")
    err := p.sendMPVCommand(map[string]interface{}{
        "command": []interface{}{ "seek", 35, "relative" },
    })
    if err != nil {
        log.Printf("[PlayerManager] Next error: %v", err)
    }
    return err
}

func (p *PlayerManager) Previous() error {
    log.Printf("[PlayerManager] Previous called")
    err := p.sendMPVCommand(map[string]interface{}{
        "command": []interface{}{ "seek", -35, "relative" },
    })
    if err != nil {
        log.Printf("[PlayerManager] Previous error: %v", err)
    }
    return err
}

//use wscat -c ws://10.0.4.41:8090/ws
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
                    sendJSON(conn, map[string]interface{}{ "status": "error", "message": "media not found" })
                } else {
                    log.Printf("[HandleWS] Found media path for media_id %v: %v", mediaID, path)
                    log.Printf("[HandleWS] Attempting to start player with path: %v", path)
                    if err := player.StartMPV(path); err != nil {
                        log.Printf("[HandleWS] Error starting player for path %v: %v", path, err)
                        sendJSON(conn, map[string]interface{}{ "status": "error", "message": err.Error() })
                    } else {
                        log.Printf("[HandleWS] Media set successfully for media_id %v", mediaID)
                        sendJSON(conn, map[string]interface{}{ "status": "media_set" })
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
                    sendJSON(conn, map[string]interface{}{ "status": "error", "message": "media not found" })
                } else {
                    log.Printf("[HandleWS] Found media path for tv_id %v: %v", tvID, path)
                    log.Printf("[HandleWS] Attempting to start player with path: %v", path)
                    if err := player.StartMPV(path); err != nil {
                        log.Printf("[HandleWS] Error starting player for path %v: %v", path, err)
                        sendJSON(conn, map[string]interface{}{ "status": "error", "message": err.Error() })
                    } else {
                        log.Printf("[HandleWS] Media set successfully for tv_id %v", tvID)
                        sendJSON(conn, map[string]interface{}{ "status": "media_set" })
                    }
                }
            } else {
                log.Printf("[HandleWS] 'tv_set_media' command received with empty tv_id")
            }
        case "stop":
            player.StopMPV()
            sendJSON(conn, map[string]interface{}{ "status": "stopped" })
        case "play":
            if err := player.Play(); err != nil {
                sendJSON(conn, map[string]interface{}{ "status": "error", "message": err.Error() })
            } else {
                sendJSON(conn, map[string]interface{}{ "status": "playing" })
            }
        case "pause":
            if err := player.Pause(); err != nil {
                sendJSON(conn, map[string]interface{}{ "status": "error", "message": err.Error() })
            } else {
                sendJSON(conn, map[string]interface{}{ "status": "paused" })
            }
        case "next":
            if err := player.Next(); err != nil {
                sendJSON(conn, map[string]interface{}{ "status": "error", "message": err.Error() })
            } else {
                sendJSON(conn, map[string]interface{}{ "status": "next" })
            }
        case "previous":
            if err := player.Previous(); err != nil {
                sendJSON(conn, map[string]interface{}{ "status": "error", "message": err.Error() })
            } else {
                sendJSON(conn, map[string]interface{}{ "status": "previous" })
            }
        case "test":
            sendJSON(conn, map[string]interface{}{ "status": "It worked" })
        default:
            sendJSON(conn, map[string]interface{}{ "status": "unknown command" })
        }
    }
}

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

type WSResponse map[string]interface{}

// getCategoryMovieImages queries the DB for movies in a given category and returns a list of image URLs
func getCategoryMovieImages(db *sql.DB, category string) []map[string]interface{} {
    query := "SELECT * FROM movies WHERE Catagory=? ORDER BY Year DESC"
    rows, err := db.Query(query, category)
    if err != nil {
        log.Println("DB error (category images):", err)
        return nil
    }
    defer rows.Close()
    cols, _ := rows.Columns()
    results := []map[string]interface{}{}
    for rows.Next() {
        vals := make([]interface{}, len(cols))
        valPtrs := make([]interface{}, len(cols))
        for i := range vals {
            valPtrs[i] = &vals[i]
        }
        if err := rows.Scan(valPtrs...); err == nil {
            row := make(map[string]interface{})
            for i, col := range cols {
                b, ok := vals[i].([]byte)
                if ok {
                    row[col] = string(b)
                } else {
                    row[col] = vals[i]
                }
            }
            results = append(results, row)
        }
    }
    return results
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

func getVideoCount(db *sql.DB) int {
    var count int
    err := db.QueryRow("SELECT COUNT(*) FROM videos").Scan(&count)
    if err != nil {
        log.Println("DB error (videocount):", err)
        return 0
    }
    return count
}

func getMoviesSizeOnDisk(db *sql.DB) string {
    var size int64
    err := db.QueryRow("SELECT SUM(Size) FROM movies").Scan(&size)
    if err != nil {
        log.Println("DB error (moviedisk):", err)
        return "0 GB"
    }
    return bytestoGB(size)
}

func getTVShowsSizeOnDisk(db *sql.DB) string {
    var size int64
    err := db.QueryRow("SELECT SUM(Size) FROM tvshows").Scan(&size)
    if err != nil {
        log.Println("DB error (tvshowdisk):", err)
        return "0 GB"
    }
    return bytestoGB(size)
}

func getVideosSizeOnDisk(db *sql.DB) string {
    var size int64
    err := db.QueryRow("SELECT SUM(Size) FROM videos").Scan(&size)
    if err != nil {
        log.Println("DB error (videodisk):", err)
        return "0 GB"
    }
    return bytestoGB(size)
}

func bytestoGB(bytes int64) string {
    gb := float64(bytes) / (1024 * 1024 * 1024)
    return fmt.Sprintf("%.2f GB", gb)
}

func freeSpaceOnDisk(path string) string {
    var stat syscall.Statfs_t
    err := syscall.Statfs(path, &stat)
    if err != nil {
        log.Println("Disk error (freespace):", err)
        return "0 GB"
    }
    free := stat.Bavail * uint64(stat.Bsize)
    return bytestoGB(int64(free))
}
