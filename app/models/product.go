package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"time"
)

type Product struct {
	ID        string    `gorm:"column:id;primary_key:true"`
	Name        string    `gorm:"column:name;type:varchar(100)"`
	CategoryId        string    `gorm:"column:category_id;type:varchar(100)"`
	Description string    `gorm:"column:description;type:varchar(255);"`
	Price       string     `gorm:"column:price;"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (p Product) TableName() string {
	return "products"
}

// BeforeCreate - Lifecycle callback - Generate UUID before persisting
func (c *Product) BeforeCreate(scope *gorm.Scope) error {
	if err := scope.SetColumn("id", uuid.New().String()); err != nil {
		log.Fatal("Error UUID Generate")
	}
	return nil
}
