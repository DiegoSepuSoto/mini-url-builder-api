package docs

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type swaggerHandler struct{}

func NewSwaggerHandler(e *echo.Echo) *swaggerHandler {
	h := new(swaggerHandler)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	return h
}
