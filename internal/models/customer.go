package models

type Customer struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone"`
}

type CustomerDTO struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

// CustomerAggregateDTO represents a customer with their accounts
type CustomerAggregateDTO struct {
	FirstName string       `json:"first_name"`
	LastName  string       `json:"last_name"`
	Email     string       `json:"email"`
	Phone     string       `json:"phone"`
	Accounts  []AccountDTO `json:"accounts,omitempty"`
}

// ToCustomerDTO converts the Customer model to CustomerDTO DTO
func (c Customer) ToCustomerDTO() CustomerDTO {
	return CustomerDTO{
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Email:     c.Email,
		Phone:     c.Phone,
	}
}

// ToCustomerAggregateDTO converts the Customer model to CustomerAggregateDTO DTO
func (c Customer) ToCustomerAggregateDTO(accounts []AccountDTO) CustomerAggregateDTO {
	return CustomerAggregateDTO{
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Email:     c.Email,
		Phone:     c.Phone,
		Accounts:  accounts,
	}
}
