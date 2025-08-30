package api

import (
	"net/http"
	"wallet-server-api/internal/api/middleware"
	v1 "wallet-server-api/internal/api/v1"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Middlewares
	router.Use(middleware.ErrorHandler())
	// router.Use(middleware.CORS())
	// router.Use(middleware.Logger())

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	// API versioning
	api := router.Group("/api")
	{
		apiV1 := api.Group("/v1")
		v1.SetupRoutes(apiV1)
	}

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
