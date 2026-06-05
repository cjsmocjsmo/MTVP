package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"sync"
	"syscall"
	"time"
	// "mtvp/setup"
)

const (
	weatherCacheTTL        = 10 * time.Minute
	weatherHTTPTimeout     = 2 * time.Second
	nasaRefreshMinInterval = 5 * time.Minute
)

type WeatherSnapshot struct {
	Location      string
	Temperature   string
	Unit          string
	Conditions    string
	WindDirection string
	WindSpeed     string
	Humidity      string
}

var (
	indexTemplateOnce sync.Once
	indexTemplate     *template.Template
	indexTemplateErr  error

	weatherCacheMu sync.RWMutex
	weatherCache   struct {
		data      WeatherSnapshot
		fetchedAt time.Time
		valid     bool
	}

	nasaRefreshMu          sync.Mutex
	nasaRefreshInFlight    bool
	nasaLastRefreshAttempt time.Time
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
	return MTVWeatherWithTimeout(db, weatherHTTPTimeout)
}

func MTVWeatherWithTimeout(db *sql.DB, timeout time.Duration) ([]byte, error) {
	// Fetch weather for Belfair, WA from National Weather Service
	latitude := 47.4281
	longitude := -122.8189
	client := &http.Client{Timeout: timeout}
	pointURL := fmt.Sprintf("https://api.weather.gov/points/%f,%f", latitude, longitude)
	pointResp, err := client.Get(pointURL)
	if err != nil {
		return nil, fmt.Errorf("weather fetch failed: %v", err)
	}
	defer pointResp.Body.Close()
	if pointResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("weather point status: %d", pointResp.StatusCode)
	}
	var pointObj map[string]interface{}
	if err := json.NewDecoder(pointResp.Body).Decode(&pointObj); err != nil {
		return nil, fmt.Errorf("weather parse failed: %v", err)
	}
	properties, ok := pointObj["properties"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("weather properties missing")
	}
	forecastURL, ok := properties["forecastHourly"].(string)
	if !ok {
		return nil, fmt.Errorf("weather forecast URL missing")
	}
	weatherResp, err := client.Get(forecastURL)
	if err != nil {
		return nil, fmt.Errorf("weather fetch failed: %v", err)
	}
	defer weatherResp.Body.Close()
	if weatherResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("weather forecast status: %d", weatherResp.StatusCode)
	}
	var weatherObj map[string]interface{}
	if err := json.NewDecoder(weatherResp.Body).Decode(&weatherObj); err != nil {
		return nil, fmt.Errorf("weather parse failed: %v", err)
	}
	weatherProps, ok := weatherObj["properties"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("weather properties missing")
	}
	periods, ok := weatherProps["periods"].([]interface{})
	if !ok || len(periods) == 0 {
		return nil, fmt.Errorf("weather no periods data")
	}
	current, ok := periods[0].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("weather current period invalid")
	}

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

func decodeWeatherSnapshot(weatherData []byte) (WeatherSnapshot, error) {
	var weatherMap map[string]interface{}
	if err := json.Unmarshal(weatherData, &weatherMap); err != nil {
		return WeatherSnapshot{}, err
	}
	s := WeatherSnapshot{}
	if v, ok := weatherMap["location"].(string); ok {
		s.Location = v
	}
	if v, ok := weatherMap["temperature"].(float64); ok {
		s.Temperature = fmt.Sprintf("%.0f", v)
	} else if v, ok := weatherMap["temperature"].(string); ok {
		s.Temperature = v
	}
	if v, ok := weatherMap["temperature_unit"].(string); ok {
		s.Unit = v
	}
	if v, ok := weatherMap["conditions"].(string); ok {
		s.Conditions = v
	}
	if v, ok := weatherMap["winddirection"].(string); ok {
		s.WindDirection = v
	}
	if v, ok := weatherMap["windspeed"].(string); ok {
		s.WindSpeed = v
	}
	if v, ok := weatherMap["humidity"].(float64); ok {
		s.Humidity = fmt.Sprintf("%.0f", v)
	} else if v, ok := weatherMap["humidity"].(string); ok {
		s.Humidity = v
	}
	return s, nil
}

func getWeatherSnapshotCached(db *sql.DB) (WeatherSnapshot, error) {
	now := time.Now()
	weatherCacheMu.RLock()
	if weatherCache.valid && now.Sub(weatherCache.fetchedAt) < weatherCacheTTL {
		cached := weatherCache.data
		weatherCacheMu.RUnlock()
		return cached, nil
	}
	hasStale := weatherCache.valid
	stale := weatherCache.data
	weatherCacheMu.RUnlock()

	weatherData, err := MTVWeatherWithTimeout(db, weatherHTTPTimeout)
	if err != nil {
		if hasStale {
			return stale, err
		}
		return WeatherSnapshot{}, err
	}

	snapshot, err := decodeWeatherSnapshot(weatherData)
	if err != nil {
		if hasStale {
			return stale, err
		}
		return WeatherSnapshot{}, err
	}

	weatherCacheMu.Lock()
	weatherCache.data = snapshot
	weatherCache.fetchedAt = now
	weatherCache.valid = true
	weatherCacheMu.Unlock()

	return snapshot, nil
}

func getLatestNASAData(db *sql.DB) (*APODResponse, error) {
	nasaData := &APODResponse{}
	fallbackRow := db.QueryRow(`SELECT Date, Explanation, HDURL, MediaType, ServiceVersion, Title, URL, ThumbnailURL, Copyright, Idx FROM nasa ORDER BY Date DESC, Idx DESC LIMIT 1`)
	if err := fallbackRow.Scan(
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
	); err != nil {
		return nil, err
	}
	return nasaData, nil
}

