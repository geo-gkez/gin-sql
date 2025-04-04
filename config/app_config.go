package config

import (
	"github.com/gin-gonic/gin"
	"org/gg/banking/controllers"
	"org/gg/banking/repository"
	"org/gg/banking/routes"
	"org/gg/banking/services"
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
