# AI Agent Instructions for weather-service-go

## Project Overview
A lightweight Go weather API (PoC) using Open-Meteo as the free weather data provider. Built with **Clean Architecture** principles: requests → handlers → service → data adapters → external APIs.

## Architecture Essentials

### Layer Responsibilities
- **`cmd/server/main.go`**: HTTP handlers only (routing, request parsing, response serialization)
- **`internal/service/`**: Business logic; implements `WeatherService` interface with dependency injection
- **`internal/config/`**: Environment variable loading with sensible defaults (PORT=9090, API_TIMEOUT_SECONDS=10)
- **`internal/entity/`**: Domain models (e.g., `CurrentWeather` struct with JSON tags for API responses)
- **`internal/openmeteo/`**: External API adapters; currently implements geocoding lookup

### Critical Data Flow for `GetCurrentWeather`
1. User calls `GET /weather?location=London` 
2. Handler extracts location, calls `ws.GetCurrentWeather(location)`
3. Service calls `openmeteo.GeocodeLocation()` to convert location name → lat/lon
4. Service calls Open-Meteo weather API with coordinates
5. Service parses weather code (0=clear, 61=rain, etc.) using `getWeatherCondition()`
6. Returns `CurrentWeather` struct; handler JSON-encodes and responds

## Key Patterns & Conventions

### Interface-Based Design
- `WeatherService` is an interface; `weatherService` is the concrete implementation
- New weather providers should implement this interface; add to service layer, not handlers
- See: [internal/service/WeatherService.go](internal/service/WeatherService.go#L9-L11)

### Configuration & Dependency Injection
- Config loaded once at startup in `init()` block; passed to service constructors
- All timeouts and URLs are configurable via environment variables (no hardcoding)
- Handler uses global `var ws` initialized in `init()` — **do not change this pattern**, it's the PoC approach

### Error Handling Convention
- Return `(*, error)` tuple; caller is responsible for wrapping with context
- Handlers write JSON error responses with appropriate HTTP status codes
- See: [cmd/server/main.go](cmd/server/main.go#L37-L42)

### Testing Strategy
- Unit tests live in `*_test.go` files in same package (e.g., [internal/service/WeatherService_test.go](internal/service/WeatherService_test.go))
- Use `httptest.NewServer()` to mock external APIs (geocoding, weather)
- Tests validate error cases (empty location) and success paths

## Build & Execution

### Development Commands
```bash
go run ./cmd/server          # Run locally on :9090
go build -o weather-server ./cmd/server  # Create binary
go test ./...                # Run all tests
go mod tidy                  # Clean dependencies
```

### Environment Variables
- `PORT`: HTTP server port (default: 9090)
- `GEOCODING_API_URL`: Open-Meteo geocoding endpoint (default: https://geocoding-api.open-meteo.com)
- `WEATHER_API_URL`: Open-Meteo forecast endpoint (default: https://api.open-meteo.com)
- `API_TIMEOUT_SECONDS`: HTTP client timeout (default: 10)

## Common Extension Points

### Adding a New Weather Endpoint (e.g., `/forecast`)
1. Create interface method in `WeatherService` (e.g., `GetForecast(location string) (*entity.Forecast, error)`)
2. Implement in `weatherService` struct — call OpenMeteo or new provider adapter
3. Add handler in `main.go` that calls `ws.GetForecast()`
4. Add unit test in `WeatherService_test.go` with mocked HTTP servers

### Adding a New Weather Provider
1. Create new adapter package (e.g., `internal/weatherapi/`) with public functions
2. Keep `WeatherService` interface unchanged; add logic inside `weatherService.GetCurrentWeather()`
3. Use config to select provider (or create new service type)
4. Example: `openmeteo.GeocodeLocation()` is an adapter to Open-Meteo's geocoding

### Fixing Broken Tests
- Mock servers in tests must match real API response structure
- Check [internal/service/WeatherService_test.go](internal/service/WeatherService_test.go#L23-L31) for pattern

## Known Issues & TODOs
- `GeocodeLocation()` in [internal/openmeteo/geocoding.go](internal/openmeteo/geocoding.go#L19) is a stub; calls `GeocodeLocationWithClient()` instead
- Middleware is defined but not wired into main.go's HTTP mux
- Global `ws` variable in main makes testing harder; consider dependency injection for future refactor

## File Structure Reference
```
cmd/server/main.go              → Handlers + HTTP setup
internal/service/WeatherService.go → Business logic (main interface)
internal/openmeteo/geocoding.go    → Adapter for geocoding API
internal/config/config.go          → Config loading from env vars
internal/entity/CurrentWeather.go  → Domain model with JSON tags
internal/middleware/logging.go     → Request/response logging (unused)
internal/service/WeatherService_test.go → Unit tests with mocked APIs
```

## Before Touching Code
- Run `go test ./...` to verify current state
- Check if environment variables are set for local testing
- Understand which layer you're modifying (handler vs. service vs. adapter)
