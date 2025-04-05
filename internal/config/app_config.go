package config

import (
	"github.com/gin-gonic/gin"
	"org/gg/banking/internal/controllers"
	"org/gg/banking/internal/repository"
	"org/gg/banking/internal/routes"
	"org/gg/banking/internal/services"
)

func SetupApp() *gin.Engine {
	// Create components
	repo := repository.NewCustomerRepository()
	service := services.NewCustomerService(repo)
	controller := controllers.NewCustomerController(service)

	// Setup router and routes
	router := routes.SetupRouter()
	routes.RegisterRoutes(router, controller)

	return router
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

//func LoadConfig(configPath string) (*AppConfiguration, error) {
//	viper.SetConfigName("config")
//	viper.SetConfigType("yml")
//
//	if configPath != "" {
//		viper.AddConfigPath(configPath)
//	}
//	viper.AddConfigPath(".")
//
//	if err := viper.ReadInConfig(); err != nil {
//		return nil, fmt.Errorf("failed to read config file: %w", err)
//	}
//
//	var config AppConfiguration
//	if err := viper.Unmarshal(&config); err != nil {
//		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
//	}
//
//	return &config, nil
//}
