package service

import (
	"gitee.com/uniqptr/media-vault.git/internal/models"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type MediaService struct {
	db    *gorm.DB
	infer *MediaInfer
}

func NewMediaService(db *gorm.DB, infer *MediaInfer) *MediaService {
	return &MediaService{
		db: db,
		infer: infer,
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
			Path: path,
			MediaType: s.infer.InferMediaType(path),
		})
	}
	err := s.db.Transaction(func(tx *gorm.DB) error {
		return tx.CreateInBatches(medias, 100).Error
	})
	return medias, errors.WithStack(err)
}
