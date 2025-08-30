package controllers

import (
	"net/http"
	"wallet-api/internal/constants"
	"wallet-api/internal/dto"
	"wallet-api/internal/services"

	"github.com/gin-gonic/gin"
)

type authController struct {
	s services.AuthService
}

func NewAuthController(service services.AuthService) *authController {
	return &authController{s: service}
}

func (c *authController) Login(ctx *gin.Context) {
	var input dto.LoginInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, user, err := c.s.Login(input)
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
