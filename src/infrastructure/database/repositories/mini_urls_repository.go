package repositories

import "context"

type MiniURLsRepository interface {
	GetIfExists(ctx context.Context, originalURL string) (string, error)
	CreateNewMiniURL(ctx context.Context, originalURL string) (string, error)
}

type MiniURLsCacheRepository interface {
	CacheNewMiniURL(ctx context.Context, originalURL, miniURL string)
}
