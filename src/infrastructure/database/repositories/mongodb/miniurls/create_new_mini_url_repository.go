package miniurls

import (
	"context"
	"github.com/DiegoSepuSoto/mini-url-builder-api/src/helpers"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/DiegoSepuSoto/mini-url-builder-api/src/infrastructure/database/repositories/mongodb/miniurls/entities"
	"github.com/DiegoSepuSoto/mini-url-builder-api/src/shared"
)

func (r *miniURLsRepository) CreateNewMiniURL(ctx context.Context, originalURL string) (string, error) {
	newMiniURL := entities.MiniURLRecord{
		OriginalURL: originalURL,
		NewURL:      helpers.CreateNewMiniURL(),
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
