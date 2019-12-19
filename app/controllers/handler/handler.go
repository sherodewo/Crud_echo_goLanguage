package handler

import (
	"echo-api-starter/app/controllers/api"
	service "echo-api-starter/app/services"
	"github.com/jinzhu/gorm"
)

func NewUserController(db *gorm.DB) *api.UserController {
	userService:=service.NewUserService(db)
	return &api.UserController{UserRepository: userService}
}
func NewProductController(db *gorm.DB) *api.ProductController {
	productService:=service.NewProductService(db)
	return &api.ProductController{ProductRepository: productService}
}