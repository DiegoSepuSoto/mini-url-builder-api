package shortener

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/DiegoSepuSoto/mini-url-builder-api/src/shared"
)

func TestBuildShortURLUseCase(t *testing.T) {
	t.Setenv("MINI_URLs_HOST", "https://diego.sepu")

	t.Run("when build short url use case is executed as expected retrieving existing mini url", func(t *testing.T) {
		mockMiniURLRepository := new(miniURLsRepositoryMock)

		mockMiniURLRepository.On("GetIfExists", mock.Anything, mock.AnythingOfType("string")).
			Return("/abc123", nil)

		shortenerUseCase := NewShortenerUseCase(mockMiniURLRepository)

		miniURLResponse, err := shortenerUseCase.BuildMiniURL(context.Background(), "https://www.google.cl")

		mockMiniURLRepository.AssertNotCalled(t, "CreateNewMiniURL")

		assert.Nil(t, err)
		assert.Equal(t, "/abc123", miniURLResponse.MiniURL)
		assert.Equal(t, "https://diego.sepu", miniURLResponse.Host)
	})

	t.Run("when build short url use case is executed as expected creating new mini url", func(t *testing.T) {
		mockMiniURLRepository := new(miniURLsRepositoryMock)

		mockMiniURLRepository.On("GetIfExists", mock.Anything, mock.AnythingOfType("string")).
			Return("", shared.BuildError(http.StatusNotFound, shared.DatabaseNotFoundError, "not found in db", "miniURLsRepository"))

		mockMiniURLRepository.On("CreateNewMiniURL", mock.Anything, mock.AnythingOfType("string")).
			Return("/abc123", nil)

		shortenerUseCase := NewShortenerUseCase(mockMiniURLRepository)

		miniURLResponse, err := shortenerUseCase.BuildMiniURL(context.Background(), "www.google.cl")

		assert.Nil(t, err)
		assert.Equal(t, "/abc123", miniURLResponse.MiniURL)
		assert.Equal(t, "https://diego.sepu", miniURLResponse.Host)
	})

	t.Run("when build short url use case gets an unwanted error from get miniurl if exists", func(t *testing.T) {
		mockMiniURLRepository := new(miniURLsRepositoryMock)

		mockMiniURLRepository.On("GetIfExists", mock.Anything, mock.AnythingOfType("string")).
			Return("", shared.BuildError(http.StatusInternalServerError, shared.DatabaseFindError, "db error", "miniURLsRepository"))

		shortenerUseCase := NewShortenerUseCase(mockMiniURLRepository)

		miniURLResponse, err := shortenerUseCase.BuildMiniURL(context.Background(), "www.google.cl")

		mockMiniURLRepository.AssertNotCalled(t, "CreateNewMiniURL")

		assert.NotNil(t, err)
		assert.Nil(t, miniURLResponse)
		assert.Equal(t, http.StatusInternalServerError, err.(*shared.ApplicationError).HTTPStatusCode)
		assert.Equal(t, shared.DatabaseFindError, err.(*shared.ApplicationError).ErrorCode)
		assert.Equal(t, "db error", err.(*shared.ApplicationError).ErrorDescription)
		assert.Equal(t, "miniURLsRepository", err.(*shared.ApplicationError).ErrorOrigin)
	})

	t.Run("when build short url use case gets an unwanted error from create new miniurl", func(t *testing.T) {
		mockMiniURLRepository := new(miniURLsRepositoryMock)

		mockMiniURLRepository.On("GetIfExists", mock.Anything, mock.AnythingOfType("string")).
			Return("", shared.BuildError(http.StatusNotFound, shared.DatabaseNotFoundError, "not found in db", "miniURLsRepository"))

		mockMiniURLRepository.On("CreateNewMiniURL", mock.Anything, mock.AnythingOfType("string")).
			Return("", shared.BuildError(http.StatusInternalServerError, shared.DatabaseFindError, "db error", "miniURLsRepository"))

		shortenerUseCase := NewShortenerUseCase(mockMiniURLRepository)

		miniURLResponse, err := shortenerUseCase.BuildMiniURL(context.Background(), "www.google.cl")

		assert.NotNil(t, err)
		assert.Nil(t, miniURLResponse)
		assert.Equal(t, http.StatusInternalServerError, err.(*shared.ApplicationError).HTTPStatusCode)
		assert.Equal(t, shared.DatabaseFindError, err.(*shared.ApplicationError).ErrorCode)
		assert.Equal(t, "db error", err.(*shared.ApplicationError).ErrorDescription)
		assert.Equal(t, "miniURLsRepository", err.(*shared.ApplicationError).ErrorOrigin)
	})
}
