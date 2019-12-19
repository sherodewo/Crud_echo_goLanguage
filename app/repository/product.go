package repository

import (
	"echo-api-starter/app/models"
	"echo-api-starter/app/request"
)

type ProductRepository interface {
	GetByID(id string) (*models.Product, error)
	FindAll() ([]models.Product, error)
	Update(string, request.ProductUpdateRequest) (*models.Product, error)
	Destroy(string) error
	Create(request request.ProductRequest) (models.Product, error)
}
