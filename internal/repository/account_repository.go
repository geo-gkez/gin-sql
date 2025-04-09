package repository

import (
	"database/sql"
	"fmt"
	"log"
	"org/gg/banking/internal/models"
	"time"
)

type IAccountRepository interface {
	FindByCustomerID(customerID int64) ([]models.Account, error)
	CreateAccount(customerID int64, account models.Account) (models.Account, error)
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
		if err := rows.Close(); err != nil {
			log.Printf("error closing rows: %v", err)
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

func (repository *accountRepository) CreateAccount(customerID int64, account models.Account) (models.Account, error) {
	query := `
		INSERT INTO accounts (customer_id, account_number, balance, account_description, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id,account_number, balance, account_description, created_at, updated_at
	`

	stmt, err := repository.db.Prepare(query)
	if err != nil {
		return models.Account{}, fmt.Errorf("error preparing statement: %v", err)
	}
	defer func(stmt *sql.Stmt) {
		if err := stmt.Close(); err != nil {
			log.Printf("error closing statement: %v", err)
		}
	}(stmt)

	var createdAccount models.Account
	err2 := stmt.QueryRow(customerID, account.AccountNumber, account.Balance, account.AccountDescription, time.Now(), time.Now()).Scan(
		&createdAccount.ID,
		&createdAccount.AccountNumber,
		&createdAccount.Balance,
		&createdAccount.AccountDescription,
		&createdAccount.CreatedAt,
		&createdAccount.UpdatedAt)

	if err2 != nil {
		return models.Account{}, fmt.Errorf("error executing statement: %v", err2)
	}

	return createdAccount, nil
}
