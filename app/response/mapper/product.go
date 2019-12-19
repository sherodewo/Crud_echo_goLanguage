package mapper

import "echo-api-starter/app/models"

type productMapper struct {
	ID       string `json:"id"`
	Name string `json:"name"`
	Description    string `json:"description"`
	Price    string `json:"price"`
	CategoryId    string `json:"category_id"`
	CategoryName string `json:"category_name"`
}

func NewProductMapper() *productMapper {
	return &productMapper{}
}
func (us *productMapper) Map(data models.Product) *productMapper {
	us.ID = data.ID
	us.Name = data.Name
	us.Description = data.Description
	us.Price = data.Price
	us.CategoryId = data.CategoryId

	return us
}

func (us *productMapper) MapList(data []models.Product) interface{} {
	var length = len(data)
	serialized := make([]productMapper, length, length)

	for k, v := range data {
		serialized[k] = productMapper{
			ID:       v.ID,
			Name: v.Name,
			Description:    v.Description,
			Price:    v.Price,
			CategoryId:    v.CategoryId,
		}
	}
	return serialized
}
