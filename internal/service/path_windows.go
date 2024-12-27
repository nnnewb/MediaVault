//go:build windows

package service

import (
	"github.com/pkg/errors"
	"golang.org/x/sys/windows"

	"github.com/nnnewb/media-vault/internal/models"
)

func (s *PathService) ListRoot() ([]models.PathEntry, error) {
	ret := make([]models.PathEntry, 0, 32)
	bitmask, err := windows.GetLogicalDrives()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	for i := 0; i < 32; i++ {
		if bitmask&(1<<i) > 0 {
			driveRoot := string(rune(int('A')+i)) + ":/"
			ret = append(ret, models.PathEntry{
				Name:  driveRoot,
				IsDir: true,
				Path:  driveRoot,
			})

		}
	}
	return ret, nil
}
