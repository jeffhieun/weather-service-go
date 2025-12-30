package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/jeffhieun/weather-service-go/internal/config"
	"github.com/jeffhieun/weather-service-go/internal/entity"
	"github.com/jeffhieun/weather-service-go/internal/openmeteo"
)

type WeatherService interface {
	GetCurrentWeather(location string) (*entity.CurrentWeather, error)	
}

type weatherService struct {
    client *http.Client
    cfg    *config.Config
}

type weatherResponse struct {
	Current struct {
		Temperature float64 `json:"temperature_2m"`
		Humidity    int     `json:"relative_humidity_2m"`
		WeatherCode int     `json:"weather_code"`
	} `json:"current"`
}

func NewWeatherService(cfg *config.Config) WeatherService {
	 return &weatherService{
        client: &http.Client{Timeout: cfg.GetAPITimeout()},
        cfg:    cfg,
    }
}

func (ws *weatherService) GetCurrentWeather(location string) (*entity.CurrentWeather, error) {
	if location == "" {
		return nil, fmt.Errorf("location cannot be empty")
	}
	
	lat, lon, err := openmeteo.GeocodeLocation(location, ws.client, ws.cfg)
	if err != nil {
		return nil, err
	}
		
	urlStr := fmt.Sprintf("%s/v1/forecast?latitude=%.4f&longitude=%.4f&current=temperature_2m,relative_humidity_2m,weather_code", ws.cfg.WeatherAPIURL, lat, lon)
	resp, err := ws.client.Get(urlStr)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("weather API returned status code: %d", resp.StatusCode)
	}
	
	var weatherResp weatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResp); err != nil {
		return nil, err
	}
	
	condition := getWeatherCondition(weatherResp.Current.WeatherCode)
	
	return &entity.CurrentWeather{
		Location:    location,
		Temperature: weatherResp.Current.Temperature,
		Humidity:    weatherResp.Current.Humidity,
		Condition:   condition,
	}, nil
}

func getWeatherCondition(code int) string {
	switch code {
	case 0:
		return "Clear sky"
	case 1, 2:
		return "Partly cloudy"
	case 3:
		return "Overcast"
	case 45, 48:
		return "Foggy"
	case 51, 53, 55:
		return "Drizzle"
	case 61, 63, 65:
		return "Rain"
	case 71, 73, 75:
		return "Snow"
	case 80, 81, 82:
		return "Rain showers"
	case 85, 86:
		return "Snow showers"
	case 95, 96, 99:
		return "Thunderstorm"
	default:
		return "Unknown"
	}
}