func maybeRefreshNASAAsync(db *sql.DB) {
	now := time.Now()
	nasaRefreshMu.Lock()
	if nasaRefreshInFlight || now.Sub(nasaLastRefreshAttempt) < nasaRefreshMinInterval {
		nasaRefreshMu.Unlock()
		return
	}
	nasaRefreshInFlight = true
	nasaLastRefreshAttempt = now
	nasaRefreshMu.Unlock()

	go func() {
		start := time.Now()
		if _, err := FetchNASAData(db); err != nil {
			log.Printf("[HomePageHandler] NASA background refresh failed after %s: %v", time.Since(start), err)
		} else {
			log.Printf("[HomePageHandler] NASA background refresh completed in %s", time.Since(start))
		}
		nasaRefreshMu.Lock()
		nasaRefreshInFlight = false
		nasaRefreshMu.Unlock()
	}()
}

func getIndexTemplate() (*template.Template, error) {
	indexTemplateOnce.Do(func() {
		indexTemplate, indexTemplateErr = template.ParseFiles("templates/index.html")
	})
	return indexTemplate, indexTemplateErr
}

// PerfHealthHandler exposes lightweight runtime state for homepage performance diagnostics.
func PerfHealthHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()

		weatherCacheMu.RLock()
		weatherValid := weatherCache.valid
		weatherFetchedAt := weatherCache.fetchedAt
		weatherLocation := weatherCache.data.Location
		weatherCacheMu.RUnlock()

		nasaRefreshMu.Lock()
		nasaInFlight := nasaRefreshInFlight
		nasaLastAttempt := nasaLastRefreshAttempt
		nasaRefreshMu.Unlock()

		latestNasaDate := ""
		latestNasaTitle := ""
		latestNasaErr := ""
		if nasaData, err := getLatestNASAData(db); err == nil {
			latestNasaDate = nasaData.Date
			latestNasaTitle = nasaData.Title
		} else {
			latestNasaErr = err.Error()
		}

		weatherAgeSeconds := -1
		if weatherValid {
			weatherAgeSeconds = int(now.Sub(weatherFetchedAt).Seconds())
		}

		nasaLastAttemptAgeSeconds := -1
		if !nasaLastAttempt.IsZero() {
			nasaLastAttemptAgeSeconds = int(now.Sub(nasaLastAttempt).Seconds())
		}

		payload := map[string]interface{}{
			"now":                             now.Format(time.RFC3339),
			"weather_cache_ttl_seconds":       int(weatherCacheTTL.Seconds()),
			"weather_cache_valid":             weatherValid,
			"weather_cache_fetched_at":        weatherFetchedAt.Format(time.RFC3339),
			"weather_cache_age_seconds":       weatherAgeSeconds,
			"weather_cached_location":         weatherLocation,
			"weather_http_timeout_seconds":    int(weatherHTTPTimeout.Seconds()),
			"nasa_refresh_in_flight":          nasaInFlight,
			"nasa_last_refresh_attempt":       nasaLastAttempt.Format(time.RFC3339),
			"nasa_last_attempt_age_seconds":   nasaLastAttemptAgeSeconds,
			"nasa_refresh_min_interval_seconds": int(nasaRefreshMinInterval.Seconds()),
			"latest_nasa_date":                latestNasaDate,
			"latest_nasa_title":               latestNasaTitle,
			"latest_nasa_error":               latestNasaErr,
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(payload); err != nil {
			http.Error(w, "Failed to encode perf health response", http.StatusInternalServerError)
		}
	}
}

// HomePageHandler serves the index.html page for the root path
func HomePageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestStart := time.Now()

		statsStart := time.Now()
		movCount := getMovieCount(db)
		tvCount := getTVShowCount(db)
		videoCount := getVideoCount(db)
		movsizeondisk := getMoviesSizeOnDisk(db)
		tvsizeondisk := getTVShowsSizeOnDisk(db)
		videosizeondisk := getVideosSizeOnDisk(db)
		freespaceondisk := freeSpaceOnDisk("/")
		log.Printf("[HomePageHandler] local stats computed in %s", time.Since(statsStart))

		weatherStart := time.Now()
		weatherSnapshot, weatherErr := getWeatherSnapshotCached(db)
		if weatherErr != nil {
			log.Printf("[HomePageHandler] weather fetch/cache warning after %s: %v", time.Since(weatherStart), weatherErr)
		}
		log.Printf("[HomePageHandler] weather stage completed in %s", time.Since(weatherStart))

		nasaStart := time.Now()
		nasaData, nasaErr := getLatestNASAData(db)
		if nasaErr != nil {
			log.Printf("[HomePageHandler] no NASA row available: %v", nasaErr)
			nasaData = &APODResponse{}
		}
		maybeRefreshNASAAsync(db)
		log.Printf("[HomePageHandler] nasa stage completed in %s", time.Since(nasaStart))
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
			WeatherLocation:      weatherSnapshot.Location,
			WeatherTemperature:   weatherSnapshot.Temperature,
			WeatherUnit:          weatherSnapshot.Unit,
			WeatherConditions:    weatherSnapshot.Conditions,
			WeatherWindDirection: weatherSnapshot.WindDirection,
			WeatherWindSpeed:     weatherSnapshot.WindSpeed,
			WeatherHumidity:      weatherSnapshot.Humidity,
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

		tmplStart := time.Now()
		tmpl, err := getIndexTemplate()
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		log.Printf("[HomePageHandler] template load in %s", time.Since(tmplStart))

		execStart := time.Now()
		err = tmpl.Execute(w, stats)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		log.Printf("[HomePageHandler] template execute in %s, total request %s", time.Since(execStart), time.Since(requestStart))
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
