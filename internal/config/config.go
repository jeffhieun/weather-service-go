package config

import (
    "log"
    "os"
    "strconv"
    "time"
)

type Config struct {
    Port              string
    GeocodingAPIURL   string
    WeatherAPIURL     string
    APITimeoutSeconds int
}

func LoadConfig() *Config {
	
    cfg := &Config{
        Port:              getEnv("PORT", "9090"),
        GeocodingAPIURL:   getEnv("GEOCODING_API_URL", "https://geocoding-api.open-meteo.com"),
        WeatherAPIURL:     getEnv("WEATHER_API_URL", "https://api.open-meteo.com"),
        APITimeoutSeconds: getEnvInt("API_TIMEOUT_SECONDS", 10),		
    }

    log.Printf("Config loaded: Port=%s, Timeout=%ds", cfg.Port, cfg.APITimeoutSeconds)
    return cfg
}

func (c *Config) GetAPITimeout() time.Duration {
    return time.Duration(c.APITimeoutSeconds) * time.Second
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
    if value := os.Getenv(key); value != "" {
        if intVal, err := strconv.Atoi(value); err == nil {
            return intVal
        }
    }
    return defaultValue
}