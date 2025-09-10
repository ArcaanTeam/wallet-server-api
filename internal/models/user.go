package models

import (
	"wallet-api/internal/constants"

	"gorm.io/gorm"
)

type UserRole string

const (
	RoleAdmin UserRole = "admin"
	RoleUser  UserRole = "user"
)

func NewUserRoleFromString(r string) (UserRole, error) {
	switch r {
	case "admin":
		return RoleAdmin, nil
	case "user":
		return RoleUser, nil
	default:
		return "", constants.ErrInvalidUserRoleString
	}
}

func (u UserRole) String() string {
	return string(u)
}

type User struct {
	gorm.Model
	Name         string `gorm:"not null"`
	Email        string `gorm:"type:citext;not null;unique"`
	Role         string `gorm:"type:varchar(20);not null;default:'user'"`
	PasswordHash string `gorm:"not null"`
}

func (u *User) Update(name string, roleString string) error {
	if name != "" {
		u.Name = name
	}
	if roleString != "" {
		role, err := NewUserRoleFromString(roleString)
		if err != nil {
			return err
		}
		u.Role = string(role)
	}

	return nil
}
