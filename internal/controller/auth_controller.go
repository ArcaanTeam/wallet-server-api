package controller

import (
	"net/http"
	"wallet-api/internal/constants"
	"wallet-api/internal/dto"
	"wallet-api/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	s service.AuthService
}

func NewAuthController(service service.AuthService) *AuthController {
	return &AuthController{s: service}
}

func (c *AuthController) Login(ctx *gin.Context) {
	stdCtx := ctx.Request.Context()

	var input dto.LoginInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, user, err := c.s.Login(stdCtx, input)
	// FIXME: move error handling to error middleware
	if err != constants.ErrNone {
		ctx.Error(err)
		ctx.Next()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"id":    user.ID,
			"email": user.Email,
			"name":  user.Name,
			"role":  user.Role,
		},
	})
}
