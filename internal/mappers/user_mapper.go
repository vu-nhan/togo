package mappers

import (
	"github.com/manabie-com/togo/internal/dtos"
	"github.com/manabie-com/togo/internal/models"
)

type UserMapper struct {

}

func NewUserMapper() UserMapper {
	return UserMapper{}
}

func (m *UserMapper) ToDto(user *models.User) *dtos.UserDto {
	return &dtos.UserDto{
		ID: user.ID,
		Password: user.Password,
	}
}

func (m *UserMapper) ToDtoList(users []*models.User) []*dtos.UserDto  {
	var userDtoList []*dtos.UserDto
	for _, user := range users {
		userDtoList = append(userDtoList, m.ToDto(user))
	}

	return userDtoList
}