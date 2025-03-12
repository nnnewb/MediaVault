package service

import (
	"gitee.com/uniqptr/media-vault.git/internal/models"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type MediaService struct {
	db *gorm.DB
}

func NewMediaService(db *gorm.DB) *MediaService {
	return &MediaService{db: db}
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
