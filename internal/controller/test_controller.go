package controller

import (
	"net/http"
	"wallet-api/internal/constants"
	"wallet-api/internal/dto"
	"wallet-api/internal/service"

	"github.com/gin-gonic/gin"
)

type TestController struct {
	service service.TestService
}

func NewTestController(service service.TestService) *TestController {
	return &TestController{
		service: service,
	}
}
func (c *TestController) GetControllerGroups() []GinControllerGroup {
	return []GinControllerGroup{
		{
			Prefix:      "/test",
			Middlewares: nil,
			RouteHandlers: []GinHandler{
				{"GET", "/ping", c.Ping},
			},
		},
	}
}

func (c *TestController) Ping(ctx *gin.Context) {
	if c.service.Ping(ctx.Request.Context()) == "" {
		ctx.Error(constants.ErrInternal)
		ctx.Next()
	} else {
		ctx.JSON(http.StatusOK, dto.NewSuccessResponse(gin.H{"message": "pong"}))
	}
}
