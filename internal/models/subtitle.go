package models

import (
	"gorm.io/gorm"
)

type Subtitle struct {
	gorm.Model
	AnimeID int64  `gorm:"anime_id"` // 动画外键 ID
	Author  string `gorm:"author"`   // 字幕组
}

func (*Subtitle) TableName() string { return "subtitle" }
