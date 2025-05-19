# Banking API

A comprehensive banking API built with Go and the Gin framework, featuring structured logging, error handling, and a clean architecture.

## Project Structure

```text
/
├── cmd/banking/main.go               # Main package & server initialization
├── configs/config.yml                # Application configuration (Viper)
├── deployments/docker-compose.yml    # Docker Compose for services
├── internal/
│   ├── config/
│   │   ├── app_config.go             # Configuration loader
│   │   └── logger/
│   │       └── logger.go             # Logger configuration
│   ├── controllers/customer_controller.go  # Gin HTTP handlers
│   ├── middleware/
│   │   ├── errors/
│   │   │   ├── custom_errors.go      # Custom error definitions
│   │   │   ├── error_handler.go      # Middleware error handler
│   │   │   └── error_response.go     # Error response DTO
│   │   └── logger/
│   │       └── http_logger.go        # HTTP request/response logging middleware
│   ├── models/
│   │   ├── account.go                # Account domain model & DTO
│   │   └── customer.go               # Customer domain model & DTO
│   ├── repository/
│   │   ├── account_repository.go     # Data access for accounts
│   │   └── customer_repository.go    # Data access for customers
│   ├── routes/router.go              # Gin router setup
│   └── services/customer_service.go  # Business logic layer
├── scripts/sql-scripts/
│   ├── create-tables.sql             # SQL migrations
│   └── dummy-data.sql                # Seed data
├── go.mod                            # Go module definition
├── go.sum                            # Go module checksums
├── Makefile                          # Build automation
└── README.md                         # Project documentation
```

## Development Workflow

### Using the Makefile

The project includes a Makefile to simplify common development tasks:

```bash
# Format Go code
make fmt

# Vet and build the application
make build

# Run the application
make run

# Start infrastructure (PostgreSQL database)
make infra

# Shut down infrastructure
make infra-down

# Clean up build artifacts
make clean

# Show all available commands
make help
```

For a complete list of available commands, run `make help`.

## Architecture

The application follows a layered architecture:

- **Controller Layer**: Handles HTTP requests and responses
- **Service Layer**: Contains business logic and rules
- **Repository Layer**: Manages data access
- **Middleware**: Provides cross-cutting concerns like error handling and logging

## Features

### Advanced Structured Logging

The API implements comprehensive request and response logging using Go's `slog` package:

- Automatic JSON pretty-printing for request and response bodies
- Request header logging with sensitive data redaction
- Response status and size tracking
- Content-type aware formatting

Example log output:

```
[2025-05-06T10:15:23.456-07:00] INFO: INCOMING REQUEST
  method: POST
  path: /api/v1/customers/
  query: 
  headers:
    Content-Type: application/json
    User-Agent: curl/7.81.0
  body:
    {
      "first_name": "Jane",
      "last_name": "Doe",
      "email": "jane.doe@example.com",
      "phone": "555-1234",
      "accounts": [
        {
          "account_number": "GR123456789012345678901234567",
          "balance": 1000.0,
          "account_description": "Main account",
          "currency": "EUR"
        }
      ]
    }
  client_ip: 127.0.0.1
```

### Error Handling

The application uses a centralized error handling approach:

- **Repository Layer**: Returns standard Go errors
- **Service Layer**: Transforms them into domain-specific errors
- **Controller Layer**: Delegates to middleware
- **Error Middleware**: Formats consistent HTTP responses

## API Endpoints

| Method | Endpoint                   | Description                               |
|--------|----------------------------|-------------------------------------------|
| GET    | /api/v1/customers          | Retrieve all customers                    |
| GET    | /api/v1/customers/:email   | Retrieve customer by email with accounts  |
| POST   | /api/v1/customers          | Create a new customer                     |
| DELETE | /api/v1/customers/:email   | Delete customer by email                  |

## Configuration with Viper

The application uses [Viper](https://github.com/spf13/viper) for configuration management. Viper allows the application
to read from various configuration sources, such as files, environment variables, and more.

Example configuration:

```yaml
# configs/config.yml
database:
  host: "localhost"
  port: 5432
  user: "your_username"
  password: "your_password"
  dbname: "your_dbname"

server:
  port: 8080
  mode: "debug"
```

## Example HTTP Requests

```http
# Get all customers
GET http://localhost:8080/api/v1/customers

# Get customer by email
GET http://localhost:8080/api/v1/customers/jane.doe@example.com

# Create a new customer
POST http://localhost:8080/api/v1/customers
Content-Type: application/json

{
  "firstName": "Jane",
  "lastName": "Doe",
  "email": "jane.doe@example.com",
  "phone": "555-1234",
  "accounts": [
    {
      "account_number": "GR123456789012345678901234567",
      "balance": 1000.0,
      "account_description": "Main account",
      "currency": "EUR"
    }
  ]
}

# Delete customer by email
DELETE http://localhost:8080/api/v1/customers/jane.doe@example.com
```

## Running the Application

### Setup the database using Docker Compose:
```bash
docker-compose -f deployments/docker-compose.yml up -d
```

### Run the application:
```bash
go run cmd/banking/main.go 
```

The server runs on http://localhost:8080

## Next Steps

### Implemented
- ✅ **Advanced Structured Logging**
  - Structured request and response logging with JSON formatting
  - Sensitive data redaction
  - Content-type aware formatting

### Future Improvements
- **Dependency Injection in Go**
  - [Google Wire - Compile-time DI](https://github.com/google/wire)
  - [Uber Dig - Runtime DI](https://github.com/uber-go/dig)

- **ORM Integration**
  - [GORM](https://gorm.io/) 

- **Graceful Shutdown**
  - Implement proper server shutdown handling

- **Authentication & Authorization**
  - Add JWT-based authentication
  - Role-based access control

- **API Documentation**
  - Integrate Swagger/OpenAPI documentation