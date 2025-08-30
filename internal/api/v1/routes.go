package v1

import (
	"wallet-server-api/internal/api/v1/controllers"
	"wallet-server-api/internal/repository"
	"wallet-server-api/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup) {
	// Initialize deps
	userRepo := repository.NewUserRepo()
	userService := service.NewUserService(userRepo)
	userController := controllers.NewUserControllerV1(userService)

	users := router.Group("/users")
	{
		users.GET("", userController.GetAllUsers)
		users.GET("/:id", userController.GetUserByID)
		users.POST("", userController.CreateUser)
		users.PUT("/:id", userController.UpdateUser)
		users.DELETE("/:id", userController.DeleteUser)
	}
}
