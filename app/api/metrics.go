package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Metrics struct {
	Prometheus *prometheus.Registry
}

func NewMetricsApi(prometheus *prometheus.Registry) *Metrics {
	return &Metrics{Prometheus: prometheus}
}

func (m *Metrics) HandleMethods(router *mux.Router) {
	router.Handle(AppRoutes["metrics"].Path, m.Metrics()).
		Methods(AppRoutes["metrics"].Method).
		Name(AppRoutes["metrics"].Name)
}

func (m *Metrics) Metrics() http.Handler {
	return promhttp.HandlerFor(m.Prometheus, promhttp.HandlerOpts{})
}
