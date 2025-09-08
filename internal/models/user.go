package models

import (
	"gorm.io/gorm"
)

type UserRole string

const (
	RoleAdmin UserRole = "admin"
	RoleUser  UserRole = "user"
)

func NewUserRoleFromString(r string) (UserRole, error) {
	// TODO implement it
	switch r {
	case "admin":
		return RoleAdmin, nil
	}
	return "", nil
}

func (u *UserRole) CheckRole() bool {
	// TODO implement it
	return false
}

type User struct {
	gorm.Model
	Name         string `gorm:"not null"`
	Email        string `gorm:"type:citext;not null;unique"`
	Role         string `gorm:"type:varchar(20);not null;default:'user'"`
	PasswordHash string `gorm:"not null"`
}

func NewUser() {

	return

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
