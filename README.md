# Financial API in Go

A RESTful API service built with Go, following clean architecture principles.

## Project Structure

```
.
├── cmd/
│   └── api/          # Application entry point
├── internal/         # Private application code
│   ├── handlers/     # HTTP request handlers
│   ├── middleware/   # HTTP middleware functions
│   ├── models/       # Data models/structs
│   ├── repository/   # Database interaction layer
│   └── service/      # Business logic layer
├── pkg/             # Public libraries
│   ├── config/      # Configuration management
│   └── utils/       # Utility functions
└── api/             # API documentation
    └── docs/        # OpenAPI/Swagger specs
```

## Getting Started

### Prerequisites

- Go 1.21 or later

### Running the Application

1. Clone the repository
2. Run the application:
   ```bash
   go run cmd/api/main.go
   ```

The server will start on port 8080.

### API Endpoints

- `GET /health` - Health check endpoint

## Development

### Adding Dependencies

```bash
go get github.com/example/package
```

### Running Tests

```bash
go test ./...
``` 