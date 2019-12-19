package request

import (
	"github.com/labstack/echo"
)

type CategoryRequest struct {
	Code    string `json:"code"`
	Name 	string `json:"name"`
}

type CategoryUpdateRequest struct {
	Code    string `json:"code"`
	Name 	string `json:"name"`
}

func (cr *CategoryRequest) Bind(c echo.Context) (*CategoryRequest, error) {
	if err := c.Bind(cr); err != nil {
		return nil, err
	}
	return cr, nil
}
