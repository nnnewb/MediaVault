package service

import (
	"fmt"
	"path/filepath"

	"gitee.com/uniqptr/media-vault.git/internal/constants"
	"gitee.com/uniqptr/media-vault.git/internal/logging"
	"gitee.com/uniqptr/media-vault.git/internal/models"
	"go.uber.org/zap"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type MediaService struct {
	db        *gorm.DB
	infer     *MediaInfer
	ff        *FFMPEGService
	coverRoot string
}

func NewMediaService(db *gorm.DB, dataRoot string, infer *MediaInfer, ff *FFMPEGService) *MediaService {
	return &MediaService{
		db:        db,
		infer:     infer,
		ff:        ff,
		coverRoot: filepath.Join(dataRoot, constants.DataFolderCovers),
	}
}

func (s *MediaService) List(options ...QueryOption) ([]*models.Media, error) {
	var medias []*models.Media
	err := s.db.Transaction(func(tx *gorm.DB) error {
		tx = tx.Model(&models.Media{})
		for _, option := range options {
			option(tx)
		}

		return tx.Preload("MediaCover").Find(&medias).Error
	})
	return medias, errors.WithStack(err)
}

func (s *MediaService) Add(paths ...string) ([]*models.Media, error) {
	var medias []*models.Media
	for _, path := range paths {
		medias = append(medias, &models.Media{
			Path:      path,
			MediaType: s.infer.InferMediaType(path),
		})
	}
	err := s.db.Transaction(func(tx *gorm.DB) error {
		return tx.CreateInBatches(medias, 100).Error
	})

	for _, media := range medias {
		if media.MediaType == models.MediaTypeVideo {
			s.extractCoverInBackground(media)
		}
	}

	return medias, errors.WithStack(err)
}

func (s *MediaService) extractCoverInBackground(media *models.Media) error {
	src := media.Path
	dest := filepath.Join(s.coverRoot, fmt.Sprintf("%d.jpg", media.ID))

	logging.GetLogger().Info("extract cover in background", zap.String("src", media.Path), zap.String("dest", dest))
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
	go s.runExtractCoverTask(src, dest, media, task)

	return nil

}

// runExtractCoverTask 提取封面任务
func (s *MediaService) runExtractCoverTask(src, dest string, media *models.Media, task *models.Task) {
	err := s.ff.ExtractCoverInBackground(src, dest)
	if err != nil {
		logging.GetLogger().Error("extract cover failed", zap.Error(err))
		task.Status = models.TaskStatusFailed
		task.StatusString = "error: " + err.Error()
	}

	mc := &models.MediaCover{
		MediaID: media.ID,
		Path:    dest,
	}
	err = s.db.Create(mc).Error
	if err != nil {
		logging.GetLogger().Error("save media cover failed", zap.Error(err))
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

// GetCover 获取封面信息
func (s *MediaService) GetCover(id uint) (*models.MediaCover, error) {
	cover := &models.MediaCover{}
	err := s.db.Model(&models.MediaCover{}).Where("id=?", id).Take(cover).Error
	return cover, errors.WithStack(err)
}

// GetMedia 获取媒体信息
func (s *MediaService) GetMedia(id uint) (*models.Media, error) {
	media := &models.Media{}
	err := s.db.Model(&models.Media{}).Where("id=?", id).Take(media).Error
	return media, errors.WithStack(err)
}
