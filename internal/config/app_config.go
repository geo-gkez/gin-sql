package config

import (
	"database/sql"
	"fmt"
	"log"
	"org/gg/banking/internal/config/logger"
	"org/gg/banking/internal/controllers"
	"org/gg/banking/internal/repository"
	"org/gg/banking/internal/routes"
	"org/gg/banking/internal/services"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func SetupApp() {
	config, err := LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}

	// config.Server.Mode
	// initialize logger
	logger.InitLogger(logger.LoggerConfig{
		Level: config.Server.Mode,
	})

	db := connectToPostgres(config)
	// Ensure DB is closed when application terminates
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println("Error closing database connection:", err)
		} else {
			log.Println("Database connection closed")
		}
	}(db)

	// Create components
	customerRepository := repository.NewCustomerRepository(db)
	accountRepository := repository.NewAccountRepository(db)
	customerService := services.NewCustomerService(customerRepository, accountRepository)
	customerController := controllers.NewCustomerController(customerService)

	// Setup router and routes
	router := routes.SetupRouter()
	routes.RegisterRoutes(router, customerController)

	gin.SetMode(config.Server.Mode)
	errRouter := router.Run(fmt.Sprintf(":%d", config.Server.Port))

	if errRouter != nil {
		panic(fmt.Sprintf("failed to start server: %v", err))
	}
}

func connectToPostgres(config *AppConfiguration) *sql.DB {
	// Setup database connection
	dbConfig := config.Database
	// Create connection string from config
	driverName := "postgres"
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName)

	// Open database connection
	db, err := sql.Open(driverName, connectionString)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	// Ping the database to verify connection
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("Connected to database successfully")
	return db
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
	viper.AddConfigPath("configs")                              // Direct subfolder
	viper.AddConfigPath("../../configs")                        // Two levels up

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config AppConfiguration
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}
