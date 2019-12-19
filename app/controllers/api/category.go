package api

import (
	"echo-api-starter/app/repository"
	"echo-api-starter/app/request"
	"echo-api-starter/app/response"
	"echo-api-starter/app/response/mapper"
	"echo-api-starter/config"
	"github.com/labstack/echo"
	"github.com/ulule/paging"
	"net/http"
)

type CategoryController struct {
	CategoryRepository repository.CategoryRepository

}

func (uc *CategoryController) Store(c echo.Context) error {
	categoryRequest := request.CategoryRequest{}
	if err := c.Bind(&categoryRequest); err != nil {
		return response.BadRequest(c, err.Error(), nil)
	}
	data, err := uc.CategoryRepository.Create(categoryRequest)
	if err != nil {
		return response.BadRequest(c, err.Error(), nil)
	}
	serialized := mapper.NewCategoryMapper().Map(data)
	return response.SingleData(c, "Success Execute request", serialized)
}

func (uc *CategoryController) GetById(c echo.Context) error {
	id := c.Param("id")
	data, err := uc.CategoryRepository.GetByID(id)
	if err != nil {
		return response.BadRequest(c, err.Error(), nil)
	}
	serialized := mapper.NewCategoryMapper().Map(*data)
	return response.SingleData(c, "Success Execute request", serialized)
}

func (uc *CategoryController) Update(c echo.Context) error {
	id := c.Param("id")
	category := request.CategoryUpdateRequest{}
	if err := c.Bind(&category);
		err != nil {
		return response.BadRequest(c, err.Error(), nil)
	}
	data, err := uc.CategoryRepository.Update(id, category)
	if err != nil{
		return response.BadRequest(c, err.Error(), nil)
	}
	return response.SingleData(c,"Update Success",data)
}

func (uc *CategoryController) Destroy(c echo.Context) error {
	id := c.Param("id")
	err := uc.CategoryRepository.Destroy(id)
	if err != nil{
		return response.BadRequest(c, err.Error(), nil)
	}
	return response.SingleData(c,"Success Delete Record", err)
}

func (uc *CategoryController) FindAll(c echo.Context) error {
	data,err := uc.CategoryRepository.FindAll()
	store, err := paging.NewGORMStore(config.GetLinkDb() ,&data)
	options := paging.NewOptions()
	r,_ := http.NewRequest("GET", c.Request().URL.String(), nil)

	paginator, _ := paging.NewOffsetPaginator(store, r, options)
	err = paginator.Page()
	if err != nil {
		return response.BadRequest(c, err.Error(), nil)
	}
	serialized := mapper.NewCategoryMapper().MapList(data)

	return response.Paginate(c,"Success Show Record", paginator,serialized)
}