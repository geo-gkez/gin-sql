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
func (c Customer) ToCustomerDTO(accounts ...AccountDTO) CustomerDTO {
	customerDto := CustomerDTO{
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Email:     c.Email,
		Phone:     c.Phone,
	}

	// Add accounts if provided
	if len(accounts) > 0 {
		customerDto.Accounts = accounts
	}

	return customerDto
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
