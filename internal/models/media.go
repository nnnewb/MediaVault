package models

import (
	"path/filepath"
	"time"

	"gorm.io/gorm"
)

type MediaType int32

const (
	MediaTypeUnknown  MediaType = iota // 默认，未知
	MediaTypeVideo                     // 视频，.mkv/.mp4
	MediaTypeSubtitle                  // 字幕, .ass/.srt 等等
)

type Media struct {
	gorm.Model
	Path          string    `gorm:"path;type:varchar(255);"` // 文件路径
	MediaType     MediaType `gorm:"title;type:integer;"`     // 媒体类型
	InformationID int64     `gorm:"information_id"`          // 信息外键 ID

	// preload fields
	MediaCover *MediaCover
}

func (*Media) TableName() string { return "media" }

func (m *Media) ToDTO() MediaDTO {
	var d MediaDTO
	m.AsDTO(&d)
	return d
}

func (m *Media) AsDTO(d *MediaDTO) {
	d.ID = m.ID
	d.CreatedAt = m.CreatedAt
	d.UpdatedAt = m.UpdatedAt
	d.DeletedAt = m.DeletedAt
	d.Path = m.Path
	d.Name = filepath.Base(m.Path)
	d.MediaType = m.MediaType
	d.InformationID = m.InformationID
	if m.MediaCover != nil {
		d.CoverID = m.MediaCover.ID
	}
}

type MediaDTO struct {
	ID            uint           `json:"id"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
	Path          string         `json:"path"`
	Name          string         `json:"name"`
	MediaType     MediaType      `json:"media_type"`
	InformationID int64          `json:"information_id"`
	CoverID       uint           `json:"cover_id"`
}

type MediaRelationType int32

const (
	InvalidMediaRelation MediaRelationType = iota //
	Series                                        // 同系列动画
)

type MediaRelation struct {
	MediaID   int64             `gorm:"media_id"`               // 番剧 ID
	RelatedID int64             `gorm:"related_id"`             // 关联番剧 ID
	Relation  MediaRelationType `gorm:"relation;type:integer;"` // 关系
}
