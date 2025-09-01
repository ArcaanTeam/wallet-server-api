package routes

import (
	"wallet-api/internal/container"
	"wallet-api/internal/controller"
	"wallet-api/internal/controller/middlewares"
	"wallet-api/internal/models"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	userController *controller.UserController
}

func NewUserRoutes(container *container.Container) *UserRoutes {
	userController := controller.NewUserController(container.UserService)
	return &UserRoutes{
		userController: userController,
	}
}

func (r *UserRoutes) Setup(router *gin.RouterGroup) {
	router.Use(middlewares.JwtAuthMiddleware(), middlewares.RBAC(models.RoleAdmin))
	{
		router.POST("", r.userController.CreateUser)
		router.GET("/:id", r.userController.GetUserByID)
		router.PUT("/:id", r.userController.UpdateUser)
	}

	protectedRoutes := router.Group("/")
	protectedRoutes.Use(middlewares.JwtAuthMiddleware())
	{
		protectedRoutes.GET("/profile", r.userController.GetProfile)
	}
}
