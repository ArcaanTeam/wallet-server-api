package container

import (
	"wallet-api/internal/repo"
	"wallet-api/internal/service"

	"gorm.io/gorm"
)

type Container struct {
	DB *gorm.DB

	// repo
	TestRepo repo.TestRepo
	AuthRepo repo.AuthRepo
	UserRepo repo.UserRepo

	// service
	TestService service.TestService
	AuthService service.AuthService
	UserService service.UserService
}

func NewContainer(db *gorm.DB) *Container {
	testRepo := repo.NewTestRepo(db)
	userRepo := repo.NewUserRepo(db)
	authRepo := repo.NewAuthRepo(db)

	testService := service.NewTestService(testRepo)
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(authRepo)

	return &Container{
		DB:       db,
		TestRepo: *testRepo,
		AuthRepo: *authRepo,
		UserRepo: *userRepo,

		TestService: *testService,
		AuthService: *authService,
		UserService: *userService,
	}
}
