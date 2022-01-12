package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/manabie-com/togo/internal/constants"
	"github.com/manabie-com/togo/internal/dtos"
	"github.com/manabie-com/togo/internal/services"
	"net/http"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(injectedUserService services.UserService) *UserHandler {
	return &UserHandler{
		userService: injectedUserService,
	}
}

// GetAll godoc
// @Summary Get All Users
// @Description Get All Users
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {object} dtos.ApiResponse
// @Router /users/ [get]
func (h *UserHandler) GetAll(ctx *gin.Context) {
	users, err := h.userService.GetAllUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &dtos.ApiResponse{Data: nil})
		return
	}

	ctx.JSON(http.StatusOK, &dtos.ApiResponse{Data: users})
}

// GetAll godoc
// @Summary Get User By Id
// @Description Get By Id
// @Param id path string true "User Id"
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {object} dtos.ApiResponse
// @Router /users/{id} [get]
func (h *UserHandler) GetById(ctx *gin.Context) {
	userId := ctx.Param("id")

	user, err := h.userService.GetById(ctx, userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &dtos.ApiResponse{Data: nil})
		return
	}
	ctx.JSON(http.StatusOK, &dtos.ApiResponse{Data: user})
}

// GetAuthToken godoc
// @Summary Get Authentication Token
// @Description Get Authentication Token
// @Param user_id query string true "User Id"
// @Param password query string true "Password"
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {object} dtos.TokenResponse
// @Router /login [get]
func (h *UserHandler) GetAuthToken(ctx *gin.Context) {
	userId := ctx.Query("user_id")
	password := ctx.Query("password")

	token, err := h.userService.GetAuthToken(ctx, userId, password)
	if errors.Is(err, constants.ErrIncorrectUserIdOrPassword) {
		ctx.JSON(http.StatusUnauthorized, dtos.NewError(err))
		return
	}

	if errors.Is(err, constants.ErrCreateToken) {
		ctx.JSON(http.StatusInternalServerError, dtos.NewError(err))
		return
	}

	ctx.JSON(http.StatusOK, &dtos.TokenResponse{Data: token})
}
