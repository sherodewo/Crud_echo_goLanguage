package mapper

import "echo-api-starter/app/models"

type userMapper struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func NewUserMapper() *userMapper {
	return &userMapper{}
}
func (us *userMapper) Map(data models.User) *userMapper {
	us.ID = data.ID
	us.Username = data.Username
	us.Email = data.Email

	return us
}

func (us *userMapper) MapList(data []models.User) interface{} {
	var length = len(data)
	serialized := make([]userMapper, length, length)

	for k, v := range data {
		serialized[k] = userMapper{
			ID:       v.ID,
			Username: v.Username,
			Email:    v.Email,
		}
	}
	return serialized
}
