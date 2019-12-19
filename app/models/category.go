package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"time"
)

type Category struct {
	ID        string    `gorm:"column:id;primary_key:true"`
	Code      string    `gorm:"type:varchar(20);column:code"`
	Name      string    `gorm:"type:varchar(50);column:name"`
	CreatedAt time.Time `gorm:"column:created_at;"`
	UpdatedAt time.Time `gorm:"column:updated_at;"`
}

func (c Category) TableName() string {
	return "categories"
}

// BeforeCreate - Lifecycle callback - Generate UUID before persisting
func (c *Category) BeforeCreate(scope *gorm.Scope) error {
	if err := scope.SetColumn("id", uuid.New().String()); err != nil {
		log.Fatal("Error UUID Generate")
	}
	return nil
}
