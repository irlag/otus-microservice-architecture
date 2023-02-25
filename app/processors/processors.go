package processors

import (
	"go.uber.org/zap"

	db "otus-microservice-architecture/app/storage/db/sqlc"
)

type Processors struct {
	HealthcheckProcessor HealthcheckProcessor
	ProductsProcessor    ProductsProcessor
}

func NewProcessor(
	store db.Store,
	log *zap.Logger,
) *Processors {
	return &Processors{
		HealthcheckProcessor: NewHealtcheckProcessor(),
		ProductsProcessor:    NewProductsProcessor(store),
	}
}
