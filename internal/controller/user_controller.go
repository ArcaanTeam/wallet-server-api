package controller

import (
	"net/http"
	"wallet-api/internal/dto"
	"wallet-api/internal/service"
	"wallet-api/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	s service.UserService
}

type GinHandler struct {
	//TODO use type
	Method  string
	Handler gin.HandlerFunc
	Path    string
}

type IGinControllerGroup interface {
	GetPrefix() string
	GetRouteHandlers() []GinHandler
	GetMiddlewares() []gin.HandlerFunc
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{s: service}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var input dto.CreateUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "malformed request body",
			"details": err.Error(),
		})
	}

	createdUser, err := c.s.CreateUser(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "unable to create user",
			"details": err.Error(),
		})
	}

	ctx.JSON(http.StatusCreated, createdUser)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	userIdString, _, err := utils.GetIdParam(ctx)
	if err != nil {
		return
	}

	var input dto.UpdateUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	updatedUser, err := c.s.UpdateUser(userIdString, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "unable to update user",
			"details": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, updatedUser)
}

func (c *UserController) GetUserByID(ctx *gin.Context) {
	idString, _, err := utils.GetIdParam(ctx)
	if err != nil {
		return
	}

	user, err := c.s.GetUserByID(idString)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error":   "user not found",
				"details": "",
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   "internal server error",
				"details": err.Error(),
			})
		}
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) GetProfile(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}

	userIdString, ok := userID.(string)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid user id",
			"details": "",
		})
	}
	user, err := c.s.GetUserByID(userIdString)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error":   "user not found",
				"details": "",
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   "internal server error",
				"details": err.Error(),
			})
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":   user.ID,
			"name": user.Name,
		},
	})
}

func (c *UserController) GetUsers(ctx *gin.Context) {
	users, err := c.s.GetUsers()
	if err != nil {
		// TODO: handle not-found branch separately
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
