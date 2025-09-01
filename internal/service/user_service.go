package service

import (
	"errors"
	"wallet-api/internal/dto"
	"wallet-api/internal/models"
	"wallet-api/internal/repo"
	"wallet-api/internal/utils"
)

type UserService struct {
	r *repo.UserRepo
}

// TODO get interface
func NewUserService(userRepo *repo.UserRepo) *UserService {
	return &UserService{r: userRepo}
}

func (s *UserService) CreateUser(
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

func (s *UserService) UpdateUser(id string, input dto.UpdateUserInput) (*dto.UpdateUserResponse, error) {
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

func (s *UserService) GetUserByID(id string) (*dto.UserResponse, error) {
	user, err := s.r.FindByID(id)
	return &dto.UserResponse{
		ID:   user.ID,
		Name: user.Name,
	}, err
}

func (s *UserService) GetUsers() ([]*dto.UserResponse, error) {
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
