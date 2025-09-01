package tests

import (
	"testing"
	"wallet-api/internal/container"
	"wallet-api/internal/controller/routes"
	"wallet-api/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestDB(t *testing.T) *gorm.DB {
	testDB, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatal("failed to connect to tets database:", err)
	}

	err = testDB.AutoMigrate(&models.User{})

	if err != nil {
		t.Fatal("Failed to migrate test database:", err)
	}

	return testDB
}

func GetTestRouter(db *gorm.DB) *gin.Engine {
	gin.SetMode(gin.TestMode)
	engine := gin.New()
	router := engine.Group("/api")
	container := container.NewContainer(db)

	routes.NewUserRoutes(container).Setup(router.Group("/users"))
	routes.NewAuthRoutes(container).Setup(router.Group("/auth"))
	routes.NewTestRoutes(container).Setup(router.Group("/test"))

	return engine
}

func TeardownTestDB(t *testing.T, db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal("Failed to get generic database interface:", err.Error())
	}

	err = sqlDB.Close()
	if err != nil {
		t.Fatal("Failed to close database connection:", err)
	}
}
