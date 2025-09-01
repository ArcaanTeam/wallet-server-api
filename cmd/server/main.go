package main

import (
	"wallet-api/internal/config"
	"wallet-api/internal/container"
	"wallet-api/internal/controller/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.GetConfig()
	cfg.Load()

	// TODO: move migration to each model
	// config.DBConfig.InitDB

	container := container.NewContainer(cfg.DBConfig.DB)
	engine := gin.Default()
	router := engine.Group("/api")

	routes.NewUserRoutes(container).Setup(router.Group("/users"))
	routes.NewAuthRoutes(container).Setup(router.Group("/auth"))
	routes.NewTestRoutes(container).Setup(router.Group("/test"))

	engine.Run(cfg.GetPortString())
}
