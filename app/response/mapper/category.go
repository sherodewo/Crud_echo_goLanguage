package mapper

import "echo-api-starter/app/models"

type categoryMapper struct {
	ID       string `json:"id"`
	Code    string `json:"code"`
	Name string `json:"name"`
}

func NewCategoryMapper() *categoryMapper {
	return &categoryMapper{}
}
func (us *categoryMapper) Map(data models.Category) *categoryMapper {
	us.ID = data.ID
	us.Code = data.Code
	us.Name = data.Name

	return us
}

func (us *categoryMapper) MapList(data []models.Category) interface{} {
	var length = len(data)
	serialized := make([]categoryMapper, length, length)

	for k, v := range data {
		serialized[k] = categoryMapper{
			ID:		v.ID,
			Code: 	v.Code,
			Name: 	v.Name,
		}
	}
	return serialized
}
