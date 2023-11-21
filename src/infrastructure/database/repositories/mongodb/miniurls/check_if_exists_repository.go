package miniurls

import (
	"context"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/DiegoSepuSoto/mini-url-builder-api/src/infrastructure/database/repositories/mongodb/miniurls/entities"
	"github.com/DiegoSepuSoto/mini-url-builder-api/src/shared"
)

func (r *miniURLsRepository) GetIfExists(ctx context.Context, originalURL string) (string, error) {
	filter := bson.D{{Key: "original_url", Value: originalURL}}

	var miniURLRecord entities.MiniURLRecord
	err := r.mongoDBCollection.FindOne(ctx, filter).Decode(&miniURLRecord)
	if err != nil {
		if err.Error() == mongo.ErrNoDocuments.Error() {
			return "", shared.BuildError(
				http.StatusNotFound,
				shared.DatabaseNotFoundError,
				err.Error(),
				"miniURLsRepository")
		}

		return "", shared.BuildError(http.StatusInternalServerError, shared.DatabaseFindError, err.Error(), "miniURLsRepository")
	}

	return fmt.Sprintf("/%s", miniURLRecord.NewURL), nil
}
