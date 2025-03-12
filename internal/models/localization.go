package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type AnimeLocalization struct {
	gorm.Model
	Title       string                      `gorm:"title"`         // 本地化标题
	Local       string                      `gorm:"local"`         // 地区代码
	AlsoKnownAs datatypes.JSONSlice[string] `gorm:"also_known_as"` // 存在多个本地化标题时，列出别名
}

func (*AnimeLocalization) TableName() string { return "anime_localization" }
