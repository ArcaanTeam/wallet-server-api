package wire

import (
	"net/http"
	"wallet-api/internal/config"
	"wallet-api/internal/container"
	"wallet-api/internal/controller"

	"github.com/gin-gonic/gin"
)

type App struct {
	Engine    *controller.GinService
	Config    *config.Config
	Container *container.Container
}

func NewApp() *App {
	// Load configurations
	cfg := config.GetConfig()
	cfg.Load()

	// Initialize container
	con := container.NewContainer(cfg.DBConfig.DB)

	gin := controller.NewGinService()

	return &App{
		Engine:    gin,
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
	controller.RegisterController(router, controller.NewTestController(a.Container.TestService))
	controller.RegisterController(router, controller.NewAuthController(a.Container.AuthService))
	controller.RegisterController(router, controller.NewUserController(a.Container.UserService))
}

func (a *App) Run() error {
	if a.Config.Flags.Migrate {
		a.Container.Migrate()
	}

	a.SetupRoutes()
	return a.Engine.Run(a.Config.GetPortString())
}
