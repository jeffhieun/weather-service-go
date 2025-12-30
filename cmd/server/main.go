package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jeffhieun/weather-service-go/internal/config"
    "github.com/jeffhieun/weather-service-go/internal/middleware"
    "github.com/jeffhieun/weather-service-go/internal/service"
)

var ws service.WeatherService

func init() {
	cfg := config.LoadConfig()
	ws = service.NewWeatherService(cfg)
}

func main() {
	cfg := config.LoadConfig()

	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/weather", weatherHandler)
	
	handler := middleware.LoggingMiddleware(mux)
	log.Printf("Starting server on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, handler))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, `{"message": "Welcome to the Weather Service! Use /weather?location=YOUR_CITY to get weather information."}`)
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	location := r.URL.Query().Get("location")
	if location == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "location query parameter is required"})
		return
	}
	
	weather, err := ws.GetCurrentWeather(location)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(weather)
}