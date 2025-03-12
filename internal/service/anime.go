package service

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/nnnewb/media-vault/internal/models"
)

type AnimeService struct {
	db *gorm.DB
}

func NewAnimeService(db *gorm.DB) *AnimeService {
	return &AnimeService{db: db}
}

func (s *AnimeService) List(pagination Pagination, ordering OrderBy) ([]*models.Anime, int64, error) {
	var ret []*models.Anime
	var count int64
	err := s.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&models.Anime{}).Scopes(pagination.WithDefault().Scope(), ordering.Scope()).Preload("AnimeEpisodes").Preload("Tags").Find(&ret).Error
		if err != nil {
			return errors.WithStack(err)
		}
		err = tx.Model(&models.Anime{}).Count(&count).Error
		if err != nil {
			return errors.WithStack(err)
		}
		return nil
	})
	return ret, count, errors.Wrap(err, "transaction failed")
}
