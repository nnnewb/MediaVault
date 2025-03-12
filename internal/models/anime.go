package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type AnimeSeason string

const (
	AnimeSeasonUndefined AnimeSeason = ""       // 未知
	AnimeSeasonSpring    AnimeSeason = "SPRING" // 春季
	AnimeSeasonSummer    AnimeSeason = "SUMMER" // 夏季
	AnimeSeasonFall      AnimeSeason = "FALL"   // 秋季
	AnimeSeasonWinter    AnimeSeason = "WINTER" // 冬季
)

type AnimeStatus string

const (
	AnimeStatusUnknown  AnimeStatus = ""
	AnimeStatusFinished AnimeStatus = "FINISHED"
	AnimeStatusOngoing  AnimeStatus = "ONGOING"
	AnimeStatusUpcoming AnimeStatus = "UPCOMING"
)

type Anime struct {
	gorm.Model
	Title         string                      `gorm:"title"`          // 番剧标题
	Synonyms      datatypes.JSONSlice[string] `gorm:"synonyms"`       // 存在多个本地化标题时，列出别名
	TotalEpisodes int32                       `gorm:"total_episodes"` // 集数
	ReleaseYear   int32                       `gorm:"release_year"`   // 发布年份
	Season        AnimeSeason                 `gorm:"season"`         // 发布季度
	Status        AnimeStatus                 `gorm:"status"`         // 播送进度
	Tags          []AnimeTag                  ``                      // 标签
	AnimeEpisodes []AnimeEpisode              ``                      // 剧集
}

func (*Anime) TableName() string { return "anime" }

type AnimeTag struct {
	ID      uint   `gorm:"primaryKey"` //
	AnimeID uint   `gorm:"anime_id"`   // 番剧ID
	Tag     string `gorm:"tag"`        // 标签
}

func (*AnimeTag) TableName() string { return "anime_tag" }

type ThumbnailSize string

const (
	ThumbnailSizeLarge ThumbnailSize = "large"
	ThumbnailSizeSmall ThumbnailSize = "small"
)

type AnimeThumbnail struct {
	ID            uint          `gorm:"primaryKey"`     //
	AnimeID       uint          `gorm:"anime_id"`       // 番剧ID
	ThumbnailSize ThumbnailSize `gorm:"thumbnail_size"` // 缩略图尺寸
	ThumbnailPath string        `gorm:"thumbnail_path"` // 缩略图路径
}

func (*AnimeThumbnail) TableName() string { return "anime_thumbnail" }

type AnimeEpisodeType string

const (
	AnimeEpisodeTypeUnknown AnimeEpisodeType = ""   // 未知
	AnimeEpisodeTypeTV      AnimeEpisodeType = "TV" // TV
	AnimeEpisodeTypeMovie   AnimeEpisodeType = "MOVIE"
	AnimeEpisodeTypeOVA     AnimeEpisodeType = "OVA"
	AnimeEpisodeTypeONA     AnimeEpisodeType = "ONA"
	AnimeEpisodeTypeSpecial AnimeEpisodeType = "SPECIAL"
)

type AnimeEpisode struct {
	ID      uint             `gorm:"primaryKey"` //
	AnimeID uint             `gorm:"anime_id"`   // 番剧ID
	MediaID uint             `gorm:"media_id"`   // 媒体ID
	Type    AnimeEpisodeType `gorm:"type"`       // 剧集类型
	Index   int              `gorm:"index"`      // 剧集索引
	Media   *Media           ``                  // 媒体
}

func (*AnimeEpisode) TableName() string { return "anime_episode" }
