package shared

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var APIRequestsMetrics = promauto.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "mini_url_builder_api",
		Help: "Mini URL Builder API HTTP requests metrics",
	}, []string{"request_url", "request_with_error", "http_method"},
)
