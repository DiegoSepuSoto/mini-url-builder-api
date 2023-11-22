package metrics

import (
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type metricsHandler struct{}

func NewMetricsHandler(e *echo.Echo) *metricsHandler {
	h := new(metricsHandler)

	e.GET("/metrics", h.getMetrics())

	return h
}

func (h *metricsHandler) getMetrics() echo.HandlerFunc {
	return echo.WrapHandler(promhttp.Handler())
}
