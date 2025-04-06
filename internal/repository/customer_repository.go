package repository

import (
	"database/sql"
	"log"
	"org/gg/banking/internal/models"
)

type ICustomerRepository interface {
	FindAll() ([]models.Customer, error)
}

type customerRepository struct {
	db         *sql.DB
	collection string
}

func NewCustomerRepository(db *sql.DB) ICustomerRepository {
	return &customerRepository{
		db:         db,
		collection: "customer_collection",
	}
}
func (c *customerRepository) FindAll() ([]models.Customer, error) {
	rows, err := c.db.Query("SELECT id, first_name, last_name, email, phone FROM customers")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	var customers []models.Customer
	for rows.Next() {
		var customer models.Customer
		if err := rows.Scan(&customer.ID, &customer.FirstName, &customer.LastName, &customer.Email, &customer.Phone); err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}
