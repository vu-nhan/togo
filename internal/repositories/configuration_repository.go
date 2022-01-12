package repositories

import (
	"context"
	"github.com/manabie-com/togo/internal/models"
	"gorm.io/gorm"
)

type configurationRepository struct {
	db *gorm.DB
}

type ConfigurationRepository interface {
	GetAll(ctx context.Context) ([]*models.Configuration, error)
	GetById(ctx context.Context, configurationId string) (*models.Configuration, error)
	CreateConfiguration(ctx context.Context, configuration *models.Configuration) (*models.Configuration, error)
	GetCapacity(ctx context.Context, userID string, date string) (int64, error)
}

func NewConfigurationRepository(injectedDB *gorm.DB) ConfigurationRepository {
	return &configurationRepository{
		db: injectedDB,
	}
}

func (r *configurationRepository) GetAll(ctx context.Context) ([]*models.Configuration, error) {
	var configurations []*models.Configuration
	if err := r.db.Find(&configurations).Error; err != nil {
		return nil, err
	}

	return configurations, nil
}

func (r *configurationRepository) GetById(ctx context.Context, configurationId string) (*models.Configuration, error) {
	var configuration = &models.Configuration{}

	if err := r.db.Where("id = ? ", configurationId).Take(&configuration).Error; err != nil {
		return nil, err
	}

	return configuration, nil
}

func (r *configurationRepository) CreateConfiguration(ctx context.Context, configuration *models.Configuration) (*models.Configuration, error) {
	if err := r.db.Create(&configuration).Error; err != nil {
		return nil, err
	}

	return configuration, nil
}

func (r *configurationRepository) GetCapacity(ctx context.Context, userID string, date string) (int64, error) {
	var capacity = &models.Configuration{}

	if err := r.db.Where("user_id = ? and date = ?", userID, date).Take(&capacity).Error; err != nil {
		return 0, err
	}

	return capacity.Capacity, nil
}
