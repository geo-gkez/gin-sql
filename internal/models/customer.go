package models

// Customer represents a customer in the banking system
type Customer struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone"`
}

// CustomerDTO represents a customer with their accounts
type CustomerDTO struct {
	FirstName string       `json:"first_name"`
	LastName  string       `json:"last_name"`
	Email     string       `json:"email"`
	Phone     string       `json:"phone"`
	Accounts  []AccountDTO `json:"accounts,omitempty"`
}

// ToCustomerDTO converts the Customer model to CustomerDTO.
// The variadic 'accounts' parameter allows passing an optional slice of AccountDTO.
// If one or more slices are provided, the first slice is used to populate the Accounts field in the DTO.
// This is useful for including related account information when converting a Customer to CustomerDTO.
func (c Customer) ToCustomerDTO(accounts ...[]AccountDTO) CustomerDTO {
	dto := CustomerDTO{
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Email:     c.Email,
		Phone:     c.Phone,
	}

	// Add accounts if provided
	if len(accounts) > 0 {
		dto.Accounts = accounts[0]
	}

	return dto
}

// ToCustomer converts the CustomerDTO DTO to Customer model
func (c CustomerDTO) ToCustomer() Customer {
	return Customer{
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Email:     c.Email,
		Phone:     c.Phone,
	}
}
