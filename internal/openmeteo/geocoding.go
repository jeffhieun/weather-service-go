package openmeteo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/jeffhieun/weather-service-go/internal/config"
)

type GeocodingResponse struct {
	Results []struct {
		Name      string  `json:"name"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"results"`
}

func GeocodeLocation(location string) (float64, float64, error) {
	return 0, 0, fmt.Errorf("location cannot be empty")
}
	
func GeocodeLocationWithClient(location string, client *http.Client, cfg *config.Config) (float64, float64, error) {
    if location == "" {
        return 0, 0, fmt.Errorf("location cannot be empty")
    }

    urlStr := fmt.Sprintf("%s/v1/search?name=%s&count=1", cfg.GeocodingAPIURL, url.QueryEscape(location))
    resp, err := client.Get(urlStr)
    if err != nil {
        return 0, 0, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return 0, 0, fmt.Errorf("geocoding API returned status code: %d", resp.StatusCode)
    }

    var geoResp GeocodingResponse
    if err := json.NewDecoder(resp.Body).Decode(&geoResp); err != nil {
        return 0, 0, err
    }

    if len(geoResp.Results) == 0 {
        return 0, 0, fmt.Errorf("no results found for location: %s", location)
    }

    return geoResp.Results[0].Latitude, geoResp.Results[0].Longitude, nil
}