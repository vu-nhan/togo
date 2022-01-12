package services

import (
	"context"
	"github.com/manabie-com/togo/internal/constants"
	"github.com/manabie-com/togo/internal/dtos"
	"github.com/manabie-com/togo/internal/helpers"
	"github.com/manabie-com/togo/internal/mappers"
	"github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/internal/repositories"
	"github.com/sirupsen/logrus"
)

type userService struct {
	tokenProvider  helpers.TokenProvider
	userRepository repositories.UserRepository
	userMapper 	   mappers.UserMapper
}

type UserService interface {
	GetAllUser(ctx context.Context) ([]*dtos.UserDto, error)
	GetById(ctx context.Context, userId string) (*models.User, error)
	GetAuthToken(ctx context.Context, userId, password string) (string, error)
}

func NewUserService(injectedTokenProvider helpers.TokenProvider,
	injectedUserRepository repositories.UserRepository,
	injectedUserMapper mappers.UserMapper) UserService {
	return &userService{
		tokenProvider:  injectedTokenProvider,
		userRepository: injectedUserRepository,
		userMapper: injectedUserMapper,
	}
}

func (s *userService) GetAllUser(ctx context.Context) ([]*dtos.UserDto, error) {
	users, err := s.userRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return s.userMapper.ToDtoList(users), nil
}

func (s *userService) GetById(ctx context.Context, userId string) (*models.User, error) {
	return s.userRepository.GetById(ctx, userId)
}

func (s *userService) GetAuthToken(ctx context.Context, userId, password string) (string, error) {
	if !s.userRepository.ValidateUser(ctx, userId, password) {
		logrus.Errorf("Validating User not match :%s", userId)
		return "", constants.ErrIncorrectUserIdOrPassword
	}

	token, err := s.tokenProvider.CreateToken(userId)
	if err != nil {
		logrus.Errorf("CreateToken error: %s", err.Error())
		return "", constants.ErrCreateToken
	}

	return token, nil
}
