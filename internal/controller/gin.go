package controller

import "github.com/gin-gonic/gin"

type GinService struct {
	*gin.Engine
}

func NewGinService() *GinService {
	engine := gin.Default()
	return &GinService{
		engine,
	}
}

type GinHandler struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}
type GinControllerGroup struct {
	Prefix        string
	Middlewares   []gin.HandlerFunc
	RouteHandlers []GinHandler
}
type IGinController interface {
	GetControllerGroups() []GinControllerGroup
}

func RegisterController(router *gin.RouterGroup, controller IGinController) {
	for _, controllerGroup := range controller.GetControllerGroups() {
		group := router.Group(controllerGroup.Prefix)
		group.Use(controllerGroup.Middlewares...)
		for _, h := range controllerGroup.RouteHandlers {
			group.Handle(h.Method, h.Path, h.Handler)
		}
	}
}
