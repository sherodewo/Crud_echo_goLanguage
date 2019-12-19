package repository

import (
	"echo-api-starter/app/models"
	"echo-api-starter/app/request"
)

type CategoryRepository interface {
	GetByID(id string) (*models.Category, error)
	FindAll() ([]models.Category, error)
	Update(string, request.CategoryUpdateRequest) (*models.Category, error)
	Destroy(string) error
	Create(request request.CategoryRequest) (models.Category, error)
}
