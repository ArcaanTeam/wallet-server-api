package models

import (
	"gorm.io/gorm"
)

const (
	RoleAdmin string = "admin"
	RoleUser  string = "user"
)

type User struct {
	gorm.Model
	Name         string `gorm:"not null"`
	Email        string `gorm:"type:citext;not null;unique"`
	Role         string `gorm:"type:varchar(20);not null;default:'user'"`
	PasswordHash string `gorm:"not null"`
}
