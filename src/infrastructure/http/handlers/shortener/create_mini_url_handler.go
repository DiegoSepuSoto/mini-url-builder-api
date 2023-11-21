package shortener

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"

	"github.com/DiegoSepuSoto/mini-url-builder-api/src/infrastructure/http/handlers/shortener/entities"
)

func (h *shortenerHandler) CreateMiniURL(c echo.Context) error {
	ctx := context.Background()

	var createMiniURLRequest *entities.CreateMiniURLRequest
	err := c.Bind(&createMiniURLRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": fmt.Sprintf("there was an error processing the request: %s", err.Error())})
	}

	miniURLResponse, err := h.shortenerUseCase.BuildMiniURL(ctx, createMiniURLRequest.OriginalURL)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "there was an error creating the mini url"})
	}

	return c.JSON(http.StatusCreated, miniURLResponse)
}
