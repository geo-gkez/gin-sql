### Banking API

A simple banking API built with Go and the Gin framework.

#### Project Structure

```
.                       # Root of the project
├── cmd               # Application entrypoint(s)
│   └── banking
│       └── main.go   # Main package with server initialization
  ├── configs           # Configuration files
│   └── config.yml    # Viper configuration
├── api               # HTTP request definitions for REST client
  │   └── customer.http # Sample API tests/requests
├── deployments       # Infrastructure setup scripts
│   └── docker-compose.yml # Docker Compose for database/service
├── internal          # Application code (non-exported)
│   ├── config        # App configuration logic
│   │   └── app_config.go
│   ├── controllers   # HTTP handler layer (Gin controllers)
│   │   └── customer_controller.go
│   ├── middleware    # Middleware for cross-cutting concerns
│   │   └── errors    # Error handling middleware
      │   │       ├── custom_errors.go
│   │       ├── error_handler.go
│   │       └── error_response.go
│   ├── models        # Domain models and DTOs
│   │   ├── account.go
│   │   └── customer.go
│   ├── repository    # Data access layer
│   │   ├── account_repository.go
│   │   └── customer_repository.go
│   ├── routes        # Gin router setup
    │   │   └── router.go
│   └── services      # Business logic layer
│       └── customer_service.go
├── scripts           # Utility scripts
│   └── sql-scripts   # SQL migrations and seed data
    │       ├── create-tables.sql
│       └── dummy-data.sql
├── go.mod            # Go module definition
├── go.sum            # Go module checksums
└── README.md         # Project documentation
```

#### Architecture

The application follows a layered architecture:

- **Controller Layer**: Handles HTTP requests and responses
- **Service Layer**: Contains business logic and rules
- **Repository Layer**: Manages data access
- **Middleware**: Provides cross-cutting concerns like error handling

#### API Endpoints

| Method | Endpoint                   | Description                               |
|--------|----------------------------|-------------------------------------------|
| GET    | /api/v1/customers          | Retrieve all customers                    |
| GET    | /api/v1/customers/:email   | Retrieve customer by email with accounts  |
| POST   | /api/v1/customers          | Create a new customer                     |
| DELETE | /api/v1/customers/:email   | Delete customer by email                  |

#### Configuration with Viper

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


#### Example HTTP Requests

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
  "phone": "555-1234"
}

# Delete customer by email
DELETE http://localhost:8080/api/v1/customers/jane.doe@example.com
```

#### Running the Application

#### Docker Compose is provided for easy database setup:
```bash
    docker-compose -f deployments/docker-compose.yml up -d
```


```bash
   go run cmd/banking/main.go 
```

The server runs on http://localhost:8080

#### Error Handling

The application uses a centralized error handling approach:

- **Repository Layer**: Returns standard Go errors
- **Service Layer**: Transforms them into domain-specific errors
- **Controller Layer**: Delegates to middleware
  - **Error Middleware**: Formats consistent HTTP responses

    #### Next Steps:

- **Dependency Injection in Go**
    - [Google Wire - Compile-time DI](https://github.com/google/wire)
    - [Uber Dig - Runtime DI](https://github.com/uber-go/dig)
    - [Article: Dependency Injection in Go](https://medium.com/@john.doe/dependency-injection-in-go-1a6a1f4eabe)
    - [Go Dependency Injection with Wire](https://medium.com/@john.doe/go-dependency-injection-with-wire-1a6a1f4eabe)

- **Use ORM library**
  - [GORM](https://gorm.io/) 

- **Graceful shutdown**
  - Find a way to gracefully shutdown the server