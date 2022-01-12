package services

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/manabie-com/togo/internal/mappers"
	"github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/tests/unit-test/services/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	UserId   = "USER-ID-HERE"
	Password = "USER-PASSWORD-HERE"
)

var (
	UserId1   = "USER-ID-1"
	Password1 = "PASSWORD-1"
	UserId2   = "USER-ID-2"
	Password2 = "PASSWORD-2"
	UserId3   = "USER-ID-3"
	Password3 = "PASSWORD-3"
)

func TestGetAllUsers(t *testing.T) {
	mockController := gomock.NewController(t)
	mockUserRepository := mocks.NewMockUserRepository(mockController)
	mockingUser := []*models.User{
		{ID: UserId1, Password: Password1},
		{ID: UserId2, Password: Password2},
		{ID: UserId3, Password: Password3},
	}

	mockUserRepository.EXPECT().
		GetAll(context.Background()).Return(mockingUser, nil).Times(1)

	userService := services.NewUserService(nil, mockUserRepository, mappers.NewUserMapper())
	users, err := userService.GetAllUser(context.Background())
	assert.Equal(t, err, nil)
	assert.Equal(t, len(users), 3)
}

func TestGetUserById(t *testing.T) {
	mockController := gomock.NewController(t)
	mockUserRepository := mocks.NewMockUserRepository(mockController)
	mockingUser := &models.User{ID: UserId, Password: Password}

	mockUserRepository.EXPECT().
		GetById(context.Background(), mockingUser.ID).
		Return(mockingUser, nil).Times(1)

	userService := services.NewUserService(nil, mockUserRepository, mappers.NewUserMapper())

	user, err := userService.GetById(context.Background(), mockingUser.ID)
	assert.Equal(t, err, nil)
	assert.Equal(t, user.ID, mockingUser.ID)
	assert.Equal(t, user.Password, mockingUser.Password)
}
