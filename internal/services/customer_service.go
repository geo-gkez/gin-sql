package services

import (
	"fmt"
	"org/gg/banking/internal/middleware/errors"
	"org/gg/banking/internal/models"
	"org/gg/banking/internal/repository"
)

type ICustomerService interface {
	FindAll() ([]models.Customer, error)
	FindCustomerWithAccounts(email string) (models.CustomerWithAccounts, error)
}

type customerService struct {
	customerRepository repository.ICustomerRepository
	accountRepository  repository.IAccountRepository
}

// NewCustomerService creates a new service with the provided repository
func NewCustomerService(customerRepository repository.ICustomerRepository, accountRepository repository.IAccountRepository) ICustomerService {
	return &customerService{
		customerRepository: customerRepository,
		accountRepository:  accountRepository,
	}
}

// FindAll GetCustomers delegates to the repository layer
func (s *customerService) FindAll() ([]models.Customer, error) {
	customers, err := s.customerRepository.FindAll()
	if err != nil {
		// Transform technical errors to domain errors
		return nil, errors.InternalServerError(fmt.Sprintf("Failed to retrieve customers: %v", err))
	}

	if len(customers) == 0 {
		// Business logic decision - empty result might be an error
		return nil, errors.NotFoundError("No customers found")
	}

	return customers, nil
}

// FindCustomerWithAccounts GetCustomerWithAccounts retrieves a customer with all their accounts
func (s *customerService) FindCustomerWithAccounts(email string) (models.CustomerWithAccounts, error) {
	customer, err := s.customerRepository.FindByEmail(email)
	if err != nil {
		return models.CustomerWithAccounts{}, err
	}

	accounts, err := s.accountRepository.FindByCustomerID(customer.ID)
	if err != nil {
		return models.CustomerWithAccounts{}, err
	}

	var accountResponses []models.AccountResponse
	for _, account := range accounts {
		accountResponses = append(accountResponses, account.ToResponse())
	}

	return customer.ToResponseWithAccounts(accountResponses), nil
}
