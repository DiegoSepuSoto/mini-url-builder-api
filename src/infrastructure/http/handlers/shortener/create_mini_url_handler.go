package shortener

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"

	"github.com/DiegoSepuSoto/mini-url-builder-api/src/domain/models"
	"github.com/DiegoSepuSoto/mini-url-builder-api/src/infrastructure/http/handlers/shortener/entities"
	"github.com/DiegoSepuSoto/mini-url-builder-api/src/shared"
)

// CreateMiniURL godoc
// @Summary      Create Mini URL
// @Description  Returns as an API Response the created mini URL from a given one
// @Tags         MiniURL
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Bearer token"
// @Param request body entities.CreateMiniURLRequest true "Request"
// @Success      200  {object}  models.MiniURLResponse "OK"
// @Failure      400  {object}  shared.EchoErrorResponse "Bad Request"
// @Failure      401  {object}  shared.EchoErrorResponse "Unauthorized"
// @Failure      500  {object}  shared.EchoErrorResponse "Application Error"
// @Router       /mini-url [post]
func (h *shortenerHandler) CreateMiniURL(c echo.Context) error {
	ctx := context.Background()

	var createMiniURLRequest *entities.CreateMiniURLRequest
	err := json.NewDecoder(c.Request().Body).Decode(&createMiniURLRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, shared.EchoErrorResponse{
			Message: fmt.Sprintf("there was an error processing the request: %s", err.Error()),
		})
	}

	miniURLResponse, err := h.shortenerUseCase.BuildMiniURL(ctx, createMiniURLRequest.OriginalURL)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, shared.EchoErrorResponse{
			Message: "there was an error creating the mini url",
		})
	}

	log.WithFields(
		log.Fields{"originalURL": createMiniURLRequest.OriginalURL, "miniURL": miniURLResponse.MiniURL},
	).Info("mini url created successfully")

	transformMiniURL(miniURLResponse)
	return c.JSON(http.StatusCreated, miniURLResponse)
}

func transformMiniURL(miniURLResponse *models.MiniURLResponse) {
	miniURLResponse.MiniURL = fmt.Sprintf("/%s", miniURLResponse.MiniURL)
}
