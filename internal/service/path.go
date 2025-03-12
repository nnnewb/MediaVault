package service

import (
	"os"
	"path/filepath"

	"gitee.com/uniqptr/media-vault.git/internal/models"
	"github.com/pkg/errors"
)

type PathService struct {
}

func NewPathService() *PathService {
	return &PathService{}
}

func (s *PathService) ReadDir(path string) ([]models.PathEntry, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	ret := make([]models.PathEntry, 0, len(entries))
	for _, entry := range entries {
		ret = append(ret, models.PathEntry{
			Name:  entry.Name(),
			IsDir: entry.IsDir(),
			Path:  filepath.Join(path, entry.Name()),
		})
	}
	return ret, nil
}
