package request

import (
	"github.com/labstack/echo"
)

type ProductRequest struct {
	Name string `json:"name"`
	Description    string `json:"description"`
	Price string `json:"price"`
	CategoryId string `json:"category_id"`
}

type ProductUpdateRequest struct {
	Name string `json:"name"`
	Description    string `json:"description"`
	Price string `json:"price"`
	CategoryId string `json:"category_id"`
}

func (cr *ProductRequest) Bind(c echo.Context) (*ProductRequest, error) {
	if err := c.Bind(cr); err != nil {
		return nil, err
	}
	return cr, nil
}
