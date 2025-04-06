package services

import (
	"fmt"
	"org/gg/banking/internal/middleware/errors"
	"org/gg/banking/internal/models"
	"org/gg/banking/internal/repository"
)

type ICustomerService interface {
	FindAll() ([]models.CustomerResponse, error)
	FindCustomerWithAccounts(email string) (models.CustomerWithAccounts, error)
	CreateCustomer(customer models.Customer) (models.CustomerResponse, error)
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
func (s *customerService) FindAll() ([]models.CustomerResponse, error) {
	customers, err := s.customerRepository.FindAll()
	if err != nil {
		// Transform technical errors to domain errors
		return nil, errors.InternalServerError(fmt.Sprintf("Failed to retrieve customers: %v", err))
	}

	if len(customers) == 0 {
		// Business logic decision - empty result might be an error
		return nil, errors.NotFoundError("No customers found")
	}

	// Convert Customer models to CustomerResponse models
	var customerResponses []models.CustomerResponse
	for _, customer := range customers {
		customerResponses = append(customerResponses, customer.ToResponse())
	}

	return customerResponses, nil
}

// FindCustomerWithAccounts GetCustomerWithAccounts retrieves a customer with all their accounts
func (s *customerService) FindCustomerWithAccounts(email string) (models.CustomerWithAccounts, error) {
	customer, err := s.customerRepository.FindByEmail(email)
	if err != nil {
		return models.CustomerWithAccounts{}, errors.NotFoundError(fmt.Sprintf("Customer with email %s not found: %v", email, err))
	}

	accounts, err := s.accountRepository.FindByCustomerID(customer.ID)
	if err != nil {
		return models.CustomerWithAccounts{}, errors.InternalServerError(fmt.Sprintf("Failed to retrieve accounts for customer with email %s: %v", email, err))
	}

	var accountResponses []models.AccountResponse
	for _, account := range accounts {
		accountResponses = append(accountResponses, account.ToResponse())
	}

	return customer.ToResponseWithAccounts(accountResponses), nil
}

// CreateCustomer creates a new customer and their accounts
func (s *customerService) CreateCustomer(customer models.Customer) (models.CustomerResponse, error) {
	createdCustomer, err := s.customerRepository.Create(customer)
	if err != nil {
		return models.CustomerResponse{}, errors.InternalServerError(fmt.Sprintf("Failed to create customer: %v", err))
	}

	// Assuming accounts are created in a separate process
	// You can add logic here to create accounts if needed

	return createdCustomer.ToResponse(), nil
}
