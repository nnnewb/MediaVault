//go:build !windows

package service

import (
	"github.com/nnnewb/media-vault/internal/models"
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
