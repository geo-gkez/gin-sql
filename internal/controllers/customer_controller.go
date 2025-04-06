package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"org/gg/banking/internal/services"
)

// ICustomerController defines the interface for customer-related HTTP handlers
type ICustomerController interface {
	GetCustomers(ctx *gin.Context)
}

type customerController struct {
	service services.ICustomerService
}

func NewCustomerController(service services.ICustomerService) ICustomerController {
	return &customerController{
		service: service,
	}
}

// GetCustomers handles the HTTP request to fetch all customers
func (c *customerController) GetCustomers(ctx *gin.Context) {
	customers, err := c.service.FindAll()
	if err != nil {
		ctx.Error(fmt.Errorf("getting customer list: %w", err))
		return
	}

	ctx.JSON(http.StatusOK, customers)
}
