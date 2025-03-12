package models

import (
	"time"
)

type PathEntry struct {
	Name      string    `json:"name"`
	IsDir     bool      `json:"is_dir"`
	Path      string    `json:"path"`
	UpdatedAt time.Time `json:"updated_at"`
}
