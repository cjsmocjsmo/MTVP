
package server

import (
    "os/exec"
    "sync"
    "net"
    "encoding/json"
    "path/filepath"
    "time"
    "database/sql"
    "log"
    "fmt"
    "github.com/gorilla/websocket"
    "html/template"
    "net/http"
    "net/url"
    "os"
    "syscall"
    "io"
    "mtvp/setup"
)

// APODResponse maps the exact JSON fields returned by NASA's API
type APODResponse struct {
    Date           string `json:"date"`
    Explanation    string `json:"explanation"`
    HDURL          string `json:"hdurl"`
    MediaType      string `json:"media_type"`
    ServiceVersion string `json:"service_version"`
    Title          string `json:"title"`
    URL            string `json:"url"`
    ThumbnailURL   string `json:"thumbnail_url,omitempty"` // Only present if video and thumbs=true
    Copyright      string `json:"copyright,omitempty"`     // Only present for copyrighted works
    Idx            int    // Database primary key
}

// FetchNASAData hits the APOD endpoint, returns the parsed payload, and inserts it into the nasa table
func FetchNASAData(db *sql.DB) (*APODResponse, error) {
    today := time.Now().Format("2006-01-02")
    // Try to get today's entry from the nasa table
    var apod APODResponse
    row := db.QueryRow(`SELECT Date, Explanation, HDURL, MediaType, ServiceVersion, Title, URL, ThumbnailURL, Copyright, Idx FROM nasa WHERE Date = ? ORDER BY Idx DESC LIMIT 1`, today)
    err := row.Scan(
        &apod.Date,
        &apod.Explanation,
        &apod.HDURL,
        &apod.MediaType,
        &apod.ServiceVersion,
        &apod.Title,
        &apod.URL,
        &apod.ThumbnailURL,
        &apod.Copyright,
        &apod.Idx,
    )
    if err == nil {
        // Found today's entry, return it
        return &apod, nil
    }
    if err != sql.ErrNoRows {
        // Some other DB error
        return nil, fmt.Errorf("failed to query nasa table: %w", err)
    }

    // Not found, fetch from NASA API
    baseURL := "https://api.nasa.gov/planetary/apod"
    apiKey := "c2MSxvl303kuIlMnkhygr6l60lc14bENZm0Mjwik"

    u, err := url.Parse(baseURL)
    if err != nil {
        return nil, err
    }
    q := u.Query()
    q.Set("api_key", apiKey)
    q.Set("thumbs", "true")
    u.RawQuery = q.Encode()

    client := &http.Client{
        Timeout: 10 * time.Second,
    }
    resp, err := client.Get(u.String())
    if err != nil {
        return nil, fmt.Errorf("failed making request to NASA API: %w", err)
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("nasa API returned status code: %d", resp.StatusCode)
    }
    if err := json.NewDecoder(resp.Body).Decode(&apod); err != nil {
        return nil, fmt.Errorf("failed to decode JSON response: %w", err)
    }
    // Insert into nasa table
    insertStmt := `INSERT INTO nasa (Date, Explanation, HDURL, MediaType, ServiceVersion, Title, URL, ThumbnailURL, Copyright)
                  VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
    _, err = db.Exec(insertStmt,
        apod.Date,
        apod.Explanation,
        apod.HDURL,
        apod.MediaType,
        apod.ServiceVersion,
        apod.Title,
        apod.URL,
        apod.ThumbnailURL,
        apod.Copyright,
    )
    if err != nil {
        return nil, fmt.Errorf("failed to insert APODResponse into nasa table: %w", err)
    }
    return &apod, nil
}

func MTVWeather() ([]byte, error) {
	// Fetch weather for Belfair, WA from National Weather Service
	latitude := 47.4281
	longitude := -122.8189
	pointURL := fmt.Sprintf("https://api.weather.gov/points/%f,%f", latitude, longitude)
	pointResp, err := http.Get(pointURL)
	if err != nil {
		return nil, fmt.Errorf("weather fetch failed: %v", err)
	}
	defer pointResp.Body.Close()
    pointData, err := io.ReadAll(pointResp.Body)
	if err != nil {
		return nil, fmt.Errorf("weather read failed: %v", err)
	}
	var pointObj map[string]interface{}
	if err := json.Unmarshal(pointData, &pointObj); err != nil {
		return nil, fmt.Errorf("weather parse failed: %v", err)
	}
	forecastURL, ok := pointObj["properties"].(map[string]interface{})["forecastHourly"].(string)
	if !ok {
		return nil, fmt.Errorf("weather forecast URL missing")
	}
	weatherResp, err := http.Get(forecastURL)
	if err != nil {
		return nil, fmt.Errorf("weather fetch failed: %v", err)
	}
	defer weatherResp.Body.Close()
    weatherData, err := io.ReadAll(weatherResp.Body)
	if err != nil {
		return nil, fmt.Errorf("weather read failed: %v", err)
	}
	var weatherObj map[string]interface{}
	if err := json.Unmarshal(weatherData, &weatherObj); err != nil {
		return nil, fmt.Errorf("weather parse failed: %v", err)
	}
	periods, ok := weatherObj["properties"].(map[string]interface{})["periods"].([]interface{})
	if !ok || len(periods) == 0 {
		return nil, fmt.Errorf("weather no periods data")
	}
	current := periods[0].(map[string]interface{})
	resp, err := json.Marshal(map[string]interface{}{
		"location": "Belfair, WA",
		"temperature": current["temperature"],
		"temperature_unit": current["temperatureUnit"],
		"conditions": current["shortForecast"],
		"winddirection": current["windDirection"],
		"windspeed": current["windSpeed"],
	})
	if err != nil {
		return nil, fmt.Errorf("weather marshal failed: %v", err)
	}
	return resp, nil
}

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
        weatherData, err := MTVWeather()
        var wdLocation, wdTemp, wdUnit, wdConditions, wdWindDir, wdWindSpeed string
        if err != nil {
            log.Println("Error fetching weather:", err)
        } else {
            var weatherMap map[string]interface{}
            if err := json.Unmarshal(weatherData, &weatherMap); err == nil {
                if v, ok := weatherMap["location"].(string); ok {
                    wdLocation = v
                }
                if v, ok := weatherMap["temperature"].(float64); ok {
                    wdTemp = fmt.Sprintf("%.0f", v)
                } else if v, ok := weatherMap["temperature"].(string); ok {
                    wdTemp = v
                }
                if v, ok := weatherMap["temperature_unit"].(string); ok {
                    wdUnit = v
                }
                if v, ok := weatherMap["conditions"].(string); ok {
                    wdConditions = v
                }
                if v, ok := weatherMap["winddirection"].(string); ok {
                    wdWindDir = v
                }
                if v, ok := weatherMap["windspeed"].(string); ok {
                    wdWindSpeed = v
                }
            }
        }
        nasaData, err := FetchNASAData(db)
        if err != nil {
            log.Println("Error fetching NASA data:", err)
        } else {
            // You can use nasaData in the template if needed
            log.Printf("Today's NASA APOD: %s - %s", nasaData.Title, nasaData.URL)
        }
        type Stats struct {
            TotalMovies    int
            TotalTVShows   int
            TotalVideos    int
            MovieSizeOnDisk string
            TVShowSizeOnDisk string
            VideoSizeOnDisk string
            FreeSpaceOnDisk string
            WeatherLocation string
            WeatherTemperature string
            WeatherUnit string
            WeatherConditions string
            WeatherWindDirection string
            WeatherWindSpeed string
            NasaDate string
            NasaExplanation string
            NasaHDURL string
            NasaMediaType string
            NasaServiceVersion string
            NasaTitle string
            NasaURL string
            NasaThumbnailURL string
            NasaCopyright string
            NasaIdx int
            IsNasaVideo bool
            IsNasaImage bool
        }
        stats := Stats{
            TotalMovies:      movCount,
            TotalTVShows:     tvCount,
            TotalVideos:      videoCount,
            MovieSizeOnDisk:  movsizeondisk,
            TVShowSizeOnDisk: tvsizeondisk,
            VideoSizeOnDisk:  videosizeondisk,
            FreeSpaceOnDisk:  freespaceondisk,
            WeatherLocation:  wdLocation,
            WeatherTemperature: wdTemp,
            WeatherUnit:      wdUnit,
            WeatherConditions: wdConditions,
            WeatherWindDirection: wdWindDir,
            WeatherWindSpeed: wdWindSpeed,
            NasaDate: nasaData.Date,
            NasaExplanation: nasaData.Explanation,
            NasaHDURL: nasaData.HDURL,
            NasaMediaType: nasaData.MediaType,
            NasaServiceVersion: nasaData.ServiceVersion,
            NasaTitle: nasaData.Title,
            NasaURL: nasaData.URL,
            NasaThumbnailURL: nasaData.ThumbnailURL,
            NasaCopyright: nasaData.Copyright,
            NasaIdx: nasaData.Idx,
            IsNasaVideo: nasaData.MediaType == "video",
            IsNasaImage: nasaData.MediaType == "image",
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

















func movScanCompare(db *sql.DB) (string, error) {
    mov_dir_path := os.Getenv("MTVGO_MOVIES_PATH")
    if mov_dir_path == "" {
        return "", fmt.Errorf("MTVGO_MOVIES_PATH not set")
    }

    // 1. Walk the directory and collect all file paths
    fsPaths := []string{}
    err := filepath.WalkDir(mov_dir_path, func(path string, d os.DirEntry, err error) error {
        if err != nil {
            return err
        }
        if !d.IsDir() {
            fsPaths = append(fsPaths, path)
        }
        return nil
    })
    if err != nil {
        return "", fmt.Errorf("error walking movie directory: %w", err)
    }

    // 2. Get all movie paths from the database
    rows, err := db.Query("SELECT Path FROM movies")
    if err != nil {
        return "", fmt.Errorf("database query error: %w", err)
    }
    defer rows.Close()
    dbPaths := map[string]struct{}{}
    for rows.Next() {
        var dbPath string
        if err := rows.Scan(&dbPath); err == nil {
            dbPaths[dbPath] = struct{}{}
        }
    }
    if err := rows.Err(); err != nil {
        return "", fmt.Errorf("database rows error: %w", err)
    }

    // 3. Compare and collect new paths
    newPaths := []string{}
    for _, path := range fsPaths {
        if _, exists := dbPaths[path]; !exists {
            newPaths = append(newPaths, path)
        }
    }

    // 4. Return new paths as a JSON array string
    result, err := json.Marshal(newPaths)
    if err != nil {
        return "", fmt.Errorf("json marshal error: %w", err)
    }
    return string(result), nil
}

func tvshowScanCompare(db *sql.DB) (string, error) {
    tv_dir_path := os.Getenv("MTVGO_TVSHOWS_PATH")
    if tv_dir_path == "" {
        return "", fmt.Errorf("MTVGO_TVSHOWS_PATH not set")
    }

    // 1. Walk the directory and collect all file paths
    fsPaths := []string{}
    err := filepath.WalkDir(tv_dir_path, func(path string, d os.DirEntry, err error) error {
        if err != nil {
            return err
        }
        if !d.IsDir() {
            fsPaths = append(fsPaths, path)
        }
        return nil
    })
    if err != nil {
        return "", fmt.Errorf("error walking TV show directory: %w", err)
    }

    // 2. Get all TV show paths from the database
    rows, err := db.Query("SELECT Path FROM tvshows")
    if err != nil {
        return "", fmt.Errorf("database query error: %w", err)
    }
    defer rows.Close()
    dbPaths := map[string]struct{}{}
    for rows.Next() {
        var dbPath string
        if err := rows.Scan(&dbPath); err == nil {
            dbPaths[dbPath] = struct{}{}
        }
    }
    if err := rows.Err(); err != nil {
        return "", fmt.Errorf("database rows error: %w", err)
    }

    // 3. Compare and collect new paths
    newPaths := []string{}
    for _, path := range fsPaths {
        if _, exists := dbPaths[path]; !exists {
            newPaths = append(newPaths, path)
        }
    }

    // 4. Return new paths as a JSON array string
    result, err := json.Marshal(newPaths)
    if err != nil {
        return "", fmt.Errorf("json marshal error: %w", err)
    }
    return string(result), nil
}

func videoScanCompare(db *sql.DB) (string, error) {
    video_dir_path := os.Getenv("MTVGO_VIDEOS_PATH")
    if video_dir_path == "" {
        return "", fmt.Errorf("MTVGO_VIDEOS_PATH not set")
    }

    // 1. Walk the directory and collect all file paths
    fsPaths := []string{}
    err := filepath.WalkDir(video_dir_path, func(path string, d os.DirEntry, err error) error {
        if err != nil {
            return err
        }
        if !d.IsDir() {
            fsPaths = append(fsPaths, path)
        }
        return nil
    })
    if err != nil {
        return "", fmt.Errorf("error walking video directory: %w", err)
    }

    // 2. Get all video paths from the database
    rows, err := db.Query("SELECT Path FROM videos")
    if err != nil {
        return "", fmt.Errorf("database query error: %w", err)
    }
    defer rows.Close()
    dbPaths := map[string]struct{}{}
    for rows.Next() {
        var dbPath string
        if err := rows.Scan(&dbPath); err == nil {
            dbPaths[dbPath] = struct{}{}
        }
    }
    if err := rows.Err(); err != nil {
        return "", fmt.Errorf("database rows error: %w", err)
    }

    // 3. Compare and collect new paths
    newPaths := []string{}
    for _, path := range fsPaths {
        if _, exists := dbPaths[path]; !exists {
            newPaths = append(newPaths, path)
        }
    }

    // 4. Return new paths as a JSON array string
    result, err := json.Marshal(newPaths)
    if err != nil {
        return "", fmt.Errorf("json marshal error: %w", err)
    }
    return string(result), nil
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
        case "home_set_media":
            homevidId, _ := data["media_id"].(string)
            log.Printf("[HandleWS] Received 'home_set_media' command. homevidId: %v", homevidId)
            if homevidId != "" {
                var path string
                log.Printf("[HandleWS] Querying DB for homevidId: %v", homevidId)
                err := db.QueryRow("SELECT VidPath FROM videos WHERE VidId = ?", homevidId).Scan(&path)
                if err != nil {
                    log.Printf("[HandleWS] Media not found for homevidId %v: %v", homevidId, err)
                    sendJSON(conn, map[string]interface{}{ "status": "error", "message": "media not found" })
                } else {
                    log.Printf("[HandleWS] Found media path for homevidId %v: %v", homevidId, path)
                    log.Printf("[HandleWS] Attempting to start player with path: %v", path)
                    if err := player.StartMPV(path); err != nil {
                        log.Printf("[HandleWS] Error starting player for path %v: %v", path, err)
                        sendJSON(conn, map[string]interface{}{ "status": "error", "message": err.Error() })
                    } else {
                        log.Printf("[HandleWS] Media set successfully for homevidId %v", homevidId)
                        sendJSON(conn, map[string]interface{}{ "status": "media_set" })
                    }
                }
            } else {
                log.Printf("[HandleWS] 'home_set_media' command received with empty homevidId")
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

