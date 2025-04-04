package models

type Customer struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Accounts  []Account `json:"accounts,omitempty" gorm:"foreignKey:CustomerID"`
}
