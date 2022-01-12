// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package providers

import (
	"github.com/manabie-com/togo/internal/handlers"
	"github.com/manabie-com/togo/internal/helpers"
	"github.com/manabie-com/togo/internal/mappers"
	"github.com/manabie-com/togo/internal/middlewares"
	"github.com/manabie-com/togo/internal/repositories"
	"github.com/manabie-com/togo/internal/services"
	"gorm.io/gorm"
)

// Injectors from provider.go:

func ProvideApplicationMiddleware(jwtSecretKey string) *middlewares.AppMiddleware {
	tokenProvider := helpers.NewTokenProvider(jwtSecretKey)
	appMiddleware := middlewares.NewAppMiddleware(tokenProvider)
	return appMiddleware
}

func ProvideTaskHandler(db *gorm.DB) *handlers.TaskHandler {
	configurationRepository := repositories.NewConfigurationRepository(db)
	configurationService := services.NewConfigurationService(configurationRepository)
	taskRepository := repositories.NewTaskRepository(db)
	taskService := services.NewTaskService(configurationService, taskRepository)
	taskHandler := handlers.NewTaskHandler(taskService)
	return taskHandler
}

func ProvideUserHandler(db *gorm.DB, jwtSecretKey string) *handlers.UserHandler {
	tokenProvider := helpers.NewTokenProvider(jwtSecretKey)
	userRepository := repositories.NewUserRepository(db)
	userMapper := mappers.NewUserMapper()
	userService := services.NewUserService(tokenProvider, userRepository, userMapper)
	userHandler := handlers.NewUserHandler(userService)
	return userHandler
}

func ProvideConfigurationHandler(db *gorm.DB) *handlers.ConfigurationHandler {
	configurationRepository := repositories.NewConfigurationRepository(db)
	configurationService := services.NewConfigurationService(configurationRepository)
	configurationHandler := handlers.NewConfigurationHandler(configurationService)
	return configurationHandler
}
