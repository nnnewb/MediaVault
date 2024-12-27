package service

import (
	"github.com/nnnewb/media-vault/internal/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type TaskService struct {
	db *gorm.DB
}

func NewTaskService(db *gorm.DB) *TaskService {
	return &TaskService{db: db}
}

func (t *TaskService) Advance(id uint, progress int) error {
	err := t.db.
		Model(&models.Task{}).
		Where("id = ?", id).
		Updates(map[string]any{
			"progress_complete": gorm.Expr("progress_complete + ?", progress),
			"status":            gorm.Expr("iif(progress_complete + ? >= progress_total, ?, status)", progress, models.TaskStatusFinished),
			"status_string":     gorm.Expr("iif(progress_complete + ? >= progress_total, ?, status_string)", progress, models.TaskStatusFinished.String()),
		}).
		Error
	return errors.WithStack(err)
}

func (t *TaskService) FinishTask(id uint) error {
	err := t.db.
		Model(&models.Task{}).
		Where("id = ?", id).
		Updates(map[string]any{
			"status":            models.TaskStatusFinished,
			"status_string":     models.TaskStatusFinished.String(),
			"progress_complete": gorm.Expr("progress_total"),
		}).
		Error
	return errors.WithStack(err)
}

func (t *TaskService) FailTask(id uint) error {
	err := t.db.
		Model(&models.Task{}).
		Where("id = ?", id).
		Updates(map[string]any{
			"status":        models.TaskStatusFailed,
			"status_string": models.TaskStatusFailed.String(),
		}).
		Error
	return errors.WithStack(err)
}

func (t *TaskService) CreateTask(name, description, category string, tags []string, total int) (*models.Task, error) {
	task := &models.Task{
		Name:             name,
		Description:      description,
		StatusString:     models.TaskStatusRunning.String(),
		Status:           models.TaskStatusRunning,
		Category:         category,
		ProgressComplete: 0,
		ProgressTotal:    total,
	}
	err := t.db.Transaction(func(tx *gorm.DB) error {
		err := t.db.Create(task).Error
		if err != nil {
			return errors.WithStack(err)
		}

		for _, tag := range tags {
			tt := &models.TaskTag{
				TaskID:  task.ID,
				TaskTag: tag,
			}
			err = t.db.Create(tt).Error
			if err != nil {
				return errors.WithStack(err)
			}
		}

		return nil
	})
	return task, errors.WithStack(err)
}
