package shortener

import (
	"context"
	"errors"
	"os"

	"github.com/DiegoSepuSoto/mini-url-builder-api/src/domain/models"
	"github.com/DiegoSepuSoto/mini-url-builder-api/src/shared"
)

func (u *shortenerUseCase) BuildMiniURL(ctx context.Context, originalURL string) (*models.MiniURLResponse, error) {
	miniURL, err := u.miniURLsRepository.GetIfExists(ctx, originalURL)
	if err != nil {
		var applicationError *shared.ApplicationError
		ok := errors.As(err, &applicationError)

		if ok && applicationError.ErrorCode == shared.DatabaseNotFoundError {
			return u.createMiniURL(ctx, originalURL)
		}

		return nil, err
	}

	return &models.MiniURLResponse{
		Host:    os.Getenv("MINI_URLs_HOST"),
		MiniURL: miniURL,
	}, nil
}

func (u *shortenerUseCase) createMiniURL(ctx context.Context, originalURL string) (*models.MiniURLResponse, error) {
	miniURL, err := u.miniURLsRepository.CreateNewMiniURL(ctx, originalURL)
	if err != nil {
		return nil, err
	}

	u.miniURLsCacheRepository.CacheNewMiniURL(ctx, originalURL, miniURL)

	return &models.MiniURLResponse{
		Host:    os.Getenv("MINI_URLs_HOST"),
		MiniURL: miniURL,
	}, nil
}
