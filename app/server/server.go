package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"

	"otus-microservice-architecture/app/api"
	authApi "otus-microservice-architecture/app/api/auth"
	productApi "otus-microservice-architecture/app/api/product"
	userApi "otus-microservice-architecture/app/api/user"
	"otus-microservice-architecture/app/config"
	"otus-microservice-architecture/app/metrics"
	appProcessors "otus-microservice-architecture/app/processors"
	"otus-microservice-architecture/app/services"
	db "otus-microservice-architecture/app/storage/db/sqlc"
)

type Server struct {
	config      *config.Config
	Prometheus  *prometheus.Registry
	Logger      *zap.Logger
	Router      *mux.Router
	Services    *services.Services
	HttpMetrics *HttpMetrics
}

func New(config *config.Config) *Server {
	logger, err := NewLogger(config.Debug)
	if err != nil {
		logger.Fatal("can't initialize zap logger", zap.Error(err))
	}

	server := &Server{
		config: config,
		Logger: logger,
	}

	server.configurePrometheus()
	server.initializeMetrics()

	appMetrics := metrics.New()
	appMetrics.MustRegisterMetrics(server.Prometheus)

	store := db.NewStore()
	err = store.Open(config.DB)
	if err != nil {
		logger.Fatal("can't initialize db store", zap.Error(err))
	}

	server.Router = NewRouter()

	server.Services = services.New(logger, config)
	processors := appProcessors.NewProcessor(store, server.Services, logger, config)

	api.NewMetricsApi(server.Prometheus).HandleMethods(server.Router)
	api.NewHealthcheckApi(processors).HandleMethods(server.Router)
	authApi.NewAuthApi(processors).HandleMethods(server.Router)
	productApi.NewProductApi(processors).HandleMethods(server.Router)
	userApi.NewUserApi(processors).HandleMethods(server.Router)

	return server
}

func (s *Server) Start() error {
	url := fmt.Sprintf("%s:%s", s.config.BindAddress, s.config.Port)

	s.Logger.Info(fmt.Sprintf("starting api server at %s", url))

	corsAllowOrigin := handlers.AllowedOrigins([]string{"*"})

	middlewares := NewMiddlewares(s.Services, s.HttpMetrics)

	s.Router.Use(
		middlewares.StartedAtMiddleware(),
		middlewares.ResponseMiddleware(),
		middlewares.JwtAuthMiddleware(),
	)

	return http.ListenAndServe(url,
		handlers.CORS(corsAllowOrigin)(
			middlewares.ContentTypeApplicationJsonMiddleware(
				handlers.CompressHandler(
					handlers.LoggingHandler(os.Stdout, s.Router),
				),
			),
		),
	)
}
