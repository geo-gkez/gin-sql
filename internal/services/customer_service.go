package services

import (
	"fmt"
	"org/gg/banking/internal/middleware/errors"
	"org/gg/banking/internal/models"
	"org/gg/banking/internal/repository"
)

type ICustomerService interface {
	FindAll() ([]models.Customer, error)
}

type customerService struct {
	repo repository.ICustomerRepository
}

// NewCustomerService creates a new service with the provided repository
func NewCustomerService(repo repository.ICustomerRepository) ICustomerService {
	return &customerService{
		repo: repo,
	}
}

// FindAll GetCustomers delegates to the repository layer
func (s *customerService) FindAll() ([]models.Customer, error) {
	customers, err := s.repo.FindAll()
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
