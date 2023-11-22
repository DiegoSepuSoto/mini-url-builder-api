package middlewares

import (
	"strconv"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/DiegoSepuSoto/mini-url-builder-api/src/shared"
)

func APIMetricsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		isError := "false"
		startTime := time.Now()
		err := next(c)
		elapsedTime := time.Since(startTime)

		if err != nil {
			isError = "true"
		}

		path := c.Path()
		statusCode := c.Response().Status

		shared.APIRequestsMetrics.WithLabelValues(
			path, isError, strconv.Itoa(statusCode),
		).Observe(elapsedTime.Seconds())

		return err
	}
}
