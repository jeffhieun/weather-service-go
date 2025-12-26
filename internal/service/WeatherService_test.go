package service

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jeffhieun/weather-service-go/internal/config"
)

func TestGetCurrentWeather_EmptyLocation(t *testing.T) {
    cfg := &config.Config{APITimeoutSeconds: 10}	
    ws := NewWeatherService(cfg)

    weather, err := ws.GetCurrentWeather("")
    if err == nil {
        t.Error("expected error for empty location, got nil")
    }
    if weather != nil {
        t.Error("expected nil weather for empty location")
    }
}

func TestGetCurrentWeather_Success(t *testing.T) {
    // Mock geocoding server
    geoServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{
            "results": [{
                "name": "London",
                "latitude": 51.5085,
                "longitude": -0.1257
            }]
        }`))
    }))
    defer geoServer.Close()

    // Mock weather server
    weatherServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{
            "current": {
                "temperature_2m": 15.5,
                "relative_humidity_2m": 72,
                "weather_code": 0
            }
        }`))
    }))

    defer weatherServer.Close()

    cfg := &config.Config{
        APITimeoutSeconds: 10,
        GeocodingAPIURL:   geoServer.URL,
        WeatherAPIURL:     weatherServer.URL,
    }
    ws := NewWeatherService(cfg)
    _ = ws
}

func TestGetWeatherCondition(t *testing.T) {
    tests := []struct {
        code     int
        expected string
    }{
        {0, "Clear sky"},
        {1, "Partly cloudy"},
        {3, "Overcast"},
        {45, "Foggy"},
        {61, "Rain"},
        {95, "Thunderstorm"},
        {999, "Unknown"},
    }

    for _, tt := range tests {
        t.Run(tt.expected, func(t *testing.T) {
            result := getWeatherCondition(tt.code)
            if result != tt.expected {
                t.Errorf("code %d: expected %s, got %s", tt.code, tt.expected, result)
            }
        })
    }
}