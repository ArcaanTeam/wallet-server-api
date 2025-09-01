package repo

import (
	"context"

	"gorm.io/gorm"
)

type TestRepo struct {
	DB *gorm.DB
}

func NewTestRepo(db *gorm.DB) *TestRepo {
	return &TestRepo{
		DB: db,
	}
}

func (repo *TestRepo) Ping(ctx context.Context) string {
	return "pong"
}

func (r *TestRepo) Migrate() error {
	return r.DB.AutoMigrate()
}
