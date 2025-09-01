package service

import (
	"context"
	"strconv"
	"wallet-api/internal/constants"
	"wallet-api/internal/dto"
	"wallet-api/internal/models"
	"wallet-api/internal/repo"
	"wallet-api/internal/utils"
)

type AuthService struct {
	r *repo.AuthRepo
}

func NewAuthService(repo *repo.AuthRepo) *AuthService {
	return &AuthService{r: repo}
}

func (s *AuthService) Login(ctx context.Context, input dto.LoginInput) (string, *models.User, constants.ErrorType) {
	user, err := s.r.GetUserByEmail(ctx, input.Email)
	if err != nil {
		// TODO: handle database errors separately
		return "", nil, constants.ErrAuthUserNotFound
	}

	if err := utils.CheckPasswordHash(input.Password, user.PasswordHash); err != nil {
		return "", nil, constants.ErrAuthUnauthorized
	}

	token, err := utils.GenerateJwtToken(strconv.Itoa(int(user.ID)), user.Email, user.Role)
	if err != nil {
		return "", nil, constants.ErrAuthGenerateTokenFailed
	}

	return token, user, constants.ErrNone
}
