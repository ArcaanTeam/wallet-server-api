package main

import (
	"wallet-api/internal/controllers/routes"
	"wallet-api/internal/providers"

	"github.com/gin-gonic/gin"
)

func main() {
	providers.LoadConfig()
	providers.InitDB()
	providers.Migrate()

	engine := gin.Default()

	routes.NewUserRoutes(providers.DB, engine).Setup()
	routes.NewAuthRoutes(providers.DB, engine).Setup()

	engine.Run(":8080")
}
