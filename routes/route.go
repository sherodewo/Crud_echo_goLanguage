package routes

import (
	"echo-api-starter/app/controllers/handler"
	"echo-api-starter/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func New() (e *echo.Echo) {
	e = echo.New()
	//Route group
	api:=e.Group("/api")
	v1:=api.Group("/v1")
	//DB
	db := config.New()
	//Validation
	e.Validator = NewValidator()

	e.Logger.SetLevel(log.DEBUG)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	//USER
	userController:=handler.NewUserController(db)
	v1.POST("/user", userController.Store)
	v1.GET("/user/:id", userController.GetByID)

	//PRODUCT
	productController:=handler.NewProductController(db)
	v1.POST("/product", productController.Store)
	v1.GET("/product/:id", productController.GetById)
	v1.PUT("/product/:id", productController.Update)
	v1.DELETE("/product/:id", productController.Destroy)
	v1.GET("/product/show", productController.FindAll)

	//CATEGORY
	categoryController:=handler.NewProductController(db)
	v1.POST("/category", categoryController.Store)
	v1.GET("/category/:id", categoryController.GetById)
	v1.PUT("/category/:id", categoryController.Update)
	v1.DELETE("/category/:id", categoryController.Destroy)
	v1.GET("/category/show", categoryController.FindAll)

	//CATEGORY
	customerController:=handler.NewCustomerController(db)
	v1.POST("/category", customerController.Store)
	v1.GET("/category/:id", customerController.GetById)
	v1.PUT("/category/:id", customerController.Update)
	v1.DELETE("/category/:id", customerController.Destroy)
	v1.GET("/category/show", customerController.FindAll)
	return e
}
