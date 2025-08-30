package main

import (
	"log"
	"wallet-server-api/internal/api"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := api.SetupRouter()

	log.Println("Server starting on :8080 ...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
