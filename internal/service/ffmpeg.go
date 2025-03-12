package service

import (
	"fmt"
	"os/exec"

	"github.com/nnnewb/media-vault/internal/logging"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	CoverWidth  = 400
	CoverHeight = 300
)

type FFMPEGService struct {
	db   *gorm.DB
	path string
}

func NewFFMPEGService(db *gorm.DB, path string) *FFMPEGService {
	return &FFMPEGService{
		db:   db,
		path: path,
	}
}

// ExtractCoverInBackground 提取封面
func (s *FFMPEGService) ExtractCoverInBackground(src, dest string) error {
	vf := fmt.Sprintf("scale=%d:%d:force_original_aspect_ratio=1,pad=%d:%d:-1:-1:color=black", CoverWidth, CoverHeight, CoverWidth, CoverHeight)
	cmd := exec.Command(s.path, "-i", src, "-vframes", "1", "-vf", vf, "-y", dest)

	logging.GetLogger().Info("execute ffmpeg command", zap.String("command", cmd.String()))
	err := cmd.Run()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
