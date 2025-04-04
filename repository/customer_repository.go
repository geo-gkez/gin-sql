package repository

import (
	"org/gg/banking/models"
)

type ICustomerRepository interface {
	GetCustomers() ([]models.Customer, error)
}

type CustomerRepository struct {
	collection string
}

func NewCustomerRepository() *CustomerRepository {
	return &CustomerRepository{"customer_collection"}
}

func (c *CustomerRepository) GetCustomers() ([]models.Customer, error) {
	customers := []models.Customer{
		{
			ID:        1,
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe@example.com",
			Phone:     "555-1234",
		},
		{
			ID:        2,
			FirstName: "Jane",
			LastName:  "Smith",
			Email:     "jane.smith@example.com",
			Phone:     "555-5678",
		},
		{
			ID:        3,
			FirstName: "George",
			LastName:  "Gkezeris",
			Email:     "gg@gmail.com",
			Phone:     "699999",
		},
	}
	return customers, nil
}
