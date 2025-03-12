package bootstrap

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/nnnewb/media-vault/internal/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type anime struct {
	Sources     []string `json:"sources"`
	Title       string   `json:"title"`
	Type        string   `json:"type"`
	Episodes    int32    `json:"episodes"`
	Status      string   `json:"status"`
	AnimeSeason struct {
		Year   int32  `json:"year"`
		Season string `json:"season"`
	} `json:"animeSeason"`
	Picture   string `json:"picture"`
	Thumbnail string `json:"thumbnail"`
	Duration  struct {
		Value int32  `json:"value"`
		Unit  string `json:"unit"`
	} `json:"duration"`
	Synonyms     []string `json:"synonyms"`
	RelatedAnime []string `json:"relatedAnime"`
	Tags         []string `json:"tags"`
}

func flatAnimeOfflineDatabase(entries []anime) ([]models.AnimeOfflineDatabase, []models.AnimeOfflineDatabaseTag, []models.AnimeOfflineDatabaseSynonym) {
	var (
		records  = make([]models.AnimeOfflineDatabase, 0, len(entries))
		tags     = make([]models.AnimeOfflineDatabaseTag, 0, len(entries)*8)
		synonyms = make([]models.AnimeOfflineDatabaseSynonym, 0, len(entries)*8)
	)
	for i, a := range entries {
		records = append(records, models.AnimeOfflineDatabase{
			ID:        uint(i + 1),
			Title:     a.Title,
			Sources:   a.Sources,
			Type:      a.Type,
			Episodes:  a.Episodes,
			Status:    a.Status,
			Year:      a.AnimeSeason.Year,
			Season:    a.AnimeSeason.Season,
			Picture:   a.Picture,
			Thumbnail: a.Thumbnail,
			Duration:  time.Duration(a.Duration.Value) * time.Second,
		})

		for _, t := range a.Tags {
			tags = append(tags, models.AnimeOfflineDatabaseTag{
				AnimeID: uint(i + 1),
				Tag:     t,
			})
		}

		for _, s := range a.Synonyms {
			synonyms = append(synonyms, models.AnimeOfflineDatabaseSynonym{
				AnimeID: uint(i + 1),
				Synonym: s,
			})
		}
	}
	return records, tags, synonyms
}

func BootstrapAnimeOfflineDatabase(db *gorm.DB, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return errors.WithStack(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var data struct {
		Data []anime `json:"data"`
	}
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		return errors.WithStack(err)
	}

	if err = db.Exec(fmt.Sprintf("DELETE FROM %s", (*models.AnimeOfflineDatabaseTag)(nil).TableName())).Error; err != nil {
		return errors.WithStack(err)
	}
	if err = db.Exec(fmt.Sprintf("DELETE FROM %s", (*models.AnimeOfflineDatabaseSynonym)(nil).TableName())).Error; err != nil {
		return errors.WithStack(err)
	}
	if err = db.Exec(fmt.Sprintf("DELETE FROM %s", (*models.AnimeOfflineDatabase)(nil).TableName())).Error; err != nil {
		return errors.WithStack(err)
	}

	records, tags, synonyms := flatAnimeOfflineDatabase(data.Data)
	if err = db.CreateInBatches(records, 500).Error; err != nil {
		return errors.WithStack(err)
	}
	if err = db.CreateInBatches(tags, 500).Error; err != nil {
		return errors.WithStack(err)
	}
	if err = db.CreateInBatches(synonyms, 500).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
