package api

import (
	"echo-api-starter/app/repository"
	"echo-api-starter/app/request"
	"echo-api-starter/app/response"
	"echo-api-starter/app/response/mapper"
	"github.com/labstack/echo"
	_ "github.com/labstack/echo"
)

type UserController struct {
	UserRepository repository.UserRepository
}

func (uc *UserController) Store(c echo.Context) error {
	userRequest := request.UserRequest{}
	if err := c.Bind(&userRequest); err != nil {
		return response.BadRequest(c, err.Error(), nil)
	}
	data, err := uc.UserRepository.Create(userRequest)
	if err != nil {
		return response.BadRequest(c, err.Error(), nil)
	}
	serialized := mapper.NewUserMapper().Map(data)
	return response.SingleData(c, "Success Execute request", serialized)
}

func (uc *UserController) GetByID(c echo.Context) error {
	id := c.Param("id")
	data, err := uc.UserRepository.GetByID(id)
	if err != nil {
		return response.BadRequest(c, err.Error(), nil)
	}
	serialized := mapper.NewUserMapper().Map(*data)
	return response.SingleData(c, "Success Execute request", serialized)
}
