package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"go.uber.org/zap"

	"otus-microservice-architecture/app/api"
	"otus-microservice-architecture/app/config"
	appProcessors "otus-microservice-architecture/app/processors"
	db "otus-microservice-architecture/app/storage/db/sqlc"
)

type Server struct {
	config *config.Config
	Logger *zap.Logger
	Router *mux.Router
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

	store := db.NewStore()
	err = store.Open(config.DB)
	if err != nil {
		logger.Fatal("can't initialize db store", zap.Error(err))
	}

	server.Router = NewRouter()

	processors := appProcessors.NewProcessor(store, logger)

	api.NewHealthcheckApi(processors).HandleMethods(server.Router)
	api.NewProductsApi(processors).HandleMethods(server.Router)

	return server
}

func (s *Server) Start() error {
	url := fmt.Sprintf("%s:%s", s.config.BindAddress, s.config.Port)

	s.Logger.Info(fmt.Sprintf("starting api server at %s", url))

	return http.ListenAndServe(url,
		handlers.CompressHandler(
			handlers.LoggingHandler(os.Stdout, s.Router),
		),
	)
}
