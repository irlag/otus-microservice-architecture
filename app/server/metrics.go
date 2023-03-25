package server

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

type HttpMetrics struct {
	statReqCount     *prometheus.CounterVec
	statReqDurations *prometheus.HistogramVec
}

func (s *Server) configurePrometheus() {
	s.Prometheus = prometheus.NewRegistry()
	s.Prometheus.MustRegister(collectors.NewGoCollector())
	s.Prometheus.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
}

func (s *Server) initializeMetrics() {
	s.HttpMetrics = &HttpMetrics{
		statReqCount: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of requests",
		}, []string{"code", "method"}),
		statReqDurations: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Request duration",
			Buckets: []float64{0.005, 0.01, 0.05, 0.1, 0.5, 1, 5},
		}, []string{"method"}),
	}

	s.Prometheus.MustRegister(
		s.HttpMetrics.statReqCount,
		s.HttpMetrics.statReqDurations,
	)
}
