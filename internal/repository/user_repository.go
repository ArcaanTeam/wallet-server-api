package repository

import (
	"wallet-server-api/internal/domain/errors"
	"wallet-server-api/internal/domain/models"
)

type IUserRepo interface {
	FindAll() ([]*models.User, error)
	FindByID(id string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Create(user *models.UserCreateRequest) (*models.User, error)
	Update(id string, user *models.UserUpdateRequest) (*models.User, error)
	Delete(id string) error
}

type userRepo struct {
	// DB connection or ORM instance
}

func NewUserRepo() IUserRepo {
	return &userRepo{}
}

func (r *userRepo) FindAll() ([]*models.User, error) {
	// Database implementation
	return []*models.User{}, nil
}

func (r *userRepo) FindByID(id string) (*models.User, error) {
	// Database implementation
	if id == "not-found" {
		return nil, errors.NewNotFoundError("user not found")
	}

	return &models.User{
		ID:        id,
		FirstName: "john",
		LastName:  "doe",
		Email:     "john.doe@gmail.com",
		Age:       30,
	}, nil
}

func (r *userRepo) FindByEmail(email string) (*models.User, error) {
	// Database implementation
	return nil, nil
}

func (r *userRepo) Create(user *models.UserCreateRequest) (*models.User, error) {
	// Database implementation
	return &models.User{
		ID:        "generated-uuid",
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Age:       user.Age,
	}, nil
}

func (r *userRepo) Update(id string, user *models.UserUpdateRequest) (*models.User, error) {
	// Database implementation
	return nil, nil
}

func (r *userRepo) Delete(id string) error {
	// Database implementation
	return nil
}
