package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type appMetrics struct {
	appErrors prometheus.Counter
}

type Metrics interface {
	MustRegisterMetrics(registry *prometheus.Registry)
	AppError()
}

func New() Metrics {
	m := &appMetrics{}

	m.appErrors = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "app_errors",
		Help: "Application errors",
	})

	return m
}

func (m *appMetrics) MustRegisterMetrics(registry *prometheus.Registry) {
	registry.MustRegister(m.appErrors)
}

func (m *appMetrics) AppError() {
	m.appErrors.Inc()
}
