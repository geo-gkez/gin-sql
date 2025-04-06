package models

import (
	"database/sql"
	"time"
)

type Account struct {
	ID                 int64        `json:"id"`
	CustomerID         int64        `json:"customer_id"`
	AccountNumber      string       `json:"account_number"`
	Balance            float64      `json:"balance"`
	AccountDescription string       `json:"account_description"`
	CreatedAt          time.Time    `json:"created_at"`
	UpdatedAt          time.Time    `json:"updated_at"`
	DeletedAt          sql.NullTime `json:"deleted_at"`
}

type AccountDTO struct {
	AccountNumber      string     `json:"account_number"`
	Balance            float64    `json:"balance"`
	AccountDescription string     `json:"account_description"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
	DeletedAt          *time.Time `json:"deleted_at,omitempty"`
}

// ToAccountDTO Convert from DB model to response model
func (a Account) ToAccountDTO() AccountDTO {
	var deletedAt *time.Time
	if a.DeletedAt.Valid {
		deletedAt = &a.DeletedAt.Time
	}

	return AccountDTO{
		AccountNumber:      a.AccountNumber,
		Balance:            a.Balance,
		AccountDescription: a.AccountDescription,
		CreatedAt:          a.CreatedAt,
		UpdatedAt:          a.UpdatedAt,
		DeletedAt:          deletedAt,
	}
}

// ToAccount Convert from response model to DB model
func (a AccountDTO) ToAccount() Account {
	return Account{
		AccountNumber:      a.AccountNumber,
		Balance:            a.Balance,
		AccountDescription: a.AccountDescription,
		CreatedAt:          a.CreatedAt,
		UpdatedAt:          a.UpdatedAt,
	}
}
