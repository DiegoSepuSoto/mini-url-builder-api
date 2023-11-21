package shortener

import (
	"context"
	"github.com/DiegoSepuSoto/mini-url-builder-api/src/domain/models"
	"github.com/stretchr/testify/mock"
)

type shortenerUseCaseMock struct {
	mock.Mock
}

func (m *shortenerUseCaseMock) BuildMiniURL(ctx context.Context, originalURL string) (*models.MiniURLResponse, error) {
	args := m.Called(ctx, originalURL)

	if args.Get(1) == nil {
		return args.Get(0).(*models.MiniURLResponse), nil
	}

	return nil, args.Error(1)
}
