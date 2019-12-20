package service

import (
	"echo-api-starter/app/models"
	"echo-api-starter/app/repository"
	"echo-api-starter/app/request"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
)

// UserService :
type CustomerService struct {
	db *gorm.DB
}
// NewUserService :
func NewCustomerService(db *gorm.DB) repository.CustomerRepository {
	return &CustomerService{
		db,
	}
}

// GetByID :
func (us *CustomerService) GetByID(id string) (*models.Customer, error) {
	var m models.Customer
	if err := us.db.Where(&models.Customer{ID: id}).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, err
		}
	}
	return &m, nil
}

// Create :
func (us *CustomerService) Create(u request.CustomerRequest) (data models.Customer, err error) {
	data.Name = u.Name
	data.Email = u.Email
	data.Address = u.Address
	data.Phone = u.Phone
	err = us.db.Create(&data).Error
	return data, err
}

func (us *CustomerService) Update(id string, u request.CustomerUpdateRequest) (*models.Customer, error) {
	data, err := us.FindById(id)
	if err != nil {
		return &data, err
	}
	data.Name = u.Name
	data.Email = u.Email
	data.Address = u.Address
	data.Phone = u.Phone

	if err := us.db.Save(&data).Error; err != nil {
		log.Debug("ERROR", err)
		return &data, err
	}
	return &data, nil
}

func (us *CustomerService) FindById(id string) (c models.Customer, err error) {
	if err := us.db.Where("id = ?", id).First(&c).Error; err != nil {
		log.Debug("ERROR", err)
		return c, err
	}
	return c, err
}

func (us *CustomerService) FindAll() (data []models.Customer, err error) {
	err = us.db.Find(&data).Error
	return data, err
}

func (us *CustomerService) Destroy(id string) error {
	var m models.Customer
	if err := us.db.Where(&models.Customer{ID: id}).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return err
		}
		return err
	}
	if err := us.db.Delete(&models.Customer{ID: id}).Error; err != nil {
		return err
	}
	return nil
}
