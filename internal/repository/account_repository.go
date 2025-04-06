package repository

import (
	"database/sql"
	"fmt"
	"log"
	"org/gg/banking/internal/models"
)

type IAccountRepository interface {
	FindByCustomerID(customerID int64) ([]models.Account, error)
}

type accountRepository struct {
	db         *sql.DB
	collection string
}

func NewAccountRepository(db *sql.DB) IAccountRepository {
	return &accountRepository{
		db:         db,
		collection: "account_collection",
	}
}

// FindByCustomerID retrieves all non-deleted accounts for a customer
func (repository *accountRepository) FindByCustomerID(customerID int64) ([]models.Account, error) {
	query := `
		SELECT id, customer_id, account_number, balance, account_description, created_at, updated_at, deleted_at
		FROM accounts
		WHERE customer_id = $1 AND deleted_at IS NULL
	`

	rows, err := repository.db.Query(query, customerID)
	if err != nil {
		return nil, fmt.Errorf("error querying customer accounts: %v", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	var accounts []models.Account
	for rows.Next() {
		var account models.Account
		if err := rows.Scan(
			&account.ID,
			&account.CustomerID,
			&account.AccountNumber,
			&account.Balance,
			&account.AccountDescription,
			&account.CreatedAt,
			&account.UpdatedAt,
			&account.DeletedAt,
		); err != nil {
			return nil, fmt.Errorf("error scanning account row: %v", err)
		}
		accounts = append(accounts, account)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating accounts: %v", err)
	}

	return accounts, nil
}
