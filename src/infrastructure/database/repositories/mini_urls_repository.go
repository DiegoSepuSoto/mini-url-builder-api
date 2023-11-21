package repositories

import "context"

type MiniURLsRepository interface {
	GetIfExists(ctx context.Context, originalURL string) (string, error)
	CreateNewMiniURL(ctx context.Context, originalURL string) (string, error)
}
