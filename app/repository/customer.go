package repository

import (
	"echo-api-starter/app/models"
	"echo-api-starter/app/request"
)

type CustomerRepository interface {
	GetByID(id string) (*models.Customer, error)
	FindAll() ([]models.Customer, error)
	Update(string, request.CustomerUpdateRequest) (*models.Customer, error)
	Destroy(string) error
	Create(request request.CustomerRequest) (models.Customer, error)
}
