package miniurls

import "github.com/DiegoSepuSoto/mini-url-builder-api/src/shared"

type miniURLsRepository struct {
	mongoDBCollection shared.MongoDBCollection
}

func NewMiniURLsRepository(mongoDBCollection shared.MongoDBCollection) *miniURLsRepository {
	return &miniURLsRepository{
		mongoDBCollection: mongoDBCollection,
	}
}
