package miniurls

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/DiegoSepuSoto/mini-url-builder-api/src/shared"
)

func TestCreateNewMiniURLRepository(t *testing.T) {
	t.Run("when create new URL helper function works as expected then returns an expected miniURL", func(t *testing.T) {
		viper.Set("COUNTER_VALUE", 130000)
		viper.Set("COUNTER_MAX", 140000)

		newMiniURL := createNewURL()

		assert.Equal(t, "000xOM", newMiniURL)
	})

	t.Run("when create new mini url repository function works as expected then returns the expected miniURL", func(t *testing.T) {
		viper.Set("COUNTER_VALUE", 130000)
		viper.Set("COUNTER_MAX", 140000)

		mockMongoCollection := new(mongoCollectionMock)

		mockMongoCollection.On("InsertOne", mock.Anything, mock.AnythingOfType("entities.MiniURLRecord"), mock.AnythingOfType("[]*options.InsertOneOptions")).
			Return(&mongo.InsertOneResult{InsertedID: "abc123"}, nil)

		miniURLsRepository := NewMiniURLsRepository(mockMongoCollection)

		miniURL, err := miniURLsRepository.CreateNewMiniURL(context.Background(), "www.google.cl")

		assert.Nil(t, err)
		assert.Equal(t, "/000xOM", miniURL)
	})

	t.Run("when create new mini url repository function gets an error from mongodb collection then returns the error", func(t *testing.T) {
		viper.Set("COUNTER_VALUE", 130000)
		viper.Set("COUNTER_MAX", 140000)

		mockMongoCollection := new(mongoCollectionMock)

		mockMongoCollection.On("InsertOne", mock.Anything, mock.AnythingOfType("entities.MiniURLRecord"), mock.AnythingOfType("[]*options.InsertOneOptions")).
			Return(nil, errors.New("database error"))

		miniURLsRepository := NewMiniURLsRepository(mockMongoCollection)

		miniURL, err := miniURLsRepository.CreateNewMiniURL(context.Background(), "www.google.cl")

		assert.NotNil(t, err)
		assert.Equal(t, "", miniURL)
		assert.Equal(t, http.StatusInternalServerError, err.(*shared.ApplicationError).HTTPStatusCode)
		assert.Equal(t, shared.DatabaseInsertError, err.(*shared.ApplicationError).ErrorCode)
		assert.Equal(t, "database error", err.(*shared.ApplicationError).ErrorDescription)
		assert.Equal(t, "miniURLsRepository", err.(*shared.ApplicationError).ErrorOrigin)
	})
}
