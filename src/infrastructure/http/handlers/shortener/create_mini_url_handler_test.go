package shortener

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/DiegoSepuSoto/mini-url-builder-api/src/domain/models"
)

func TestCreateMiniURLHandler(t *testing.T) {
	t.Run("when create mini url handler is executed and returns as expected", func(t *testing.T) {
		e := echo.New()
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/mini-url", strings.NewReader(`{"api_key":"xyz789","original_url": "www.google.cl"}`))

		c := e.NewContext(req, rec)

		mockShortenerUseCase := new(shortenerUseCaseMock)

		mockShortenerUseCase.On("BuildMiniURL", mock.Anything, mock.AnythingOfType("string")).
			Return(&models.MiniURLResponse{
				Host:    "https://diego.sepu",
				MiniURL: "abc123",
			}, nil)

		shortenerHandler := NewShortenerHandler(e, mockShortenerUseCase)

		err := shortenerHandler.CreateMiniURL(c)

		assert.Nil(t, err)

		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String(), `{"host":"https://diego.sepu","mini_url":"/abc123"}`)
	})

	t.Run("when create mini url handler sends an error from json decode", func(t *testing.T) {
		e := echo.New()
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/mini-url", strings.NewReader(`<>`))

		c := e.NewContext(req, rec)

		mockShortenerUseCase := new(shortenerUseCaseMock)

		mockShortenerUseCase.On("BuildMiniURL", mock.Anything, mock.AnythingOfType("string")).
			Return(&models.MiniURLResponse{
				Host:    "https://diego.sepu",
				MiniURL: "abc123",
			}, nil)

		shortenerHandler := NewShortenerHandler(e, mockShortenerUseCase)

		err := shortenerHandler.CreateMiniURL(c)

		assert.Nil(t, err)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Contains(t, rec.Body.String(), "there was an error processing the request")
	})

	t.Run("when create mini url handler sends an error from use case", func(t *testing.T) {
		e := echo.New()
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/mini-url", strings.NewReader(`{"original_url": "www.google.cl"}`))

		c := e.NewContext(req, rec)

		mockShortenerUseCase := new(shortenerUseCaseMock)

		mockShortenerUseCase.On("BuildMiniURL", mock.Anything, mock.AnythingOfType("string")).
			Return(nil, errors.New("internal logic error"))

		shortenerHandler := NewShortenerHandler(e, mockShortenerUseCase)

		err := shortenerHandler.CreateMiniURL(c)

		assert.Nil(t, err)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Contains(t, rec.Body.String(), "there was an error creating the mini url")
	})
}
