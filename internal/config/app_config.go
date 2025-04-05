package config

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
