

# weather-service-go

A lightweight Go weather API powered by Open-Meteo, designed with Clean Architecture and optimized for quick PoC use, easy extension, and simple REST access.

# weather-service-go

[![Build Status](https://img.shields.io/github/actions/workflow/status/yourusername/weather-service-go/ci.yml?branch=main)](https://github.com/yourusername/weather-service-go/actions)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A lightweight Go weather API powered by Open-Meteo, designed with Clean Architecture and optimized for quick PoC use, easy extension, and simple REST access.

## Table of Contents

1. [Features](#features)
2. [Project Structure](#project-structure)
3. [Getting Started](#getting-started)
4. [Usage](#usage)
5. [API Endpoints](#api-endpoints)
6. [Testing](#testing)
7. [CI/CD](#cicd)
8. [API Documentation (Swagger)](#api-documentation-swagger)
9. [Setting Up Deploy Keys](#setting-up-deploy-keys)
10. [Estimated Delivery Timeline](#estimated-delivery-timeline)
11. [Dependencies](#dependencies)
12. [FAQ](#faq)
13. [Contributing](#contributing)
14. [Contact](#contact)
15. [License](#license)

## Features
## API Endpoints

Here are some example endpoints:

- **Get current weather**
	- `GET /weather?location=London`
	- Response:
		```json
	- `GET /forecast?location=London`
	- Response:
		```

- **Location search/autocomplete**
	- `GET /location/search?query=Lon`
	- Response:
		```json
		["London", "Londonderry", "Long Beach"]
		```

See Swagger docs for full details and more endpoints.
## Dependencies

- Go 1.20+
- [swaggo/swag](https://github.com/swaggo/swag) for Swagger docs
- Other dependencies listed in `go.mod`

To update dependencies:
```sh
go get -u
go mod tidy
```
## FAQ

**Q: Why Clean Architecture?**
A: It keeps code maintainable, testable, and scalable.

**Q: How do I add a new weather provider?**
A: Implement a new service in `internal/service/` and update the usecase/repository as needed.

**Q: How do I run tests?**
A: Use `go test ./...` from the project root.

**Q: How do I regenerate Swagger docs?**
A: Run `swag init` after updating code comments.

**Q: Who do I contact for support?**
A: See the Contact section below.
## Contact

For questions, issues, or support, please open an issue on GitHub or contact [your.email@example.com](mailto:your.email@example.com).

### Core Features
- Retrieve current weather data by location
- 5-day weather forecast
- Historical weather data lookup
- Location search and autocomplete
- Support for multiple weather providers (e.g., Open-Meteo, WeatherAPI)
- Caching for frequent queries
- API authentication (e.g., API key, JWT)
- Rate limiting for API requests

### Non-Features (Project Qualities)
- Clean Architecture for maintainability and scalability
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

## Setting Up Deploy Keys

1. **Generate SSH Key Pair on Your Server:**
	 ```sh
	 ssh-keygen -t rsa -b 4096 -C "deploy-key-for-weather-service-go"
	 ```
	 - Save the key pair in a secure location (e.g., `/home/youruser/.ssh/weather-service-go_deploy`).

2. **Add the Public Key to GitHub:**
	 - Copy the contents of the public key file (e.g., `weather-service-go_deploy.pub`).
	 - Go to your repository on GitHub.
	 - Click on **Settings** > **Deploy keys** > **Add deploy key**.
	 - Give it a title and paste the public key.
	 - (Optional) Check **Allow write access** if needed.
	 - Click **Add key**.

3. **Configure Your Server to Use the Private Key:**
	 - Ensure the private key is accessible and has correct permissions:
		 ```sh
		 chmod 600 /home/youruser/.ssh/weather-service-go_deploy
		 ```
	 - Use this key when cloning or pulling from the repository:
		 ```sh
		 GIT_SSH_COMMAND='ssh -i /home/youruser/.ssh/weather-service-go_deploy' git clone git@github.com:yourusername/weather-service-go.git
		 ```

## Estimated Delivery Timeline

| Week | Tasks                                                                 |
|------|-----------------------------------------------------------------------|
| 1    | Project setup, Clean Architecture scaffolding, repo initialization    |
| 2    | Implement core entities, domain models, and basic configuration       |
| 3    | Develop weather data retrieval (current weather) and REST endpoints   |
| 4    | Add forecast, historical data, and location search features           |
| 5    | Integrate external weather providers, caching, and authentication     |
| 6    | Implement rate limiting, logging, and health check endpoints          |
| 7    | Add Swagger/OpenAPI documentation and finalize API docs               |
| 8    | Write unit and integration tests, CI/CD setup (GitHub Actions)        |
| 9    | Final review, bug fixes, and deployment preparation                   |

**Total:** ~2 months for a robust MVP  
*Adjust based on team size, feature complexity, and resource availability.*

## Contributing

Feel free to open issues or submit pull requests for improvements!

## License

MIT
