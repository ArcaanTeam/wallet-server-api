package services

import (
	"errors"
	"strconv"
	"wallet-api/internal/constants"
	"wallet-api/internal/dto"
	"wallet-api/internal/models"
	"wallet-api/internal/repositories"
	"wallet-api/internal/utils"
)

type AuthService interface {
	Login(input dto.LoginInput) (
		string,
		*models.User,
		error,
	)
}

type authService struct {
	r repositories.AuthRepository
}

func NewAuthService(repo repositories.AuthRepository) AuthService {
	return &authService{r: repo}
}

func (s *authService) Login(input dto.LoginInput) (string, *models.User, error) {
	user, err := s.r.GetUserByEmail(input.Email)
	if err != nil {
		// TODO: handle database errors separately
		return "", nil, errors.New(constants.ErrAuthUserNotFound)
	}

	if err := utils.CheckPasswordHash(input.Password, user.PasswordHash); err != nil {
		return "", nil, errors.New(constants.ErrAuthUnauthorized)
	}

	token, err := utils.GenerateJwtToken(strconv.Itoa(int(user.ID)), user.Email, user.Role)
	if err != nil {
		return "", nil, errors.New(constants.ErrAuthGenerateTokenFailed)
	}

	return token, user, nil
}
