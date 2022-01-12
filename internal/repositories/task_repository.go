package repositories

import (
	"context"
	"github.com/manabie-com/togo/internal/models"
	"gorm.io/gorm"
 )

type taskRepository struct {
	db *gorm.DB
}

type TaskRepository interface {
	GetAll(ctx context.Context) ([]*models.Task, error)
	GetById(ctx context.Context, taskId string) (*models.Task, error)
	Create(ctx context.Context, task *models.Task) (*models.Task, error)
	GetTasks(ctx context.Context, userID string, createdDate string) ([]*models.Task, error)
	CountTotalInDate(ctx context.Context, userID string, createdDate string) (int64, error)
}

func NewTaskRepository(injectedDB *gorm.DB) TaskRepository {
	return &taskRepository{
		db: injectedDB,
	}
}

func (r *taskRepository) GetAll(ctx context.Context) ([]*models.Task, error) {
	var tasks []*models.Task

	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *taskRepository) GetById(ctx context.Context, taskId string) (*models.Task, error) {
	var task = &models.Task{}

	if err := r.db.Where("id = ? ", taskId).Take(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (r *taskRepository) Create(ctx context.Context, task *models.Task) (*models.Task, error) {
	if err := r.db.Create(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (r *taskRepository) GetTasks(ctx context.Context, userID string, createdDate string) ([]*models.Task, error) {
	var tasks []*models.Task

	if err := r.db.Model(&models.Task{}).
		Where("user_id = ? AND created_date = ?", userID, createdDate).
		Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *taskRepository) CountTotalInDate(ctx context.Context, userID string, createdDate string) (int64, error) {
	var total int64
	if err := r.db.Model(&models.Task{}).
		Where("user_id = ? AND created_date = ?", userID, createdDate).
		Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil

}
