package sync

import (
	"github.com/DiegoSepuSoto/mini-url-builder-api/src/shared"
)

type syncRepository struct {
	HTTPClient shared.HTTPClient
}

func NewSyncRepository(httpClient shared.HTTPClient) *syncRepository {
	return &syncRepository{
		HTTPClient: httpClient,
	}
}
