package services

import (
	"errors"
	"wallet-api/internal/dto"
	"wallet-api/internal/models"
	"wallet-api/internal/repositories"
	"wallet-api/internal/utils"
)

type IUserService interface {
	CreateUser(input dto.CreateUserInput) (*dto.CreateUserResponse, error)
	GetUserByID(id string) (*dto.UserResponse, error)
	GetUsers() ([]*dto.UserResponse, error)
	UpdateUser(id string, input dto.UpdateUserInput) (*dto.UpdateUserResponse, error)
}

type userService struct {
	r repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) IUserService {
	return &userService{r: userRepo}
}

func (s *userService) CreateUser(
	input dto.CreateUserInput,
) (*dto.CreateUserResponse, error) {
	if len(input.Password) < 6 {
		return nil, errors.New("password must be at least 6 characters long")
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	var role string
	if !utils.IsValidRole(role) || input.Role == "" {
		role = models.RoleUser // Default role
	} else {
		role = input.Role
	}

	user := models.User{
		Name:         input.Name,
		Email:        input.Email,
		PasswordHash: hashedPassword,
		Role:         role,
	}

	if err := s.r.Create(&user); err != nil {
		return nil, err
	}
	return &dto.CreateUserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}, nil
}

func (s *userService) UpdateUser(id string, input dto.UpdateUserInput) (*dto.UpdateUserResponse, error) {
	userToUpdate, err := s.r.FindByID(id)
	if err != nil {
		return nil, err
	}

	if input.Name != "" {
		userToUpdate.Name = input.Name
	}
	if input.Email != "" {
		userToUpdate.Email = input.Email
	}
	if input.Role != "" && utils.IsValidRole(input.Role) {
		userToUpdate.Role = input.Role
	}
	if input.Password != "" && len(input.Password) > 5 {
		hashedNewPassword, err := utils.HashPassword(input.Password)
		if err != nil {
			return nil, err
		}
		userToUpdate.PasswordHash = hashedNewPassword
	}

	if err := s.r.Update(&userToUpdate); err != nil {
		return nil, err
	}
	return &dto.UpdateUserResponse{
		ID:    userToUpdate.ID,
		Name:  userToUpdate.Name,
		Email: userToUpdate.Email,
		Role:  userToUpdate.Role,
	}, nil
}

func (s *userService) GetUserByID(id string) (*dto.UserResponse, error) {
	user, err := s.r.FindByID(id)
	return &dto.UserResponse{
		ID:   user.ID,
		Name: user.Name,
	}, err
}

func (s *userService) GetUsers() ([]*dto.UserResponse, error) {
	users, err := s.r.FindUsers()
	if err != nil {
		return []*dto.UserResponse{}, nil
	}
	response := make([]*dto.UserResponse, len(users))
	for i := range users {
		response[i] = &dto.UserResponse{
			ID:   users[i].ID,
			Name: users[i].Name,
		}
	}
	return response, nil
}
