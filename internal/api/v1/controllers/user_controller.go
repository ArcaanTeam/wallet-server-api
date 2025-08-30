package controllers

import (
	"net/http"
	"wallet-server-api/internal/domain/errors"
	"wallet-server-api/internal/domain/models"
	"wallet-server-api/internal/service"

	"github.com/gin-gonic/gin"
)

type UserControllerV1 struct {
	userService service.IUserService
}

func NewUserControllerV1(service service.IUserService) *UserControllerV1 {
	return &UserControllerV1{
		userService: service,
	}
}

// GetAllUsers returns all users
// @Summary Get all users
// @Description Get a list of all users
// @Tags users-v1
// @Produce json
// @Success 200 {object} models.SuccessResponse{data=[]models.User}
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/users [get]
func (controller *UserControllerV1) GetAllUsers(ctx *gin.Context) {
	users, err := controller.userService.GetAllUsers()
	if err != nil {
		ctx.JSON(err.(*errors.AppError).Code, models.NewErrorResponse(err))
	}

	// TODO: paginate the users
	ctx.JSON(http.StatusOK, models.NewPaginatedResponse(users, "operation successful"))
}

// GetUserByID returns a user by ID
// @Summary Get user by ID
// @Description Get user details by user ID
// @Tags users-v1
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.SuccessResponse{data=models.User}
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/users/{id} [get]
func (c *UserControllerV1) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := c.userService.GetUserByID(id)
	if err != nil {
		ctx.JSON(err.(*errors.AppError).Code, models.NewErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, models.NewSuccessResponse(user))
}

// CreateUser creates a new user
// @Summary Create a new user
// @Description Create a new user with the provided data
// @Tags users-v1
// @Accept json
// @Produce json
// @Param user body models.UserCreateRequest true "User data"
// @Success 201 {object} models.SuccessResponse{data=models.User}
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/users [post]
func (c *UserControllerV1) CreateUser(ctx *gin.Context) {
	var userReq models.UserCreateRequest
	if err := ctx.ShouldBindJSON(&userReq); err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewErrorResponse(
			errors.NewBadRequestError("invalid request body"),
		))
		return
	}

	user, err := c.userService.CreateUser(&userReq)
	if err != nil {
		ctx.JSON(err.(*errors.AppError).Code, models.NewErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, models.NewSuccessResponse(user))
}

// UpdateUser updates an existing user
// @Summary Update user
// @Description Update an existing user with the provided data
// @Tags users-v1
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body models.UserUpdateRequest true "User data"
// @Success 200 {object} models.SuccessResponse{data=models.User}
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/users/{id} [put]
func (c *UserControllerV1) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var userReq models.UserUpdateRequest
	if err := ctx.ShouldBindJSON(&userReq); err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewErrorResponse(
			errors.NewBadRequestError("invalid request body"),
		))
		return
	}

	user, err := c.userService.UpdateUser(id, &userReq)
	if err != nil {
		ctx.JSON(err.(*errors.AppError).Code, models.NewErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, models.NewSuccessResponse(user))
}

// DeleteUser deletes a user
// @Summary Delete user
// @Description Delete a user by ID
// @Tags users-v1
// @Produce json
// @Param id path string true "User ID"
// @Success 204
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/users/{id} [delete]
func (c *UserControllerV1) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.userService.DeleteUser(id)
	if err != nil {
		ctx.JSON(err.(*errors.AppError).Code, models.NewErrorResponse(err))
		return
	}
	ctx.Status(http.StatusNoContent)
}
