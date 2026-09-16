package server

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

func seedWeatherCacheForTest(t *testing.T, snapshot WeatherSnapshot, fetchedAt time.Time, valid bool) {
	seedWeatherCacheForLocationForTest(t, defaultWeatherLocationKey, snapshot, fetchedAt, valid)
}

func seedWeatherCacheForLocationForTest(t *testing.T, locationKey string, snapshot WeatherSnapshot, fetchedAt time.Time, valid bool) {
	t.Helper()
	weatherCacheMu.Lock()
	previous := weatherCacheByKey[locationKey]
	weatherCacheByKey[locationKey] = &weatherCacheEntry{data: snapshot, fetchedAt: fetchedAt, valid: valid}
	weatherCacheMu.Unlock()

	t.Cleanup(func() {
		weatherCacheMu.Lock()
		if previous == nil {
			delete(weatherCacheByKey, locationKey)
		} else {
			weatherCacheByKey[locationKey] = previous
		}
		weatherCacheMu.Unlock()
	})
}

func TestBuildWeatherTemplateData_UsesCachedSnapshot(t *testing.T) {
	seedWeatherCacheForTest(t, WeatherSnapshot{
		Location:      "Belfair, WA",
		Temperature:   "62",
		Unit:          "F",
		Conditions:    "Clear",
		WindDirection: "NW",
		WindSpeed:     "5 mph",
	}, time.Now(), true)

	got, err := buildWeatherTemplateData(nil)
	if err != nil {
		t.Fatalf("buildWeatherTemplateData returned error: %v", err)
	}

	if got.WeatherLocation != "Belfair, WA" {
		t.Fatalf("expected WeatherLocation Belfair, WA, got %q", got.WeatherLocation)
	}
	if got.WeatherTemperature != "62" || got.WeatherUnit != "F" {
		t.Fatalf("expected temperature 62 F, got %q %q", got.WeatherTemperature, got.WeatherUnit)
	}
	if got.WeatherConditions != "Clear" {
		t.Fatalf("expected conditions Clear, got %q", got.WeatherConditions)
	}
}

func TestRadarPageHandler_RendersWeatherDivValues(t *testing.T) {
	originalWD, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd failed: %v", err)
	}
	if err := os.Chdir(".."); err != nil {
		t.Fatalf("chdir failed: %v", err)
	}
	t.Cleanup(func() {
		_ = os.Chdir(originalWD)
	})

	seedWeatherCacheForTest(t, WeatherSnapshot{
		Location:      "Belfair, WA",
		Temperature:   "59",
		Unit:          "F",
		Conditions:    "Cloudy",
		WindDirection: "S",
		WindSpeed:     "7 mph",
	}, time.Now(), true)

	req := httptest.NewRequest(http.MethodGet, "/radar", nil)
	rr := httptest.NewRecorder()

	RadarPageHandler(nil).ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d body=%s", rr.Code, rr.Body.String())
	}

	body := rr.Body.String()
	if !strings.Contains(body, "class=\"weather-div\"") {
		t.Fatalf("expected weather-div markup in radar page")
	}
	if !strings.Contains(body, "Belfair, WA") {
		t.Fatalf("expected location text in radar page")
	}
	if !strings.Contains(body, "59 F") {
		t.Fatalf("expected temperature text in radar page")
	}
	if !strings.Contains(body, "Cloudy") {
		t.Fatalf("expected conditions text in radar page")
	}
}

func TestWeatherAPIHandler_ReturnsCachedSnapshotForLocation(t *testing.T) {
	seedWeatherCacheForLocationForTest(t, "durant", WeatherSnapshot{
		Location:      "Durant, OK",
		Temperature:   "88",
		Unit:          "F",
		Conditions:    "Sunny",
		WindDirection: "SW",
		WindSpeed:     "10 mph",
	}, time.Now(), true)

	req := httptest.NewRequest(http.MethodGet, "/api/weather?location=durant", nil)
	rr := httptest.NewRecorder()

	WeatherAPIHandler(nil).ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d body=%s", rr.Code, rr.Body.String())
	}
	body := rr.Body.String()
	if !strings.Contains(body, "Durant, OK") || !strings.Contains(body, "Sunny") {
		t.Fatalf("expected Durant weather data in response, got %s", body)
	}
}

func TestWeatherAPIHandler_UnknownLocationReturnsBadRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/weather?location=nowhere", nil)
	rr := httptest.NewRecorder()

	WeatherAPIHandler(nil).ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rr.Code)
	}
}
