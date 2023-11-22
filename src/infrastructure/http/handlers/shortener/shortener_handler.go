package shortener

import (
	"github.com/labstack/echo/v4"

	"github.com/DiegoSepuSoto/mini-url-builder-api/src/application/usecase"
	"github.com/DiegoSepuSoto/mini-url-builder-api/src/infrastructure/http/handlers/middlewares"
)

type shortenerHandler struct {
	shortenerUseCase usecase.ShortenerUseCase
}

func NewShortenerHandler(e *echo.Echo, shortenerUseCase usecase.ShortenerUseCase) *shortenerHandler {
	h := &shortenerHandler{
		shortenerUseCase: shortenerUseCase,
	}

	g := e.Group("", middlewares.APIMetricsMiddleware, middlewares.AuthMiddleware)

	g.POST("/mini-url", h.CreateMiniURL)

	return h
}
