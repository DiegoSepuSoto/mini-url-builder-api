package miniurls

import "github.com/DiegoSepuSoto/mini-url-builder-api/src/shared"

type miniURLsCacheRepository struct {
	redisClient shared.RedisClient
}

func NewMiniURLsCacheRepository(redisClient shared.RedisClient) *miniURLsCacheRepository {
	return &miniURLsCacheRepository{
		redisClient: redisClient,
	}
}
