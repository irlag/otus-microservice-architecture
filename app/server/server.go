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
	authApi "otus-microservice-architecture/app/api/auth"
	productApi "otus-microservice-architecture/app/api/product"
	"otus-microservice-architecture/app/api/responses"
	userApi "otus-microservice-architecture/app/api/user"
	"otus-microservice-architecture/app/config"
	appProcessors "otus-microservice-architecture/app/processors"
	"otus-microservice-architecture/app/services"
	db "otus-microservice-architecture/app/storage/db/sqlc"
)

type Server struct {
	config   *config.Config
	Logger   *zap.Logger
	Router   *mux.Router
	Services *services.Services
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

	server.Services = services.New(logger, config)
	processors := appProcessors.NewProcessor(store, server.Services, logger, config)

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

	s.Router.Use(s.jwtAuthMiddleware)

	return http.ListenAndServe(url,
		handlers.CORS(corsAllowOrigin)(
			s.contentTypeApplicationJsonMiddleware(
				handlers.CompressHandler(
					handlers.LoggingHandler(os.Stdout, s.Router),
				),
			),
		),
	)
}

func (s *Server) contentTypeApplicationJsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		next.ServeHTTP(w, r)
	})
}

func (s *Server) jwtAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		currentRoute, err := api.AppRoutes.Find(mux.CurrentRoute(request).GetName())
		if err != nil {
			responses.NewErrorResponse(http.StatusBadRequest, err).WriteErrorResponse(writer)

			return
		}

		if currentRoute.Secure {
			status, err := s.Services.Authenticator.Authenticate(request)
			if err != nil {
				responses.NewErrorResponse(status, err).WriteErrorResponse(writer)

				return
			}
		}

		next.ServeHTTP(writer, request)
	})
}
