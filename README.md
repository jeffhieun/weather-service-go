
# weather-service-go

A lightweight Go weather API powered by Open-Meteo, designed with Clean Architecture and optimized for quick PoC use, easy extension, and simple REST access.

## Features

- Clean Architecture for maintainability and scalability
- RESTful API endpoints for weather data
- Easy integration with external weather services (Open-Meteo)
- Simple configuration and extension
- Ready for unit testing
- CI/CD ready (GitHub Actions example)
- Swagger/OpenAPI documentation support

## Project Structure

Below is an explanation of each main folder in this project:

- **cmd/**  
	Contains the application's entry point. Typically, you place the main executable code here (e.g., `main.go`).

- **internal/**  
	Holds the core business logic and application code. This folder is protected from external imports and is divided into several subfolders:
	- **config/**: Configuration files and logic (e.g., environment variables, app settings, security-related configuration).
	- **domain/**: Domain models and business entities.
	- **usecase/**: Application use cases, orchestrating business rules.
	- **repository/**: Data access logic, interfaces, and implementations for persistence. Add caching logic here if needed.
	- **handler/**: HTTP handlers, controllers for routing and request processing. Implement authentication, authorization, input validation, and Swagger handler here.
	- **service/**: Integration with external services (e.g., third-party APIs, external security or caching services).
	- **entity/**: Definitions of core entities used throughout the application.

- **pkg/**  
	Shared libraries and utilities that can be reused across different parts of the project or even in other projects. Place reusable cache utilities here.

- **docs/**  
	Contains Swagger/OpenAPI specification files and auto-generated API documentation.

- **.github/workflows/**  
	Contains CI/CD workflow files for GitHub Actions.

This structure supports maintainability, scalability, security, caching, and API documentation by organizing related logic into appropriate folders.

## Getting Started

### Prerequisites

- Go 1.20+
- Internet connection (for external API calls)

### Installation

```sh
git clone https://github.com/yourusername/weather-service-go.git
cd weather-service-go
go mod tidy
```

### Running the Service

```sh
go run cmd/app/main.go
```

## Usage

- Access the API at `http://localhost:8080/weather?location=YOUR_LOCATION`
- Access Swagger API docs at `http://localhost:8080/swagger/index.html`
- See API documentation for available endpoints.

## Testing

Unit tests are placed alongside the code they test, with filenames ending in `_test.go`.

Run all unit tests:

```sh
go test ./...
```

## CI/CD

GitHub Actions workflows are configured in `.github/workflows/`.  
Public repositories are free; private repositories have a free tier with limited minutes.

## API Documentation (Swagger)

- Swagger/OpenAPI documentation is available in the `docs/` folder.
- You can auto-generate docs using [swaggo/swag](https://github.com/swaggo/swag).
- Serve Swagger UI via a handler in `internal/handler/`.

## Contributing

Feel free to open issues or submit pull requests for improvements!

## License

MIT
