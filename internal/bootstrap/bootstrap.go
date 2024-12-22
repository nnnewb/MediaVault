package bootstrap

import (
	"gitee.com/uniqptr/media-vault.git/internal/models"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func BootstrapDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Media{},
		&models.MediaRelation{},
		&models.AnimeInformation{},
		&models.AnimeLocalization{},
		&models.Subtitle{},
	)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
