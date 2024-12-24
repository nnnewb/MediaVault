package service

import (
	"bytes"
	"fmt"
	"os/exec"

	"gitee.com/uniqptr/media-vault.git/internal/logging"
	"gitee.com/uniqptr/media-vault.git/internal/models"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type FFMPEGService struct {
	db   *gorm.DB
	path string
}

func NewFFMPEGService(db *gorm.DB, path string) *FFMPEGService {
	return &FFMPEGService{
		db:   db,
		path: path,
	}
}

// ExtractCoverInBackground 后台提取封面
func (s *FFMPEGService) ExtractCoverInBackground(src, dest string) error {
	logging.GetLogger().Info("extract cover in background", zap.String("src", src), zap.String("dest", dest))
	// create task
	task := &models.Task{
		Name:             "extract cover",
		Description:      fmt.Sprintf("extract cover from %s to %s", src, dest),
		StatusString:     "running",
		Status:           models.TaskStatusRunning,
		Category:         models.TaskCategoryMediaDiscover,
		ProgressComplete: 0,
		ProgressTotal:    1,
	}
	err := s.db.Create(task).Error
	if err != nil {
		return errors.WithStack(err)
	}

	// fire and forget
	go s.runExtractCoverTask(src, dest, task)

	return nil
}

// runExtractCoverTask 提取封面任务
func (s *FFMPEGService) runExtractCoverTask(src, dest string, task *models.Task) {
	stderr := &bytes.Buffer{}
	cmd := exec.Command(s.path, "-i", src, "-vframes", "1", dest)
	cmd.Stderr = stderr

	logging.GetLogger().Info("execute ffmpeg command", zap.String("command", cmd.String()))
	err := cmd.Run()
	if err != nil {
		logging.GetLogger().Error("extract cover failed", zap.Error(err), zap.String("stderr", stderr.String()))
		task.Status = models.TaskStatusFailed
		task.StatusString = "error: " + err.Error()
	}
	task.ProgressTotal = 1
	task.ProgressComplete = 1

	err = s.db.Save(task).Error
	if err != nil {
		logging.GetLogger().Error("save task failed", zap.Error(err))
	}
}
