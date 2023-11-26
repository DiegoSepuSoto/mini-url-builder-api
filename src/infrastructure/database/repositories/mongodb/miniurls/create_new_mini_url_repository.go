package miniurls

import (
	"context"
	"math/big"
	"net/http"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/DiegoSepuSoto/mini-url-builder-api/src/helpers"
	"github.com/DiegoSepuSoto/mini-url-builder-api/src/infrastructure/database/repositories/mongodb/miniurls/entities"
	"github.com/DiegoSepuSoto/mini-url-builder-api/src/shared"
)

const base62 = 62

func (r *miniURLsRepository) CreateNewMiniURL(ctx context.Context, originalURL string) (string, error) {
	newMiniURL := entities.MiniURLRecord{
		OriginalURL: originalURL,
		NewURL:      createNewURL(),
		CreatedAt:   time.Now(),
	}

	_, err := r.mongoDBCollection.InsertOne(ctx, newMiniURL)
	if err != nil {
		return "", shared.BuildError(
			http.StatusInternalServerError,
			shared.DatabaseInsertError,
			err.Error(),
			"miniURLsRepository")
	}

	log.Println("record created!")

	return newMiniURL.NewURL, nil
}

func createNewURL() string {
	newURLExpectedLen := 6

	counterValue := viper.GetInt64("COUNTER_VALUE")
	counterMaxValue := viper.GetInt64("COUNTER_MAX")

	if !helpers.IsRangeConfigValid(counterValue, counterMaxValue) {
		err := helpers.LoadIDRanges()
		if err != nil {
			panic(err)
		}
	}

	newURL := big.NewInt(counterValue).Text(base62)
	newURLLen := len(newURL)

	if newURLLen < newURLExpectedLen {
		paddingLength := newURLExpectedLen - newURLLen

		newURL = strings.Repeat("0", paddingLength) + newURL
	}

	viper.Set("COUNTER_VALUE", counterValue+1)

	return newURL
}
