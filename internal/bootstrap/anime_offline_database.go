package bootstrap

import (
	"encoding/json"
	"os"
	"strings"

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
	} `json:"anime_season"`
	Picture   string `json:"picture"`
	Thumbnail string `json:"thumbnail"`
	Duration  struct {
		Value int32  `json:"value"`
		Unit  string `json:"unit"`
	} `json:"duration"`
	Synonyms     []string `json:"synonyms"`
	RelatedAnime []string `json:"related_anime"`
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

	for _, a := range data.Data {
		var season models.AnimeSeason
		switch strings.ToUpper(a.AnimeSeason.Season) {
		case "SPRING":
			season = models.AnimeSeasonSpring
		case "SUMMER":
			season = models.AnimeSeasonSummer
		case "FALL":
			season = models.AnimeSeasonFall
		case "WINTER":
			season = models.AnimeSeasonWinter
		default:
			season = models.AnimeSeasonUndefined
		}

		var status models.AnimeStatus
		switch strings.ToUpper(a.Status) {
		case "FINISHED":
			status = models.AnimeStatusFinished
		case "ONGOING":
			status = models.AnimeStatusOngoing
		case "UPCOMING":
			status = models.AnimeStatusUpcoming
		default:
			status = models.AnimeStatusUnknown
		}

		record := models.Anime{
			Title:         a.Title,
			Synonyms:      a.Synonyms,
			TotalEpisodes: a.Episodes,
			ReleaseYear:   a.AnimeSeason.Year,
			Season:        season,
			Status:        status,
			Tags:          []models.AnimeTag{},
		}
		for _, t := range a.Tags {
			record.Tags = append(record.Tags, models.AnimeTag{Tag: t})
		}

		logging.GetLogger().Info("import anime", zap.String("title", record.Title))
		err = db.Save(&record).Error
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
