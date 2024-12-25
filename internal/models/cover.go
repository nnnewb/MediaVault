package models

import "gorm.io/gorm"

type MediaCover struct {
	gorm.Model
	MediaID uint   `gorm:"media_id;unique;"`
	Path    string `gorm:"path;varchar(255);"`
}

func (*MediaCover) TableName() string { return "media_cover" }
