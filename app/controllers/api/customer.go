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

type CustomerController struct {
	CustomerRepository repository.CustomerRepository
}

func (uc *CustomerController) Store(c echo.Context) error {
	customerRequest := request.CustomerRequest{}
	if err := c.Bind(&customerRequest); err != nil {
		return response.BadRequest(c, err.Error(), nil)
	}
	data, err := uc.CustomerRepository.Create(customerRequest)
	if err != nil {
		return response.BadRequest(c, err.Error(), nil)
	}
	serialized := mapper.NewCustomerMapper().Map(data)
	return response.SingleData(c, "Success Execute request", serialized)
}

func (uc *CustomerController) GetById(c echo.Context) error {
	id := c.Param("id")
	data, err := uc.CustomerRepository.GetByID(id)
	if err != nil {
		return response.BadRequest(c, err.Error(), nil)
	}
	serialized := mapper.NewCustomerMapper().Map(*data)
	return response.SingleData(c, "Success Execute request", serialized)
}

func (uc *CustomerController) Update(c echo.Context) error {
	id := c.Param("id")
	customer := request.CustomerUpdateRequest{}
	if err := c.Bind(&customer);
		err != nil {
		return response.BadRequest(c, err.Error(), nil)
	}
	data, err := uc.CustomerRepository.Update(id, customer)
	if err != nil{
		return response.BadRequest(c, err.Error(), nil)
	}
	return response.SingleData(c,"Update Success",data)
}

func (uc *CustomerController) Destroy(c echo.Context) error {
	id := c.Param("id")
	err := uc.CustomerRepository.Destroy(id)
	if err != nil{
		return response.BadRequest(c, err.Error(), nil)
	}
	return response.SingleData(c,"Success Delete Record", err)
}

func (uc *CustomerController) FindAll(c echo.Context) error {
	data,err := uc.CustomerRepository.FindAll()
	store, err := paging.NewGORMStore(config.GetLinkDb() ,&data)
	options := paging.NewOptions()
	r,_ := http.NewRequest("GET", c.Request().URL.String(), nil)

	paginator, _ := paging.NewOffsetPaginator(store, r, options)
	err = paginator.Page()
	if err != nil {
		return response.BadRequest(c, err.Error(), nil)
	}
	serialized := mapper.NewCustomerMapper().MapList(data)

	return response.Paginate(c,"Success Show Record", paginator,serialized)
}