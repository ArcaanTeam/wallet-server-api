package utils

import "wallet-api/internal/models"

func IsValidRole(role string) bool {
	return role == models.RoleAdmin ||
		role == models.RoleUser
}
