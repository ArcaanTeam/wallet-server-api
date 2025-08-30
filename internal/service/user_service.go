package service

import (
	"wallet-server-api/internal/domain/errors"
	"wallet-server-api/internal/domain/models"
	"wallet-server-api/internal/repository"
)

type IUserService interface {
	GetAllUsers() ([]*models.User, error)
	GetUserByID(id string) (*models.User, error)
	CreateUser(user *models.UserCreateRequest) (*models.User, error)
	UpdateUser(id string, user *models.UserUpdateRequest) (*models.User, error)
	DeleteUser(id string) error
}

type userService struct {
	userRepo repository.IUserRepo
}

func NewUserService(repo repository.IUserRepo) *userService {
	return &userService{
		userRepo: repo,
	}
}

func (service *userService) GetAllUsers() ([]*models.User, error) {
	return service.userRepo.FindAll()
}

func (service *userService) GetUserByID(id string) (*models.User, error) {
	return service.userRepo.FindByID(id)
}

func (service *userService) CreateUser(user *models.UserCreateRequest) (*models.User, error) {
	if user.Age < 18 {
		return nil, errors.NewBadRequestError("user must be at least 18 years old")
	}
	if existingUser, _ := service.userRepo.FindByEmail(user.Email); existingUser != nil {
		return nil, errors.NewBadRequestError("email already exists")
	}

	return service.userRepo.Create(user)
}

func (service *userService) UpdateUser(id string, user *models.UserUpdateRequest) (*models.User, error) {
	_, err := service.userRepo.FindByID(id)
	if err != nil {
		return nil, errors.NewNotFoundError("user not found")
	}

	return service.userRepo.Update(id, user)
}

func (service *userService) DeleteUser(id string) error {
	if _, err := service.userRepo.FindByID(id); err != nil {
		return errors.NewNotFoundError("user not found")
	}

	return service.userRepo.Delete(id)
}
