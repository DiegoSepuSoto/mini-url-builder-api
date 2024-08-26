package shortener

import "github.com/DiegoSepuSoto/mini-url-builder-api/src/infrastructure/database/repositories"

type shortenerUseCase struct {
	miniURLsRepository      repositories.MiniURLsRepository
	miniURLsCacheRepository repositories.MiniURLsCacheRepository
}

func NewShortenerUseCase(miniURLsRepository repositories.MiniURLsRepository,
	miniURLsCacheRepository repositories.MiniURLsCacheRepository) *shortenerUseCase {
	return &shortenerUseCase{
		miniURLsRepository:      miniURLsRepository,
		miniURLsCacheRepository: miniURLsCacheRepository,
	}
}
