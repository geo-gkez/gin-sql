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

| Method | Endpoint                   | Description                        |
|--------|----------------------------|------------------------------------|
| GET    | /api/v1/customers          | Retrieve all customers             |
| GET    | /api/v1/customers/:email   | Retrieve customer by email with accounts |
| POST   | /api/v1/customers          | Create a new customer              |

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

Example usage in the application:

```go
import (
"fmt"
"github.com/gin-gonic/gin"
"github.com/spf13/viper"
"org/gg/banking/internal/controllers"
"org/gg/banking/internal/repository"
"org/gg/banking/internal/routes"
"org/gg/banking/internal/services"
"os"
)

func SetupApp() (*gin.Engine, ServerConfiguration) {
config, err := LoadConfig()
if err != nil {
panic(fmt.Sprintf("failed to load config: %v", err))
}
// Create components
repo := repository.NewCustomerRepository()
service := services.NewCustomerService(repo)
controller := controllers.NewCustomerController(service)

// Setup router and routes
router := routes.SetupRouter()
routes.RegisterRoutes(router, controller)

return router, config.Server
}

type AppConfiguration struct {
Database DatabaseConfiguration
Server   ServerConfiguration
}

type DatabaseConfiguration struct {
Host     string
Port     int
User     string
Password string
DBName   string
}

type ServerConfiguration struct {
Port int
Mode string
}

func LoadConfig() (*AppConfiguration, error) {
var env = os.Getenv("ENV")

viper.SetConfigName(env)
viper.SetConfigName("config")
viper.SetConfigType("yml")

// Get the project root directory
projectRoot, err := os.Getwd()
if err != nil {
return nil, fmt.Errorf("failed to get working directory: %w", err)
}

// Try multiple possible config locations
viper.AddConfigPath(fmt.Sprintf("%s/configs", projectRoot)) // From project root
viper.AddConfigPath("configs")       // Direct subfolder
viper.AddConfigPath("../../configs") // Two levels up

if err := viper.ReadInConfig(); err != nil {
return nil, fmt.Errorf("failed to read config file: %w", err)
}

var config AppConfiguration
if err := viper.Unmarshal(&config); err != nil {
return nil, fmt.Errorf("failed to unmarshal config: %w", err)
}

return &config, nil
}
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