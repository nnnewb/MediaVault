package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type AnimeInformation struct {
	gorm.Model
	LocalizationID int64                       `gorm:"localization_id"` // 本地化信息外键 ID
	Title          string                      `gorm:"title"`           // 番剧标题
	AlsoKnownAs    datatypes.JSONSlice[string] `gorm:"also_known_as"`   // 存在多个标题时，列出别名
	ReleaseDate    time.Time                   `gorm:"release_date"`    // 发布日期
	Episodes       int32                       `gorm:"episodes"`        // 集数
}

func (*AnimeInformation) TableName() string { return "anime_information" }
