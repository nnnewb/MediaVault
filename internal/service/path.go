package service

import (
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"gitee.com/uniqptr/media-vault.git/internal/logging"
	"gitee.com/uniqptr/media-vault.git/internal/models"
)

type PathService struct {
}

func NewPathService() *PathService {
	return &PathService{}
}

func (s *PathService) ReadDir(path string) ([]models.PathEntry, error) {
	if path == "" {
		return s.ListRoot()
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	ret := make([]models.PathEntry, 0, len(entries))
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			logging.GetLogger().Warn("can not get info of entry", zap.Error(err))
		}

		var updatedAt time.Time
		if info != nil {
			updatedAt = info.ModTime()
		}

		ret = append(ret, models.PathEntry{
			Name:      entry.Name(),
			IsDir:     entry.IsDir(),
			Path:      filepath.Join(path, entry.Name()),
			UpdatedAt: updatedAt,
		})
	}
	return ret, nil
}
