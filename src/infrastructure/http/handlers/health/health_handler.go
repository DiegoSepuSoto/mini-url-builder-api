package health

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type healthHandler struct{}

type appDetails struct {
	AppName     string `json:"app_name"`
	Status      string `json:"status"`
	Version     string `json:"version"`
	Environment string `json:"environment"`
}

func NewHealthHandler(e *echo.Echo) {
	h := &healthHandler{}

	e.GET("/health", h.HealthCheck)
}

func (h *healthHandler) HealthCheck(c echo.Context) error {
	currentAppDetails := appDetails{
		AppName:     "mini-url-builder-api",
		Status:      "UP",
		Version:     os.Getenv("APP_VERSION"),
		Environment: os.Getenv("APP_ENV"),
	}

	return c.JSON(http.StatusOK, currentAppDetails)
}
