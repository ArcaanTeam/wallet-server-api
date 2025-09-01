package repo

import "gorm.io/gorm"

type TestRepo struct {
	DB *gorm.DB
}

func NewTestRepo(db *gorm.DB) *TestRepo {
	return &TestRepo{
		DB: db,
	}
}

func (repo *TestRepo) Ping() string {
	return "pong"
}
