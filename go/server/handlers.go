package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"syscall"
	"time"
	// "mtvp/setup"
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

func MTVWeather(db *sql.DB) ([]byte, error) {
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

	temperature := ""
	switch v := current["temperature"].(type) {
	case float64:
		temperature = fmt.Sprintf("%.0f", v)
	case string:
		temperature = v
	}

	temperatureUnit, _ := current["temperatureUnit"].(string)
	conditions, _ := current["shortForecast"].(string)
	windDirection, _ := current["windDirection"].(string)
	windSpeed, _ := current["windSpeed"].(string)
	humidity := ""
	if rh, ok := current["relativeHumidity"].(map[string]interface{}); ok {
		switch v := rh["value"].(type) {
		case float64:
			humidity = fmt.Sprintf("%.0f", v)
		case string:
			humidity = v
		}
	}

	_, err = db.Exec(
		`INSERT INTO weather (FetchedAt, Location, Temperature, TemperatureUnit, Conditions, WindDirection, WindSpeed, Humidity)
         VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		time.Now().Format(time.RFC3339),
		"Belfair, WA",
		temperature,
		temperatureUnit,
		conditions,
		windDirection,
		windSpeed,
		humidity,
	)
	if err != nil {
		return nil, fmt.Errorf("weather insert failed: %v", err)
	}

	resp, err := json.Marshal(map[string]interface{}{
		"location":         "Belfair, WA",
		"temperature":      temperature,
		"temperature_unit": temperatureUnit,
		"conditions":       conditions,
		"winddirection":    windDirection,
		"windspeed":        windSpeed,
		"humidity":         humidity,
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
		weatherData, err := MTVWeather(db)
		var wdLocation, wdTemp, wdUnit, wdConditions, wdWindDir, wdWindSpeed, wdHumidity string
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
				if v, ok := weatherMap["humidity"].(float64); ok {
					wdHumidity = fmt.Sprintf("%.0f", v)
				} else if v, ok := weatherMap["humidity"].(string); ok {
					wdHumidity = v
				}
			}
		}
		nasaData, err := FetchNASAData(db)
		if err != nil {
			log.Println("Error fetching NASA data:", err)
			// Fallback: use the most recent entry already in the DB
			nasaData = &APODResponse{}
			fallbackRow := db.QueryRow(`SELECT Date, Explanation, HDURL, MediaType, ServiceVersion, Title, URL, ThumbnailURL, Copyright, Idx FROM nasa ORDER BY Date DESC LIMIT 1`)
			if scanErr := fallbackRow.Scan(
				&nasaData.Date,
				&nasaData.Explanation,
				&nasaData.HDURL,
				&nasaData.MediaType,
				&nasaData.ServiceVersion,
				&nasaData.Title,
				&nasaData.URL,
				&nasaData.ThumbnailURL,
				&nasaData.Copyright,
				&nasaData.Idx,
			); scanErr != nil {
				log.Println("No NASA fallback row available:", scanErr)
			} else {
				log.Printf("Using NASA fallback row: %s - %s", nasaData.Title, nasaData.URL)
			}
		} else {
			// You can use nasaData in the template if needed
			log.Printf("Today's NASA APOD: %s - %s", nasaData.Title, nasaData.URL)
		}
		type Stats struct {
			TotalMovies          int
			TotalTVShows         int
			TotalVideos          int
			MovieSizeOnDisk      string
			TVShowSizeOnDisk     string
			VideoSizeOnDisk      string
			FreeSpaceOnDisk      string
			WeatherLocation      string
			WeatherTemperature   string
			WeatherUnit          string
			WeatherConditions    string
			WeatherWindDirection string
			WeatherWindSpeed     string
			WeatherHumidity      string
			NasaDate             string
			NasaExplanation      string
			NasaHDURL            string
			NasaMediaType        string
			NasaServiceVersion   string
			NasaTitle            string
			NasaURL              string
			NasaThumbnailURL     string
			NasaCopyright        string
			NasaIdx              int
			IsNasaVideo          bool
			IsNasaImage          bool
		}
		stats := Stats{
			TotalMovies:          movCount,
			TotalTVShows:         tvCount,
			TotalVideos:          videoCount,
			MovieSizeOnDisk:      movsizeondisk,
			TVShowSizeOnDisk:     tvsizeondisk,
			VideoSizeOnDisk:      videosizeondisk,
			FreeSpaceOnDisk:      freespaceondisk,
			WeatherLocation:      wdLocation,
			WeatherTemperature:   wdTemp,
			WeatherUnit:          wdUnit,
			WeatherConditions:    wdConditions,
			WeatherWindDirection: wdWindDir,
			WeatherWindSpeed:     wdWindSpeed,
			WeatherHumidity:      wdHumidity,
			NasaDate:             nasaData.Date,
			NasaExplanation:      nasaData.Explanation,
			NasaHDURL:            nasaData.HDURL,
			NasaMediaType:        nasaData.MediaType,
			NasaServiceVersion:   nasaData.ServiceVersion,
			NasaTitle:            nasaData.Title,
			NasaURL:              nasaData.URL,
			NasaThumbnailURL:     nasaData.ThumbnailURL,
			NasaCopyright:        nasaData.Copyright,
			NasaIdx:              nasaData.Idx,
			IsNasaVideo:          nasaData.MediaType == "video",
			IsNasaImage:          nasaData.MediaType == "image",
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

// RadarPageHandler serves the NOAA/NWS radar page.
func RadarPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/radar.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, nil)
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
