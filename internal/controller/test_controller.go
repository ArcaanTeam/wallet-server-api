package controller

import (
	"errors"
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

func (c *TestController) Ping(ctx *gin.Context) {
	if c.service.Ping() == "" {
		ctx.Error(errors.New(constants.ErrInternal))
		ctx.Next()
	} else {
		ctx.JSON(http.StatusOK, dto.NewSuccessResponse(gin.H{"message": "ping"}))
	}
}
