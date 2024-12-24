package models

import (
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
	Name             string
	Description      string
	StatusString     string
	Status           TaskStatus
	Category         string
	ProgressComplete int
	ProgressTotal    int
}

func (*Task) TableName() string {
	return "task"
}

type TaskTag struct {
	TaskID  int    `gorm:"task_id"`
	TaskTag string `gorm:"task_tag"`
}

func (*TaskTag) TableName() string {
	return "task_tag"
}
