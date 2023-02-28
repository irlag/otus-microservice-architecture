package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"otus-microservice-architecture/app/api"
	"otus-microservice-architecture/app/config"
	appProcessors "otus-microservice-architecture/app/processors"
)

type Server struct {
	config *config.Config
	Logger *zap.Logger
	Router *mux.Router
}

func New(config *config.Config) *Server {
	logger, err := NewLogger(config.Debug)

	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	server := &Server{
		config: config,
		Logger: logger,
	}

	server.Router = NewRouter()

	processors := appProcessors.NewProcessor(logger)

	api.NewHealthcheckApi(processors).HandleMethods(server.Router)

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
