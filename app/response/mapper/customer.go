package mapper

import "echo-api-starter/app/models"

type customerMapper struct {
	ID       string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Address string `json:"address"`
	Phone string `json:"phone"`
}

func NewCustomerMapper() *customerMapper {
	return &customerMapper{}
}
func (us *customerMapper) Map(data models.Customer) *customerMapper {
	us.ID = data.ID
	us.Name = data.Name
	us.Email = data.Email
	us.Address = data.Address
	us.Phone = data.Phone

	return us
}

func (us *customerMapper) MapList(data []models.Customer) interface{} {
	var length = len(data)
	serialized := make([]customerMapper, length, length)

	for k, v := range data {
		serialized[k] = customerMapper{
			ID:			v.ID,
			Name: 		v.Name,
			Email: 		v.Email,
			Address: 	v.Address,
			Phone: 		v.Phone,
		}
	}
	return serialized
}
