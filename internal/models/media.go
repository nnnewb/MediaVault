package models

import (
	"gorm.io/gorm"
)

type MediaType int32

const (
	Invalid       MediaType = iota // 默认
	AnimeTVSeries                  // TV 动画
	AnimeOVA                       // OVA 动画
	AnimeOAD                       // OAD 动画
	AnimeSP                        // SP 动画
	AnimeMovie                     // 动画电影
)

type Media struct {
	gorm.Model
	Path          string    `gorm:"path"`
	MediaType     MediaType `gorm:"title"`          // 媒体类型
	InformationID int64     `gorm:"information_id"` // 信息外键 ID
}

func (*Media) TableName() string { return "media" }

type MediaRelationType int32

const (
	InvalidMediaRelation MediaRelationType = iota //
	Series                                        // 同系列动画
)

type MediaRelation struct {
	MediaID   int64             `gorm:"media_id"`   // 番剧 ID
	RelatedID int64             `gorm:"related_id"` // 关联番剧 ID
	Relation  MediaRelationType `gorm:"relation"`   // 关系
}
