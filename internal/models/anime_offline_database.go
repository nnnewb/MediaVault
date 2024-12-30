package models

import (
	"time"

	"gorm.io/datatypes"
)

type AnimeOfflineDatabase struct {
	ID        uint                          `gorm:"primaryKey"`
	CreatedAt time.Time                     `gorm:"created_at"`
	UpdatedAt time.Time                     `gorm:"updated_at"`
	Title     string                        `gorm:"title"`
	Sources   datatypes.JSONSlice[string]   `gorm:"sources"`
	Type      string                        `gorm:"type"`
	Episodes  int32                         `gorm:"episodes"`
	Status    string                        `gorm:"status"`
	Year      int32                         `gorm:"year"`
	Season    string                        `gorm:"season"`
	Picture   string                        `gorm:"picture"`
	Thumbnail string                        `gorm:"thumbnail"`
	Duration  time.Duration                 `gorm:"duration"`
	Synonyms  []AnimeOfflineDatabaseSynonym `gorm:"foreignKey:AnimeID"`
	Tags      []AnimeOfflineDatabaseTag     `gorm:"foreignKey:AnimeID"`
}

func (*AnimeOfflineDatabase) TableName() string {
	return "anime_offline_database"
}

func (a *AnimeOfflineDatabase) ToDTO() *AnimeOfflineDatabaseDTO {
	var dto AnimeOfflineDatabaseDTO
	a.AsDTO(&dto)
	return &dto
}

func (a *AnimeOfflineDatabase) AsDTO(dto *AnimeOfflineDatabaseDTO) {
	dto.ID = a.ID
	dto.Title = a.Title
	dto.Sources = a.Sources
	dto.Type = a.Type
	dto.Episodes = a.Episodes
	dto.Status = a.Status
	dto.Year = a.Year
	dto.Season = a.Season
	dto.Picture = a.Picture
	dto.Thumbnail = a.Thumbnail
	dto.Duration = a.Duration
	dto.Synonyms = make([]string, 0, len(a.Synonyms))
	for _, synonym := range a.Synonyms {
		dto.Synonyms = append(dto.Synonyms, synonym.Synonym)
	}
	dto.Tags = make([]string, 0, len(a.Tags))
	for _, tag := range a.Tags {
		dto.Tags = append(dto.Tags, tag.Tag)
	}
}

type AnimeOfflineDatabaseSynonym struct {
	ID      uint   `gorm:"id"`
	AnimeID uint   `gorm:"anime_id"`
	Synonym string `gorm:"synonym"`
}

func (*AnimeOfflineDatabaseSynonym) TableName() string { return "anime_offline_database_synonym" }

type AnimeOfflineDatabaseTag struct {
	ID      uint   `gorm:"id"`
	AnimeID uint   `gorm:"anime_id"`
	Tag     string `gorm:"tag"`
}

func (*AnimeOfflineDatabaseTag) TableName() string { return "anime_offline_database_tag" }

type AnimeOfflineDatabaseDTO struct {
	ID        uint          `json:"id"`
	Title     string        `json:"title"`
	Sources   []string      `json:"sources"`
	Type      string        `json:"type"`
	Episodes  int32         `json:"episodes"`
	Status    string        `json:"status"`
	Year      int32         `json:"year"`
	Season    string        `json:"season"`
	Picture   string        `json:"picture"`
	Thumbnail string        `json:"thumbnail"`
	Duration  time.Duration `json:"duration"`
	Synonyms  []string      `json:"synonyms"`
	Tags      []string      `json:"tags"`
}
