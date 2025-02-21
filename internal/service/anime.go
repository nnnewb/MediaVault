package service

import (
	"github.com/nnnewb/media-vault/internal/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
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
func (s *AnimeService) Search(term string, pagination Pagination, by OrderBy) ([]*models.AnimeOfflineDatabase, int64, error) {
	var ret []*models.AnimeOfflineDatabase
	var count int64
	err := s.db.Transaction(func(tx *gorm.DB) error {
		err := tx.
			Model(&models.AnimeOfflineDatabase{}).
			Scopes(pagination.WithDefault().Scope(), by.Scope()).
			Where(
				"title like ? OR id IN (?)",
				"%"+term+"%",
				s.db.
					Model(&models.AnimeOfflineDatabaseSynonym{}).
					Distinct("anime_id").
					Where("synonym like ?", "%"+term+"%"),
			).
			Preload("Synonyms").
			Preload("Tags").
			Find(&ret).
			Error
		if err != nil {
			return errors.WithStack(err)
		}

		err = tx.
			Model(&models.AnimeOfflineDatabase{}).
			Where(
				"title like ? OR id IN (?)",
				"%"+term+"%",
				s.db.
					Model(&models.AnimeOfflineDatabaseSynonym{}).
					Distinct("anime_id").
					Where("synonym like ?", "%"+term+"%"),
			).
			Count(&count).
			Error
		if err != nil {
			return errors.WithStack(err)
		}
		return nil
	})
	return ret, count, errors.Wrap(err, "transaction failed")
}

// Info 在离线动画数据库里按id获取动画数据
func (s *AnimeService) Info(id uint) (*models.AnimeOfflineDatabase, error) {
	var ret models.AnimeOfflineDatabase
	err := s.db.Model(&models.AnimeOfflineDatabase{}).Where("id = ?", id).Preload("Synonyms").Preload("Tags").First(&ret).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ret, nil
}
