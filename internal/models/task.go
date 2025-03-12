package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	TaskCategoryMediaDiscover = "media-discover"
)

type TaskStatus int

const (
	TaskStatusInvalid  TaskStatus = iota // 未运行
	TaskStatusRunning                    // 正在运行
	TaskStatusFinished                   // 完成
	TaskStatusFailed                     // 失败
)

func (s TaskStatus) String() string {
	switch s {
	case TaskStatusRunning:
		return "running"
	case TaskStatusFinished:
		return "finished"
	case TaskStatusFailed:
		return "failed"
	default:
		return "invalid"
	}
}

type Task struct {
	gorm.Model
	Name             string     `gorm:"name"`
	Description      string     `gorm:"description"`
	StatusString     string     `gorm:"status_string"`
	Status           TaskStatus `gorm:"status"`
	Category         string     `gorm:"category"`
	ProgressComplete int        `gorm:"progress_complete"`
	ProgressTotal    int        `gorm:"progress_total"`
	Tags             []TaskTag
}

func (*Task) TableName() string {
	return "task"
}

func (t *Task) ToDTO() *TaskDTO {
	dto := &TaskDTO{}
	t.AsDTO(dto)
	return dto
}

func (t *Task) AsDTO(dto *TaskDTO) {
	dto.ID = t.ID
	dto.CreatedAt = t.CreatedAt
	dto.UpdatedAt = t.UpdatedAt
	dto.Name = t.Name
	dto.Description = t.Description
	dto.StatusString = t.StatusString
	dto.Status = t.Status
	dto.Category = t.Category
	dto.ProgressComplete = t.ProgressComplete
	dto.ProgressTotal = t.ProgressTotal
	dto.Tags = t.Tags
}

type TaskTag struct {
	TaskID  uint   `gorm:"task_id" json:"-"`
	TaskTag string `gorm:"task_tag" json:"task_tag"`
}

func (*TaskTag) TableName() string {
	return "task_tag"
}

type TaskDTO struct {
	ID               uint       `json:"id"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	Name             string     `json:"name"`
	Description      string     `json:"description"`
	StatusString     string     `json:"status_string"`
	Status           TaskStatus `json:"status"`
	Category         string     `json:"category"`
	ProgressComplete int        `json:"progress_complete"`
	ProgressTotal    int        `json:"progress_total"`
	Tags             []TaskTag  `json:"tags"`
}
