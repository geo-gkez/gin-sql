package repository

import (
	"database/sql"
	"org/gg/banking/internal/models"
)

type ICustomerRepository interface {
	GetCustomers() ([]models.Customer, error)
}

type CustomerRepository struct {
	db         *sql.DB
	collection string
}

func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{
		db:         db,
		collection: "customer_collection",
	}
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
