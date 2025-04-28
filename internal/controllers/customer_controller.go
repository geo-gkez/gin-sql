package controllers

import (
	"fmt"
	"net/http"
	"org/gg/banking/internal/models"
	"org/gg/banking/internal/services"

	"github.com/gin-gonic/gin"
)

// ICustomerController defines the interface for customer-related HTTP handlers
type ICustomerController interface {
	GetCustomers(ctx *gin.Context)
	GetCustomerByEmail(ctx *gin.Context)
	CreateCustomer(ctx *gin.Context)
	DeleteCustomerByEmail(ctx *gin.Context) // New method
}

type customerController struct {
	customerService services.ICustomerService
}

func NewCustomerController(service services.ICustomerService) ICustomerController {
	return &customerController{
		customerService: service,
	}
}

// GetCustomers handles the HTTP request to fetch all customers
func (c *customerController) GetCustomers(ctx *gin.Context) {
	customers, err := c.customerService.FindAll()
	if err != nil {
		ctx.Error(fmt.Errorf("getting customer list: %w", err))
		return
	}

	ctx.JSON(http.StatusOK, customers)
}

func (c *customerController) GetCustomerByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	if email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email parameter is required"})
		return
	}

	customer, err := c.customerService.FindCustomerWithAccounts(email)
	if err != nil {
		ctx.Error(fmt.Errorf("getting customer by email: %w", err))
		return
	}

	ctx.JSON(http.StatusOK, customer)
}

// CreateCustomer handles the HTTP request to create a new customer
func (c *customerController) CreateCustomer(ctx *gin.Context) {
	var customer models.CustomerDTO
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	createdCustomer, err := c.customerService.CreateCustomer(customer)
	if err != nil {
		ctx.Error(fmt.Errorf("creating customer: %w", err))
		return
	}

	ctx.JSON(http.StatusCreated, createdCustomer)
}

// Implement the delete method
func (c *customerController) DeleteCustomerByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	if email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email parameter is required"})
		return
	}

	err := c.customerService.DeleteCustomerByEmail(email)
	if err != nil {
		ctx.Error(fmt.Errorf("deleting customer by email: %w", err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Customer successfully deleted"})
}
