package miniurls

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/DiegoSepuSoto/mini-url-builder-api/src/infrastructure/database/repositories/mongodb/miniurls/entities"
	"github.com/DiegoSepuSoto/mini-url-builder-api/src/shared"
)

func TestCheckIfExistsRepository(t *testing.T) {
	t.Run("when check if exists gets executed as expected and finds the record in the database", func(t *testing.T) {
		mockMongoCollection := new(mongoCollectionMock)
		mockMongoSingleResult := new(mongoSingleResultMock)

		mockMongoSingleResult.On("Decode", mock.AnythingOfType("*entities.MiniURLRecord")).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*entities.MiniURLRecord)

			arg.NewURL = "abc123"
		}).Return(nil)

		mockMongoCollection.On("FindOne", mock.Anything, mock.Anything, mock.AnythingOfType("[]*options.FindOneOptions")).
			Return(mockMongoSingleResult)

		miniURLsRepository := NewMiniURLsRepository(mockMongoCollection)

		miniURL, err := miniURLsRepository.GetIfExists(context.Background(), "www.google.cl")

		assert.Nil(t, err)
		assert.Equal(t, "abc123", miniURL)
	})

	t.Run("when check if exists gets executed as expected and does not find record in the database", func(t *testing.T) {
		mockMongoCollection := new(mongoCollectionMock)
		mockMongoSingleResult := new(mongoSingleResultMock)

		mockMongoSingleResult.On("Decode", mock.AnythingOfType("*entities.MiniURLRecord")).Return(errors.New("mongo: no documents in result"))

		mockMongoCollection.On("FindOne", mock.Anything, mock.Anything, mock.AnythingOfType("[]*options.FindOneOptions")).
			Return(mockMongoSingleResult)

		miniURLsRepository := NewMiniURLsRepository(mockMongoCollection)

		miniURL, err := miniURLsRepository.GetIfExists(context.Background(), "www.google.cl")

		assert.NotNil(t, err)
		assert.Equal(t, "", miniURL)
		assert.Equal(t, http.StatusNotFound, err.(*shared.ApplicationError).HTTPStatusCode)
		assert.Equal(t, shared.DatabaseNotFoundError, err.(*shared.ApplicationError).ErrorCode)
		assert.Equal(t, "mongo: no documents in result", err.(*shared.ApplicationError).ErrorDescription)
		assert.Equal(t, "miniURLsRepository", err.(*shared.ApplicationError).ErrorOrigin)
	})

	t.Run("when check if exists gets an error finding record then returns the error", func(t *testing.T) {
		mockMongoCollection := new(mongoCollectionMock)
		mockMongoSingleResult := new(mongoSingleResultMock)

		mockMongoSingleResult.On("Decode", mock.AnythingOfType("*entities.MiniURLRecord")).Return(errors.New("mongo: error"))

		mockMongoCollection.On("FindOne", mock.Anything, mock.Anything, mock.AnythingOfType("[]*options.FindOneOptions")).
			Return(mockMongoSingleResult)

		miniURLsRepository := NewMiniURLsRepository(mockMongoCollection)

		miniURL, err := miniURLsRepository.GetIfExists(context.Background(), "www.google.cl")

		assert.NotNil(t, err)
		assert.Equal(t, "", miniURL)
		assert.Equal(t, http.StatusInternalServerError, err.(*shared.ApplicationError).HTTPStatusCode)
		assert.Equal(t, shared.DatabaseFindError, err.(*shared.ApplicationError).ErrorCode)
		assert.Equal(t, "mongo: error", err.(*shared.ApplicationError).ErrorDescription)
		assert.Equal(t, "miniURLsRepository", err.(*shared.ApplicationError).ErrorOrigin)
	})
}
