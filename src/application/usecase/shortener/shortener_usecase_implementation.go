package shortener

import "github.com/DiegoSepuSoto/mini-url-builder-api/src/infrastructure/database/repositories"

type shortenerUseCase struct {
	miniURLsRepository repositories.MiniURLsRepository
}

func NewShortenerUseCase(miniURLsRepository repositories.MiniURLsRepository) *shortenerUseCase {
	return &shortenerUseCase{
		miniURLsRepository: miniURLsRepository,
	}
}
