package wire

import (
	"net/http"
	"wallet-api/internal/config"
	"wallet-api/internal/container"
	"wallet-api/internal/controller/routes"

	"github.com/gin-gonic/gin"
)

type App struct {
	Engine    *gin.Engine
	Config    *config.Config
	Container *container.Container
}

func NewApp() *App {
	// Load configurations
	cfg := config.GetConfig()
	cfg.Load()

	// Initialize container
	con := container.NewContainer(cfg.DBConfig.DB)

	// Setup gin engine
	eng := gin.Default()

	return &App{
		Engine:    eng,
		Config:    cfg,
		Container: con,
	}
}

func (a *App) SetupRoutes() {
	// Health check route
	a.Engine.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	router := a.Engine.Group("/api")

	// Setup all routes
	routes.NewTestRoutes(a.Container).Setup(router.Group("/test"))
	routes.NewAuthRoutes(a.Container).Setup(router.Group("/auth"))
	routes.NewUserRoutes(a.Container).Setup(router.Group("/users"))
}

func (a *App) Run() error {
	a.SetupRoutes()
	return a.Engine.Run(a.Config.GetPortString())
}
