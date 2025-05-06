package routes

import (
	"org/gg/banking/internal/controllers"
	"org/gg/banking/internal/middleware/errors"
	"org/gg/banking/internal/middleware/logger"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes the Gin router and applies middleware
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(logger.HTTPLoggerMiddleware())
	// Register error middleware
	router.Use(errors.ErrorHandlerMiddleware())

	return router
}

// RegisterRoutes adds all application routes to the router
func RegisterRoutes(router *gin.Engine, customerController controllers.ICustomerController) {
	// API v1 group
	v1 := router.Group("/api/v1")

	// Customer routes
	customerGroup := v1.Group("/customers")
	{
		customerGroup.GET("/", customerController.GetCustomers)
		customerGroup.GET("/:email", customerController.GetCustomerByEmail)
		customerGroup.POST("/", customerController.CreateCustomer)
		customerGroup.DELETE("/:email", customerController.DeleteCustomerByEmail)
	}

}
