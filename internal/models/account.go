package models

import "time"

type Account struct {
	ID                 int        `json:"id"`
	CustomerID         int        `json:"customer_id"`
	AccountNumber      string     `json:"account_number"`
	Balance            float64    `json:"balance"`
	AccountDescription string     `json:"account_description"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
	DeletedAt          *time.Time `json:"deleted_at,omitempty"`
}
