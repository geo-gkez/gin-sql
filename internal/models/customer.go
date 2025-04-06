package models

type Customer struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone"`
}

type CustomerResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

// CustomerWithAccounts represents a customer with their accounts
type CustomerWithAccounts struct {
	FirstName string            `json:"first_name"`
	LastName  string            `json:"last_name"`
	Email     string            `json:"email"`
	Phone     string            `json:"phone"`
	Accounts  []AccountResponse `json:"accounts"`
}

// ToResponse converts the Customer model to CustomerResponse DTO
func (c Customer) ToResponse() CustomerResponse {
	return CustomerResponse{
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Email:     c.Email,
		Phone:     c.Phone,
	}
}

// ToResponseWithAccounts converts the Customer model to CustomerWithAccounts DTO
func (c Customer) ToResponseWithAccounts(accounts []AccountResponse) CustomerWithAccounts {
	return CustomerWithAccounts{
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Email:     c.Email,
		Phone:     c.Phone,
		Accounts:  accounts,
	}
}
