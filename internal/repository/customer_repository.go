package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"org/gg/banking/internal/models"
)

type ICustomerRepository interface {
	FindAll() ([]models.Customer, error)
	FindByEmail(email string) (models.Customer, error)
	Create(customer models.Customer) (models.Customer, error)
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

func (repo *customerRepository) FindAll() ([]models.Customer, error) {
	rows, err := repo.db.Query("SELECT id, first_name, last_name, email, phone FROM customers")
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

// FindByEmail retrieves a customer by email
func (repo *customerRepository) FindByEmail(email string) (models.Customer, error) {
	var customer models.Customer

	query := `
		SELECT id, first_name, last_name, email, phone
		FROM customers
		WHERE email = $1
	`

	err := repo.db.QueryRow(query, email).Scan(
		&customer.ID,
		&customer.FirstName,
		&customer.LastName,
		&customer.Email,
		&customer.Phone,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Customer{}, fmt.Errorf("customer with email %s not found", email)
		}
		return models.Customer{}, fmt.Errorf("error querying customer by email: %v", err)
	}

	return customer, nil
}

// Create inserts a new customer into the database
// Create inserts a new customer into the database and returns the created customer
func (repo *customerRepository) Create(customer models.Customer) (models.Customer, error) {
	query := `
		INSERT INTO customers (first_name, last_name, email, phone)
		VALUES ($1, $2, $3, $4)
		RETURNING id, first_name, last_name, email, phone
	`

	var createdCustomer models.Customer
	err := repo.db.QueryRow(query, customer.FirstName, customer.LastName, customer.Email, customer.Phone).Scan(
		&createdCustomer.ID,
		&createdCustomer.FirstName,
		&createdCustomer.LastName,
		&createdCustomer.Email,
		&createdCustomer.Phone,
	)

	if err != nil {
		return models.Customer{}, fmt.Errorf("error inserting new customer: %v", err)
	}

	return createdCustomer, nil
}
