package models

import (
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
}

func (*Media) TableName() string { return "media" }

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
