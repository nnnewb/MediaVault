package service

import (
	"fmt"
	"path/filepath"

	"gitee.com/uniqptr/media-vault.git/internal/constants"
	"gitee.com/uniqptr/media-vault.git/internal/models"

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
		for _, option := range options {
			option(tx)
		}

		return tx.Find(&medias).Error
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
			s.ff.ExtractCoverInBackground(media.Path, filepath.Join(s.coverRoot, fmt.Sprintf("%d.jpg", media.ID)))
		}
	}

	return medias, errors.WithStack(err)
}
