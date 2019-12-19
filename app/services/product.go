package service

import (
	"echo-api-starter/app/models"
	"echo-api-starter/app/repository"
	"echo-api-starter/app/request"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
)

// UserService :
type ProductService struct {
	db *gorm.DB
}
// NewUserService :
func NewProductService(db *gorm.DB) repository.ProductRepository {
	return &ProductService{
		db,
	}
}

// GetByID :
func (us *ProductService) GetByID(id string) (*models.Product, error) {
	var m models.Product
	if err := us.db.Where(&models.Product{ID: id}).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, err
		}
	}
	return &m, nil
}

// Create :
func (us *ProductService) Create(u request.ProductRequest) (data models.Product, err error) {
	data.Name = u.Name
	data.Description = u.Description
	data.Price = u.Price
	data.CategoryId = u.CategoryId
	err = us.db.Create(&data).Error
	return data, err
}

func (us *ProductService) Update(id string, um request.ProductUpdateRequest) (*models.Product, error) {
	data, err := us.FindById(id)
	if err != nil {
		return &data, err
	}
	data.Name = um.Name
	data.Description = um.Description
	data.Price = um.Price
	data.CategoryId = um.CategoryId

	if err := us.db.Save(&data).Error; err != nil {
		log.Debug("ERROR", err)
		return &data, err
	}
	return &data, nil
}

func (us *ProductService) FindById(id string) (c models.Product, err error) {
	if err := us.db.Where("id = ?", id).First(&c).Error; err != nil {
		log.Debug("ERROR", err)
		return c, err
	}
	return c, err
}

func (us *ProductService) FindAll() (data []models.Product, err error) {
	err = us.db.Find(&data).Error
	return data, err
}

func (us *ProductService) Destroy(id string) error {
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
