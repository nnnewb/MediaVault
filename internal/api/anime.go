package api

import (
	"gorm.io/gorm"
)

type AnimeControllerV1 struct {
	db *gorm.DB
}

func NewAnimeControllerV1(db *gorm.DB) *AnimeControllerV1 {
	return &AnimeControllerV1{db: db}
}
