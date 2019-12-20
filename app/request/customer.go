package request

import (
	"github.com/labstack/echo"
)

type CustomerRequest struct {
	Name 		string `json:"name"`
	Email 		string `json:"email"`
	Address 	string `json:"address"`
	Phone 		string `json:"phone"`
}

type CustomerUpdateRequest struct {
	Name 		string `json:"name"`
	Email 		string `json:"email"`
	Address 	string `json:"address"`
	Phone 		string `json:"phone"`
}

func (cr *CustomerRequest) Bind(c echo.Context) (*CustomerRequest, error) {
	if err := c.Bind(cr); err != nil {
		return nil, err
	}
	return cr, nil
}
