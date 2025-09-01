package routes

import (
	"wallet-api/internal/container"
	"wallet-api/internal/controller"

	"github.com/gin-gonic/gin"
)

type AuthRoutes struct {
	authController *controller.AuthController
}

func NewAuthRoutes(container *container.Container) *AuthRoutes {
	authController := controller.NewAuthController(container.AuthService)
	return &AuthRoutes{
		authController: authController,
	}
}

func (r *AuthRoutes) Setup(router *gin.RouterGroup) {
	router.POST("/login", r.authController.Login)
}
