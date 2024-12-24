package service

import (
	"mime"
	"path/filepath"
	"strings"

	"gitee.com/uniqptr/media-vault.git/internal/models"
)

type MediaInfer struct {
}

func NewMediaInfer() *MediaInfer {
	return &MediaInfer{}
}

func (m *MediaInfer) InferMediaType(path string) models.MediaType {
	cleaned := filepath.Clean(path)

	suffix := filepath.Ext(cleaned)
	mimeType := mime.TypeByExtension(suffix)
	switch {
	case strings.HasPrefix(mimeType, "video/"):
		return models.MediaTypeVideo
	case strings.EqualFold(suffix, ".srt") || strings.EqualFold(suffix, ".ass"):
		return models.MediaTypeSubtitle
	default:
		return models.MediaTypeUnknown
	}
}
