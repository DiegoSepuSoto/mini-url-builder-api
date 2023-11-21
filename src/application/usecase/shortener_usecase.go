package usecase

import (
	"context"

	"github.com/DiegoSepuSoto/mini-url-builder-api/src/domain/models"
)

type ShortenerUseCase interface {
	BuildMiniURL(ctx context.Context, originalURL string) (*models.MiniURLResponse, error)
}
