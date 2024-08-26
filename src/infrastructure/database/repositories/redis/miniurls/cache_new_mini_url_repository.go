package miniurls

import (
	"context"

	log "github.com/sirupsen/logrus"
)

func (r *miniURLsCacheRepository) CacheNewMiniURL(ctx context.Context, originalURL, miniURL string) {
	_, err := r.redisClient.Set(ctx, miniURL, originalURL).Result()
	if err != nil {
		log.Printf("%s key could not be created in cache repository", miniURL)
	}
}
