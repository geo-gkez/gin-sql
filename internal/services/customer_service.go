package services

import (
	"fmt"
	"org/gg/banking/internal/middleware/errors"
	"org/gg/banking/internal/models"
	"org/gg/banking/internal/repository"
)

type ICustomerService interface {
	FindAll() ([]models.CustomerDTO, error)
	FindCustomerWithAccounts(email string) (models.CustomerDTO, error)
	CreateCustomer(customer models.CustomerDTO) (models.CustomerDTO, error)
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
func (s *customerService) FindAll() ([]models.CustomerDTO, error) {
	customers, err := s.customerRepository.FindAll()
	if err != nil {
		// Transform technical errors to domain errors
		return nil, errors.InternalServerError(fmt.Sprintf("Failed to retrieve customers: %v", err))
	}

	if len(customers) == 0 {
		// Business logic decision - empty result might be an error
		return nil, errors.NotFoundError("No customers found")
	}

	// Convert Customer models to CustomerDTO models
	var customerResponses []models.CustomerDTO
	for _, customer := range customers {
		customerResponses = append(customerResponses, customer.ToCustomerDTO())
	}

	return customerResponses, nil
}

// FindCustomerWithAccounts retrieves a CustomerDTO containing customer details and associated account information
func (s *customerService) FindCustomerWithAccounts(email string) (models.CustomerDTO, error) {
	customer, err := s.customerRepository.FindByEmail(email)
	if err != nil {
		return models.CustomerDTO{}, errors.NotFoundError(fmt.Sprintf("Customer with email %s not found: %v", email, err))
	}

	accounts, err := s.accountRepository.FindByCustomerID(customer.ID)
	if err != nil {
		return models.CustomerDTO{}, errors.InternalServerError(fmt.Sprintf("Failed to retrieve accounts for customer with email %s: %v", email, err))
	}

	var accountResponses []models.AccountDTO
	for _, account := range accounts {
		accountResponses = append(accountResponses, account.ToAccountDTO())
	}

	return customer.ToCustomerDTO(accountResponses), nil
}

// CreateCustomer creates a new customer and their accounts inline within this method
func (s *customerService) CreateCustomer(customerDto models.CustomerDTO) (models.CustomerDTO, error) {
	createdCustomer, err := s.customerRepository.Create(customerDto.ToCustomer())
	if err != nil {
		return models.CustomerDTO{}, errors.InternalServerError(fmt.Sprintf("Failed to create customer: %v", err))
	}

	var createdAccountDtos []models.AccountDTO
	if len(customerDto.Accounts) > 0 {
		for _, accountDto := range customerDto.Accounts {
			account := accountDto.ToAccount()
			account.CustomerID = createdCustomer.ID
			createdAccount, err := s.accountRepository.CreateAccount(createdCustomer.ID, account)
			if err != nil {
				return models.CustomerDTO{}, errors.InternalServerError(fmt.Sprintf("Failed to create account for customer: %v", err))
			}
			createdAccountDtos = append(createdAccountDtos, createdAccount.ToAccountDTO())
		}
	}

	return createdCustomer.ToCustomerDTO(createdAccountDtos), nil
}
