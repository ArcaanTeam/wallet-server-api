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

	gin := NewGinService()

	return &App{
		Engine:    gin.Engine,
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

type GinHandler struct {
	//TODO use type
	Method  string
	Handler gin.HandlerFunc
	Path    string
}

type IGinControllerGroup interface {
	GetPrefix() string
	GetRouteHandlers() []GinHandler
	GetMiddlewares() []gin.HandlerFunc
}
type GinService struct {
	*gin.Engine
	Config interface{}
}

func NewGinService() {
	engine := gin.Default()
}

func (g *GinService) RegisterController(controller IGinControllerGroup) {
	router := g.Engine.Group(controller.GetPrefix())
	router.Use(controller.GetMiddlewares()...)
	for _, h := range controller.GetRouteHandlers() {
		router.Handle(h.Method, h.Path, h.Handler)
	}
}

func (a *App) Run() error {
	if a.Config.Flags.Migrate {
		a.Container.Migrate()
	}

	a.SetupRoutes()
	return a.Engine.Run(a.Config.GetPortString())
}
