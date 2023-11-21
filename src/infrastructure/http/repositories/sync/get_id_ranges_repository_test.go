package sync

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/DiegoSepuSoto/mini-url-builder-api/src/shared"
)

func TestGetIDRangesRepository(t *testing.T) {
	t.Run("when get id ranges gets executed as expected", func(t *testing.T) {
		mockHTTPClient := new(HTTPClientMock)

		mockResponse := "{\"rangeIDs\": \"10-20\"}"

		mockHTTPClient.On("Get", mock.AnythingOfType("string")).Return(&http.Response{Body: io.NopCloser(bytes.NewBuffer([]byte(mockResponse)))}, nil)

		syncRepository := NewSyncRepository(mockHTTPClient)

		ranges, err := syncRepository.GetIDRanges()

		assert.Nil(t, err)
		assert.Equal(t, "10-20", ranges)
	})

	t.Run("when get id ranges gets an error from http client", func(t *testing.T) {
		mockHTTPClient := new(HTTPClientMock)

		mockHTTPClient.On("Get", mock.AnythingOfType("string")).Return(nil, errors.New("http error"))

		syncRepository := NewSyncRepository(mockHTTPClient)

		ranges, err := syncRepository.GetIDRanges()

		assert.NotNil(t, err)
		assert.Equal(t, "", ranges)
		assert.Equal(t, http.StatusInternalServerError, err.(*shared.ApplicationError).HTTPStatusCode)
		assert.Equal(t, shared.SyncCommunicationError, err.(*shared.ApplicationError).ErrorCode)
		assert.Equal(t, "http error", err.(*shared.ApplicationError).ErrorDescription)
		assert.Equal(t, "syncRepository", err.(*shared.ApplicationError).ErrorOrigin)
	})
}
