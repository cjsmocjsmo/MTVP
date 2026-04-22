package handlers

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"database/sql"
	"net/http"
	"io/ioutil"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"log"
	"os/exec"
)

func HandleUtilityCommand(conn *websocket.Conn, db *sql.DB, command string) {
	switch command {
	case "weather":
		// Fetch weather for Belfair, WA from National Weather Service
		latitude := 47.4281
		longitude := -122.8189
		pointURL := fmt.Sprintf("https://api.weather.gov/points/%f,%f", latitude, longitude)
		pointResp, err := http.Get(pointURL)
		if err != nil {
			log.Printf("weather fetch failed: %v", err)
			conn.WriteMessage(websocket.TextMessage, []byte(`{"status":"error","message":"weather fetch failed"}`))
			return
		}
		defer pointResp.Body.Close()
		pointData, err := ioutil.ReadAll(pointResp.Body)
		if err != nil {
			log.Printf("weather read failed: %v", err)
			conn.WriteMessage(websocket.TextMessage, []byte(`{"status":"error","message":"weather fetch failed"}`))
			return
		}
		var pointObj map[string]interface{}
		if err := json.Unmarshal(pointData, &pointObj); err != nil {
			log.Printf("weather parse failed: %v", err)
			conn.WriteMessage(websocket.TextMessage, []byte(`{"status":"error","message":"weather parse failed"}`))
			return
		}
		forecastURL, ok := pointObj["properties"].(map[string]interface{})["forecastHourly"].(string)
		if !ok {
			log.Printf("weather forecast URL missing")
			conn.WriteMessage(websocket.TextMessage, []byte(`{"status":"error","message":"weather parse failed"}`))
			return
		}
		weatherResp, err := http.Get(forecastURL)
		if err != nil {
			log.Printf("weather fetch failed: %v", err)
			conn.WriteMessage(websocket.TextMessage, []byte(`{"status":"error","message":"weather fetch failed"}`))
			return
		}
		defer weatherResp.Body.Close()
		weatherData, err := ioutil.ReadAll(weatherResp.Body)
		if err != nil {
			log.Printf("weather read failed: %v", err)
			conn.WriteMessage(websocket.TextMessage, []byte(`{"status":"error","message":"weather fetch failed"}`))
			return
		}
		var weatherObj map[string]interface{}
		if err := json.Unmarshal(weatherData, &weatherObj); err != nil {
			log.Printf("weather parse failed: %v", err)
			conn.WriteMessage(websocket.TextMessage, []byte(`{"status":"error","message":"weather parse failed"}`))
			return
		}
		periods, ok := weatherObj["properties"].(map[string]interface{})["periods"].([]interface{})
		if !ok || len(periods) == 0 {
			log.Printf("weather no periods data")
			conn.WriteMessage(websocket.TextMessage, []byte(`{"status":"error","message":"no weather data"}`))
			return
		}
		current := periods[0].(map[string]interface{})
		resp, _ := json.Marshal(map[string]interface{}{
			"location": "Belfair, WA",
			"temperature": current["temperature"],
			"temperature_unit": current["temperatureUnit"],
			"conditions": current["shortForecast"],
			"winddirection": current["windDirection"],
			"windspeed": current["windSpeed"],
		})
		conn.WriteMessage(websocket.TextMessage, resp)
	case "test":
		conn.WriteMessage(websocket.TextMessage, []byte(`{"status":"Fuck it worked"}`))

	case "getarch":
		arch := "unknown"
		out, err := exec.Command("uname", "-m").Output()
		if err == nil {
			arch = strings.TrimSpace(string(out))
		}
		resp, _ := json.Marshal(arch)
		conn.WriteMessage(websocket.TextMessage, resp)

	case "sqlite3check":
		err := exec.Command("which", "sqlite3").Run()
		result := err == nil
		resp, _ := json.Marshal(result)
		conn.WriteMessage(websocket.TextMessage, resp)

	case "mpvcheck":
		err := exec.Command("which", "mpv").Run()
		result := err == nil
		resp, _ := json.Marshal(result)
		conn.WriteMessage(websocket.TextMessage, resp)

	case "python3mpvcheck":
		err := exec.Command("python3", "-c", "import mpv").Run()
		result := err == nil
		resp, _ := json.Marshal(result)
		conn.WriteMessage(websocket.TextMessage, resp)

	case "python3pilcheck":
		err := exec.Command("python3", "-c", "import PIL").Run()
		result := err == nil
		resp, _ := json.Marshal(result)
		conn.WriteMessage(websocket.TextMessage, resp)

	case "python3dotenvcheck":
		err := exec.Command("python3", "-c", "import dotenv").Run()
		result := err == nil
		resp, _ := json.Marshal(result)
		conn.WriteMessage(websocket.TextMessage, resp)

	case "python3websocketscheck":
		err := exec.Command("python3", "-c", "import websockets").Run()
		result := err == nil
		resp, _ := json.Marshal(result)
		conn.WriteMessage(websocket.TextMessage, resp)

	case "imgwalkdirs":
		dir := os.Getenv("MTV_IMAGES_PATH")
		var jpglist []string
		filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() { return nil }
			if strings.ToLower(filepath.Ext(path)) == ".jpg" {
				jpglist = append(jpglist, path)
			}
			return nil
		})
		resp, _ := json.Marshal(jpglist)
		conn.WriteMessage(websocket.TextMessage, resp)

	case "tvimgwalkdirs":
		dir := os.Getenv("MTV_TV_IMAGES_PATH")
		var imglist []string
		filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() { return nil }
			if strings.ToLower(filepath.Ext(path)) == ".jpg" {
				imglist = append(imglist, path)
			}
			return nil
		})
		resp, _ := json.Marshal(imglist)
		conn.WriteMessage(websocket.TextMessage, resp)
	case "checkformovupdates":
		// Compare movie files on disk with DB and insert new ones
		movieDir := os.Getenv("MTV_MOVIES_PATH")
		rows, err := db.Query("SELECT Path FROM movies")
		if err != nil {
			log.Printf("checkformovupdates: DB query failed: %v", err)
			conn.WriteMessage(websocket.TextMessage, []byte(`[]`))
			return
		}
		defer rows.Close()
		dbPaths := map[string]struct{}{}
		for rows.Next() {
			var path string
			rows.Scan(&path)
			dbPaths[path] = struct{}{}
		}
		var newMovies []string
		exts := map[string]struct{}{ ".mp4":{}, ".mkv":{}, ".avi":{} }
		filepath.Walk(movieDir, func(path string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() { return nil }
			if _, ok := exts[strings.ToLower(filepath.Ext(path))]; ok {
				if _, exists := dbPaths[path]; !exists {
					newMovies = append(newMovies, path)
				}
			}
			return nil
		})
		// Insert new movies into DB
		for _, path := range newMovies {
			name := filepath.Base(path)
			size := int64(0)
			if fi, err := os.Stat(path); err == nil {
				size = fi.Size()
			}
			_, err := db.Exec("INSERT OR IGNORE INTO movies (Name, Path, Size) VALUES (?, ?, ?)", name, path, size)
			if err != nil {
				log.Printf("checkformovupdates: DB insert failed for %s: %v", path, err)
			}
		}
		resp, _ := json.Marshal(newMovies)
		conn.WriteMessage(websocket.TextMessage, resp)
	case "checkfortvupdates":
		// Compare TV show files on disk with DB and insert new ones
		tvDir := os.Getenv("MTV_TV_PATH")
		rows, err := db.Query("SELECT Path FROM tvshows")
		if err != nil {
			log.Printf("checkfortvupdates: DB query failed: %v", err)
			conn.WriteMessage(websocket.TextMessage, []byte(`[]`))
			return
		}
		defer rows.Close()
		dbPaths := map[string]struct{}{}
		for rows.Next() {
			var path string
			rows.Scan(&path)
			dbPaths[path] = struct{}{}
		}
		var newTV []string
		exts := map[string]struct{}{ ".mp4":{}, ".mkv":{}, ".avi":{} }
		filepath.Walk(tvDir, func(path string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() { return nil }
			if _, ok := exts[strings.ToLower(filepath.Ext(path))]; ok {
				if _, exists := dbPaths[path]; !exists {
					newTV = append(newTV, path)
				}
			}
			return nil
		})
		// Insert new TV shows into DB
		for _, path := range newTV {
			name := filepath.Base(path)
			size := int64(0)
			if fi, err := os.Stat(path); err == nil {
				size = fi.Size()
			}
			_, err := db.Exec("INSERT OR IGNORE INTO tvshows (Name, Path, Size) VALUES (?, ?, ?)", name, path, size)
			if err != nil {
				log.Printf("checkfortvupdates: DB insert failed for %s: %v", path, err)
			}
		}
		resp, _ := json.Marshal(newTV)
		conn.WriteMessage(websocket.TextMessage, resp)
	case "movcount":
		var count int
		err := db.QueryRow("SELECT COUNT(*) FROM movies").Scan(&count)
		if err != nil {
			log.Printf("movcount: DB query failed: %v", err)
			conn.WriteMessage(websocket.TextMessage, []byte(`0`))
			return
		}
		resp, _ := json.Marshal(count)
		conn.WriteMessage(websocket.TextMessage, resp)
	case "tvcount":
		var count int
		err := db.QueryRow("SELECT COUNT(*) FROM tvshows").Scan(&count)
		if err != nil {
			log.Printf("tvcount: DB query failed: %v", err)
			conn.WriteMessage(websocket.TextMessage, []byte(`0`))
			return
		}
		resp, _ := json.Marshal(count)
		conn.WriteMessage(websocket.TextMessage, resp)
	case "movsizeondisk":
		rows, err := db.Query("SELECT Size FROM movies")
		if err != nil {
			log.Printf("movsizeondisk: DB query failed: %v", err)
			conn.WriteMessage(websocket.TextMessage, []byte(`0`))
			return
		}
		var total int64
		for rows.Next() {
			var sz int64
			rows.Scan(&sz)
			total += sz
		}
		rows.Close()
		gb := float64(total) / (1024 * 1024 * 1024)
		resp, _ := json.Marshal(gb)
		conn.WriteMessage(websocket.TextMessage, resp)
	case "tvsizeondisk":
		rows, err := db.Query("SELECT Size FROM tvshows")
		if err != nil {
			log.Printf("tvsizeondisk: DB query failed: %v", err)
			conn.WriteMessage(websocket.TextMessage, []byte(`0`))
			return
		}
		var total int64
		for rows.Next() {
			var sz int64
			rows.Scan(&sz)
			total += sz
		}
		rows.Close()
		gb := float64(total) / (1024 * 1024 * 1024)
		resp, _ := json.Marshal(gb)
		conn.WriteMessage(websocket.TextMessage, resp)
	// Add more utility commands here
	}
}
