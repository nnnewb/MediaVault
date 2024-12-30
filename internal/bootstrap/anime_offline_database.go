package bootstrap

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/nnnewb/media-vault/internal/logging"
	"github.com/nnnewb/media-vault/internal/models"
	"github.com/pkg/errors"
	"go.uber.org/zap"
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

	// truncate tables
	err = db.Exec(fmt.Sprintf("DELETE FROM %s", (&models.AnimeOfflineDatabase{}).TableName())).Error
	if err != nil {
		return errors.WithStack(err)
	}

	err = db.Exec(fmt.Sprintf("DELETE FROM %s", (&models.AnimeOfflineDatabaseSynonym{}).TableName())).Error
	if err != nil {
		return errors.WithStack(err)
	}

	err = db.Exec(fmt.Sprintf("DELETE FROM %s", (&models.AnimeOfflineDatabaseTag{}).TableName())).Error
	if err != nil {
		return errors.WithStack(err)
	}

	// import json data
	for _, a := range data.Data {
		synonyms := make([]models.AnimeOfflineDatabaseSynonym, 0, len(a.Synonyms))
		for _, s := range a.Synonyms {
			synonyms = append(synonyms, models.AnimeOfflineDatabaseSynonym{Synonym: s})
		}

		tags := make([]models.AnimeOfflineDatabaseTag, 0, len(a.Tags))
		for _, t := range a.Tags {
			tags = append(tags, models.AnimeOfflineDatabaseTag{Tag: t})
		}

		record := models.AnimeOfflineDatabase{
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
			Synonyms:  synonyms,
			Tags:      tags,
		}

		logging.GetLogger().Info("import anime", zap.String("title", record.Title))
		err = db.Save(&record).Error
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
