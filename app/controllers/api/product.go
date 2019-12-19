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

type ProductController struct {
	ProductRepository repository.ProductRepository
}

func (uc *ProductController) Store(c echo.Context) error {
	productRequest := request.ProductRequest{}
	if err := c.Bind(&productRequest); err != nil {
		return response.BadRequest(c, err.Error(), nil)
	}
	data, err := uc.ProductRepository.Create(productRequest)
	if err != nil {
		return response.BadRequest(c, err.Error(), nil)
	}
	serialized := mapper.NewProductMapper().Map(data)
	return response.SingleData(c, "Success Execute request", serialized)
}

func (uc *ProductController) GetById(c echo.Context) error {
	id := c.Param("id")
	data, err := uc.ProductRepository.GetByID(id)
	if err != nil {
		return response.BadRequest(c, err.Error(), nil)
	}
	serialized := mapper.NewProductMapper().Map(*data)
	return response.SingleData(c, "Success Execute request", serialized)
}

func (uc *ProductController) Update(c echo.Context) error {
	id := c.Param("id")
	product := request.ProductUpdateRequest{}
	if err := c.Bind(&product);
		err != nil {
		return response.BadRequest(c, err.Error(), nil)
		}
	data, err := uc.ProductRepository.Update(id, product)
	if err != nil{
		return response.BadRequest(c, err.Error(), nil)
	}
	return response.SingleData(c,"Update Success",data)
}

func (uc *ProductController) Destroy(c echo.Context) error {
	id := c.Param("id")
	err := uc.ProductRepository.Destroy(id)
	if err != nil{
		return response.BadRequest(c, err.Error(), nil)
	}
	return response.SingleData(c,"Success Delete Record", err)
}

func (uc *ProductController) FindAll(c echo.Context) error {
	data,err := uc.ProductRepository.FindAll()
	store, err := paging.NewGORMStore(config.GetLinkDb() ,&data)
	options := paging.NewOptions()
	r,_ := http.NewRequest("GET", c.Request().URL.String(), nil)

	paginator, _ := paging.NewOffsetPaginator(store, r, options)
	err = paginator.Page()
	if err != nil {
		return response.BadRequest(c, err.Error(), nil)
	}
	serialized := mapper.NewProductMapper().MapList(data)

	return response.Paginate(c,"Success Show Record", paginator,serialized)
}