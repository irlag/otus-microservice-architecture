package server

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"

	"otus-microservice-architecture/app/api"
	"otus-microservice-architecture/app/api/responses"
	"otus-microservice-architecture/app/models"
	"otus-microservice-architecture/app/services"
)

type Middlewares struct {
	Services    *services.Services
	HttpMetrics *HttpMetrics
}

func NewMiddlewares(services *services.Services, httpMetrics *HttpMetrics) *Middlewares {
	return &Middlewares{Services: services, HttpMetrics: httpMetrics}
}

func (m *Middlewares) ContentTypeApplicationJsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		next.ServeHTTP(w, r)
	})
}

func (m *Middlewares) ResponseMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			respWriter := NewResponseWriter(w)

			defer func() {
				if api.AppRoutes["metrics"].Path == r.RequestURI {
					return
				}

				startedAt := r.Context().Value(models.RequestStartedAtKey).(time.Time)

				m.HttpMetrics.statReqDurations.With(
					prometheus.Labels{"method": r.RequestURI},
				).Observe(time.Since(startedAt).Seconds())

				m.HttpMetrics.statReqCount.With(
					prometheus.Labels{"method": r.RequestURI, "code": strconv.Itoa(respWriter.statusCode)},
				).Inc()
			}()

			next.ServeHTTP(respWriter, r)
		})
	}
}

func (m *Middlewares) StartedAtMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startedAt := time.Now()
			r = r.Clone(context.WithValue(r.Context(), models.RequestStartedAtKey, startedAt))

			next.ServeHTTP(w, r)
		})
	}
}

func (m *Middlewares) JwtAuthMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			currentRoute, err := api.AppRoutes.Find(mux.CurrentRoute(r).GetName())
			if err != nil {
				responses.NewErrorResponse(http.StatusBadRequest, err).WriteErrorResponse(w)

				return
			}

			if currentRoute.Secure {
				status, err := m.Services.Authenticator.Authenticate(r)
				if err != nil {
					responses.NewErrorResponse(status, err).WriteErrorResponse(w)

					return
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}
