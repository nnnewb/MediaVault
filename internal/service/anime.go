package service

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

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

// Search 在离线动画数据库里搜索动画数据
// TODO: 搜索应该包含 synonyms 的搜索结果，考虑接 sonic 做成全文搜索，或者直接 sqlite 全文搜索
func (s *AnimeService) Search(term string, pagination Pagination, by OrderBy) ([]*models.AnimeOfflineDatabase, int64, error) {
	var ret []*models.AnimeOfflineDatabase
	var count int64
	err := s.db.Transaction(func(tx *gorm.DB) error {
		err := tx.
			Model(&models.AnimeOfflineDatabase{}).
			Where(clause.Like{Column: "title", Value: "%" + term + "%"}).
			Scopes(pagination.WithDefault().Scope(), by.Scope()).
			Preload("Synonyms").
			Preload("Tags").
			Find(&ret).
			Error
		if err != nil {
			return errors.WithStack(err)
		}

		err = tx.
			Model(&models.AnimeOfflineDatabase{}).
			Where(clause.Like{Column: "title", Value: "%" + term + "%"}).
			Count(&count).
			Error
		if err != nil {
			return errors.WithStack(err)
		}
		return nil
	})
	return ret, count, errors.Wrap(err, "transaction failed")
}
