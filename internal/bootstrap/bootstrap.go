package bootstrap

import (
	"os"
	"path/filepath"

	"github.com/nnnewb/media-vault/internal/constants"
	"github.com/nnnewb/media-vault/internal/models"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func BootstrapDataFolder(dataRoot string) error {
	if err := os.MkdirAll(dataRoot, 0o755); err != nil {
		return errors.WithStack(err)
	}

	if err := os.MkdirAll(filepath.Join(dataRoot, constants.DataFolderCovers), 0o755); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func BootstrapDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Media{},
		&models.MediaCover{},
		&models.MediaRelation{},
		&models.AnimeInformation{},
		&models.AnimeLocalization{},
		&models.Subtitle{},
		&models.Task{},
		&models.TaskTag{},
	)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
