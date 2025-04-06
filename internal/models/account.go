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

type AccountResponse struct {
	AccountNumber      string     `json:"account_number"`
	Balance            float64    `json:"balance"`
	AccountDescription string     `json:"account_description"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
	DeletedAt          *time.Time `json:"deleted_at,omitempty"`
}

// ToResponse Convert from DB model to response model
func (a Account) ToResponse() AccountResponse {
	var deletedAt *time.Time
	if a.DeletedAt.Valid {
		deletedAt = &a.DeletedAt.Time
	}

	return AccountResponse{
		AccountNumber:      a.AccountNumber,
		Balance:            a.Balance,
		AccountDescription: a.AccountDescription,
		CreatedAt:          a.CreatedAt,
		UpdatedAt:          a.UpdatedAt,
		DeletedAt:          deletedAt,
	}
}
