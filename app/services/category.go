package service

import (
	"echo-api-starter/app/models"
	"echo-api-starter/app/repository"
	"echo-api-starter/app/request"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
)

// UserService :
type CategoryService struct {
	db *gorm.DB
}
// NewUserService :
func NewCategoryService(db *gorm.DB) repository.ProductRepository {
	return &ProductService{
		db,
	}
}

// GetByID :
func (us *CategoryService) GetByID(id string) (*models.Category, error) {
	var m models.Category
	if err := us.db.Where(&models.Category{ID: id}).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, err
		}
	}
	return &m, nil
}

// Create :
func (us *CategoryService) Create(u request.CategoryRequest) (data models.Category, err error) {
	data.Name = u.Name
	data.Code = u.Code
	err = us.db.Create(&data).Error
	return data, err
}

func (us *CategoryService) Update(id string, um request.CategoryUpdateRequest) (*models.Category, error) {
	data, err := us.FindById(id)
	if err != nil {
		return &data, err
	}
	data.Name = um.Name
	data.Code = um.Code

	if err := us.db.Save(&data).Error; err != nil {
		log.Debug("ERROR", err)
		return &data, err
	}
	return &data, nil
}

func (us *CategoryService) FindById(id string) (c models.Category, err error) {
	if err := us.db.Where("id = ?", id).First(&c).Error; err != nil {
		log.Debug("ERROR", err)
		return c, err
	}
	return c, err
}

func (us *CategoryService) FindAll() (data []models.Product, err error) {
	err = us.db.Find(&data).Error
	return data, err
}

func (us *CategoryService) Destroy(id string) error {
	var m models.Product
	if err := us.db.Where(&models.Product{ID: id}).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return err
		}
		return err
	}
	if err := us.db.Delete(&models.Product{ID: id}).Error; err != nil {
		return err
	}
	return nil
}
