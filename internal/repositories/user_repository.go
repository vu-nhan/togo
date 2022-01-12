package repositories

import (
	"context"
	"github.com/manabie-com/togo/internal/models"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	GetAll(ctx context.Context) ([]*models.User, error)
	GetById(ctx context.Context, userId string) (*models.User, error)
	ValidateUser(ctx context.Context, userID, password string) bool
}

func NewUserRepository(injectedDB *gorm.DB) UserRepository {
	return &userRepository{
		db: injectedDB,
	}
}

func (r *userRepository) GetAll(ctx context.Context) ([]*models.User, error) {
	var users []*models.User

	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) GetById(ctx context.Context, userId string) (*models.User, error) {
	var task = &models.User{}

	if err := r.db.Where("id = ? ", userId).Take(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (r *userRepository) ValidateUser(ctx context.Context, userID, password string) bool {
	var user = &models.User{}
	if err := r.db.Model(user).
		Where("id = ? AND password = ?", userID, password).Error; err != nil {
		return false
	}

	return true
}
