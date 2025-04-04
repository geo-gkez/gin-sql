package repository

import "org/gg/banking/models"

type CustomerRepository interface {
	GetCUstomers() ([]models.Customer, error)
}
