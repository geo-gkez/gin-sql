package controllers

import (
	"fmt"
	"net/http"
	"org/gg/banking/services"

	"github.com/gin-gonic/gin"
)

// ICustomerController defines the interface for customer-related HTTP handlers
type ICustomerController interface {
	GetCustomers(ctx *gin.Context)
}

type CustomerController struct {
	service services.ICustomerService
}

func NewCustomerController(service services.ICustomerService) *CustomerController {
	return &CustomerController{
		service: service,
	}
}

// GetCustomers handles the HTTP request to fetch all customers
func (c *CustomerController) GetCustomers(ctx *gin.Context) {
	customers, err := c.service.GetCustomers()
	if err != nil {
		ctx.Error(fmt.Errorf("getting customer list: %w", err))
		return
	}

	ctx.JSON(http.StatusOK, customers)
}
