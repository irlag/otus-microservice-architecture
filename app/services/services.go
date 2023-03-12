package services

import (
	"go.uber.org/zap"

	"otus-microservice-architecture/app/config"
	"otus-microservice-architecture/app/services/authenticator"
	"otus-microservice-architecture/app/services/password"
)

type Services struct {
	PasswordHash  password.PasswordHash
	Authenticator authenticator.Authenticator
}

func New(log *zap.Logger, config *config.Config) *Services {
	return &Services{
		PasswordHash: password.New(password.Config{Log: log}),
		Authenticator: authenticator.New(authenticator.Config{
			Log:    log,
			Config: config.Auth,
		}),
	}
}
