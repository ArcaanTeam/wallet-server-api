package repo

import (
	"context"
	"wallet-api/internal/models"

	"gorm.io/gorm"
)

type UpdateUserInput struct{}

// TODO add context to all functions
type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepo) FindByID(ID string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, ID).Error
	return user, err
}

func (r *UserRepo) FindByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "email = ?", email).Error
	return user, err
}

func (r *UserRepo) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepo) FindUsers() ([]*models.User, error) {
	var users []*models.User
	if err := r.db.Find(&users).Error; err != nil {
		return []*models.User{}, nil
	}
	return users, nil
}

func (r *UserRepo) Migrate(ctx context.Context) error {
	return r.db.AutoMigrate(
		&models.User{},
	)
}

type IMigrateRepo interface {
	Migrate(ctx context.Context) error
}
