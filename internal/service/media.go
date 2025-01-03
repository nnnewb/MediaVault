package service

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/nnnewb/media-vault/internal/constants"
	"github.com/nnnewb/media-vault/internal/logging"
	"github.com/nnnewb/media-vault/internal/models"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type MediaService struct {
	db        *gorm.DB
	infer     *MediaInfer
	ff        *FFMPEGService
	task      *TaskService
	coverRoot string
}

func NewMediaService(db *gorm.DB, dataRoot string, infer *MediaInfer, ff *FFMPEGService, task *TaskService) *MediaService {
	return &MediaService{
		db:        db,
		infer:     infer,
		ff:        ff,
		task:      task,
		coverRoot: filepath.Join(dataRoot, constants.DataFolderCovers),
	}
}

func (s *MediaService) List(q string, pagination Pagination, by OrderBy) ([]*models.Media, int64, error) {
	var medias []*models.Media
	var count int64

	qScope := func(db *gorm.DB) *gorm.DB {
		if q != "" {
			return db.Where("path like ?", "%"+q+"%")
		}
		return db
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		err := tx.
			Model(&models.Media{}).
			Scopes(pagination.WithDefault().Scope(), by.Scope(), qScope).
			Preload("MediaCover").
			Find(&medias).
			Error
		if err != nil {
			return errors.WithStack(err)
		}

		err = tx.Model(&models.Media{}).Scopes(qScope).Count(&count).Error
		if err != nil {
			return errors.WithStack(err)
		}
		return nil
	})
	return medias, count, err
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
			err = s.extractCoverInBackground(media)
			if err != nil {
				logging.GetLogger().Error("extract cover failed", zap.Error(err))
			}
		}
	}

	return medias, errors.WithStack(err)
}

// StartScanMedia 启动后台任务，扫描指定目录下的媒体文件
func (s *MediaService) StartScanMedia(paths ...string) error {
	for _, path := range paths {
		task, err := s.task.CreateTask(fmt.Sprintf("scan %s", path), fmt.Sprintf("scan %s", path), models.TaskCategoryMediaDiscover, []string{}, 1)
		if err != nil {
			return err
		}

		go s.scanInBackground(path, task)
	}

	return nil
}

// runScanInBackground 后台扫描目录下的媒体文件
func (s *MediaService) scanInBackground(path string, task *models.Task) {
	err := filepath.WalkDir(path, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			logging.GetLogger().Error("walk dir failed", zap.String("path", p), zap.Error(err))
			return nil
		}

		if d.IsDir() {
			return nil
		}

		if s.infer.InferMediaType(p) == models.MediaTypeVideo {
			_, err = s.Add(p)
			if err != nil {
				logging.GetLogger().Error("add media failed", zap.String("path", p), zap.Error(err))
				return nil
			}
		}

		return nil
	})
	if err != nil {
		err = s.task.FailTask(task.ID)
		if err != nil {
			logging.GetLogger().Error("mark task fail failed", zap.Error(err))
		}
	}
	err = s.task.FinishTask(task.ID)
	if err != nil {
		logging.GetLogger().Error("mark task finish failed", zap.Error(err))
	}
}

func (s *MediaService) extractCoverInBackground(media *models.Media) error {
	src := media.Path
	dest := filepath.Join(s.coverRoot, fmt.Sprintf("%d.jpg", media.ID))

	logging.GetLogger().Info("extract cover in background", zap.String("src", media.Path), zap.String("dest", dest))
	// create task
	task, err := s.task.CreateTask("extract cover", fmt.Sprintf("extract cover from %s to %s", src, dest), models.TaskCategoryMediaDiscover, []string{}, 1)
	if err != nil {
		return err
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
		err = s.task.FailTask(task.ID)
		if err != nil {
			logging.GetLogger().Error("mark task fail failed", zap.Uint("task_id", task.ID), zap.Error(err))
		}
		return
	}

	mc := &models.MediaCover{
		MediaID: media.ID,
		Path:    dest,
	}
	err = s.db.Create(mc).Error
	if err != nil {
		logging.GetLogger().Error("save media cover failed", zap.Error(err))
		err = s.task.FailTask(task.ID)
		if err != nil {
			logging.GetLogger().Error("mark task fail failed", zap.Uint("task_id", task.ID), zap.Error(err))
		}
		return
	}

	err = s.task.FinishTask(task.ID)
	if err != nil {
		logging.GetLogger().Error("mark task finish failed", zap.Uint("task_id", task.ID), zap.Error(err))
	}
	return
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
