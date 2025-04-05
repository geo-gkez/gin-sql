### Banking API
A simple banking API built with Go and the Gin framework.

#### Project Structure
```
├── cmd/
│   └── banking/                # Main application entry point
│       └── main.go
├── internal/                   # Private application code
│   ├── config/                 # Configuration logic
│   ├── controllers/            # HTTP handlers
│   ├── middleware/             # HTTP middleware
│   ├── models/                 # Data models
│   ├── repository/             # Data access layer
│   ├── services/               # Business logic
│   └── routes/                 # API route definitions
├── pkg/                        # Public reusable libraries
├── configs/                    # Configuration files
│   └── config.yml              # YAML configuration
├── api/                        # API definitions
│   └── http_requests/          # HTTP request examples
├── deployments/                # Deployment configurations
│   └── docker-compose.yml
└── scripts/                    # Utility scripts
    └── sql-scripts/            # SQL initialization scripts
```

#### Architecture
The application follows a layered architecture:
- **Controller Layer**: Handles HTTP requests and responses
- **Service Layer**: Contains business logic and rules
- **Repository Layer**: Manages data access
- **Middleware**: Provides cross-cutting concerns like error handling

#### API Endpoints

| Method | Endpoint   | Description            |
|--------|------------|------------------------|
| GET    | /customers | Retrieve all customers |

#### Running the Application

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

#### TODO
**Database Integration**
- [ ] Add database configuration
- [ ] Implement database connection pooling
- [ ] Update repositories to use the database
- [ ] Add migrations system

**Dependency Injection Improvements**
- [ ] Consider using a DI framework like Wire or Dig
- [ ] Support different environments (dev, test, prod)
- [ ] Add configuration through environment variables

#### Resources
- **Dependency Injection in Go**
    - [Google Wire - Compile-time DI](https://github.com/google/wire)
    - [Uber Dig - Runtime DI](https://github.com/uber-go/dig)
    - [Article: Dependency Injection in Go](https://medium.com/@john.doe/dependency-injection-in-go-1a6a1f4eabe)
    - [Go Dependency Injection with Wire](https://medium.com/@john.doe/go-dependency-injection-with-wire-1a6a1f4eabe)