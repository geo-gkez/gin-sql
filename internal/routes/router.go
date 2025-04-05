package routes

import (
	"github.com/gin-gonic/gin"
	"org/gg/banking/internal/controllers"
	"org/gg/banking/internal/middleware/errors"
)

// SetupRouter initializes the Gin router and applies middleware
func SetupRouter() *gin.Engine {
	router := gin.Default()
	// Register error middleware
	router.Use(errors.ErrorHandlerMiddleware())

	return router
}

// RegisterRoutes adds all application routes to the router
func RegisterRoutes(router *gin.Engine, customerController *controllers.CustomerController) {
	// Customer routes
	router.GET("/customers", customerController.GetCustomers)
	//router.POST("/customers", customerController.CreateCustomer)
	//router.GET("/customers/:id", customerController.GetCustomerByID)
}
