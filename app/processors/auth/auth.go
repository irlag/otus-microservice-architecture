package auth

import (
	"context"

	"otus-microservice-architecture/app/api/parameters"
	"otus-microservice-architecture/app/config"
	"otus-microservice-architecture/app/models"
	"otus-microservice-architecture/app/services"
	db "otus-microservice-architecture/app/storage/db/sqlc"
)

type Auth interface {
	Login(ctx context.Context, params *parameters.AuthTokenParams) (models.Token, error)
}

type authProcessor struct {
	store    db.Store
	services *services.Services
	config   *config.AuthConfig
}

func NewAuthProcessor(store db.Store, services *services.Services, config *config.AuthConfig) Auth {
	return &authProcessor{
		store:    store,
		services: services,
		config:   config,
	}
}
