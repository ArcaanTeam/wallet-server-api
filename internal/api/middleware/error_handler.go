package middleware

import (
	"net/http"
	"wallet-server-api/internal/domain/models"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) > 0 {
			err := ctx.Errors.Last()
			ctx.JSON(http.StatusInternalServerError, models.NewErrorResponse(err))
			ctx.Abort()
		}
	}
}
