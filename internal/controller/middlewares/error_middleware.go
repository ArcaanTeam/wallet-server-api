package middlewares

import (
	"net/http"
	"wallet-api/internal/constants"
	"wallet-api/internal/dto"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) > 0 {
			err := ctx.Errors.Last()

			// Check if the error is of type constants.ErrorType
			var statusCode int
			if customErr, ok := err.Err.(constants.ErrorType); ok {
				statusCode = mapStatusCode(customErr)
			} else {
				statusCode = http.StatusInternalServerError
			}
			ctx.JSON(statusCode, dto.NewErrorResponse(err))

			ctx.Abort()
		}
	}
}

func mapStatusCode(err constants.ErrorType) int {
	switch err {
	case constants.ErrAuthUserNotFound:
		return http.StatusNotFound
	case constants.ErrAuthUnauthorized:
		return http.StatusUnauthorized
	case constants.ErrAuthGenerateTokenFailed:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
