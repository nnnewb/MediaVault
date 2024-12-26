//go:build !windows

package service

import (
	"gitee.com/uniqptr/media-vault.git/internal/models"
)

func (s *PathService) ListDrives() ([]models.PathEntry, error) {
	return []models.PathEntry{
		{
			Name:  "/",
			IsDir: true,
			Path:  "/",
		},
	}, nil
}
