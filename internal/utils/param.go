package utils

import (
	"net/http"
	"strconv"
	"wallet-api/internal/constants"

	"github.com/gin-gonic/gin"
)

func GetIdParam(ctx *gin.Context) (string, int, error) {
	idString := ctx.Param("id")
	if idString == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id param is required"})
		return "", -1, constants.ErrIDParamNotProvided
	}
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return idString, -1, constants.ErrInvalidIDParam
	}
	return idString, id, nil
}
