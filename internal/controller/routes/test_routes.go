package routes

import (
	"wallet-api/internal/container"
	"wallet-api/internal/controller"

	"github.com/gin-gonic/gin"
)

type TestRoutes struct {
	testController *controller.TestController
}

func NewTestRoutes(container *container.Container) *TestRoutes {
	testController := controller.NewTestController(container.TestService)
	return &TestRoutes{
		testController: testController,
	}
}

func (routes *TestRoutes) Setup(router *gin.RouterGroup) {
	router.GET("/ping", routes.testController.Ping)
}
