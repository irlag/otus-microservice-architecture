package processors

import (
	"go.uber.org/zap"

	"otus-microservice-architecture/app/config"
	"otus-microservice-architecture/app/processors/auth"
	"otus-microservice-architecture/app/processors/product"
	"otus-microservice-architecture/app/processors/user"
	"otus-microservice-architecture/app/services"
	db "otus-microservice-architecture/app/storage/db/sqlc"
)

type Processors struct {
	HealthcheckProcessor HealthcheckProcessor
	UserProcessor        user.User
	ProductProcessor     product.Product
	AuthProcessor        auth.Auth
}

func NewProcessor(
	store db.Store,
	services *services.Services,
	log *zap.Logger,
	config *config.Config,
) *Processors {
	return &Processors{
		HealthcheckProcessor: NewHealtcheckProcessor(),
		ProductProcessor:     product.NewProductProcessor(store),
		AuthProcessor:        auth.NewAuthProcessor(store, services, config.Auth),
		UserProcessor:        user.NewUserProcessor(store, services),
	}
}
