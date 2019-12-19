package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"time"
)

type Customer struct {
	ID        int       `gorm:"column:id;type:varchar(255);primary_key:true"`
	Name      string    `gorm:"column:name;type:varchar(100)"`
	Email     string    `gorm:"column:email;type:varchar(50);unique"`
	Address   string    `gorm:"column:address;type:varchar(255)"`
	Phone     string    `gorm:"column:phone;type:varchar(13);unique"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column;updated_at"`
}

func (c Customer) TableName() string {
	return "customers"
}

// BeforeCreate - Lifecycle callback - Generate UUID before persisting
func (c *Customer) BeforeCreate(scope *gorm.Scope) error {
	if err := scope.SetColumn("id", uuid.New().String()); err != nil {
		log.Fatal("Error UUID Generate")
	}
	return nil
}
