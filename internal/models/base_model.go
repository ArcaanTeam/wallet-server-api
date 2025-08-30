package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ID uuid.UUID

type BaseModel struct {
	ID        ID `gorm:"primarykey;type:uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
