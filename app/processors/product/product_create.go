package product

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"otus-microservice-architecture/app/api/parameters"
	db "otus-microservice-architecture/app/storage/db/sqlc"
)

func (p *productProcessor) Create(ctx context.Context, params parameters.ProductsCreateParams) (db.Product, error) {
	version, err := uuid.NewUUID()
	if err != nil {
		return db.Product{}, err
	}

	productFromRequest := db.CreateProductParams{
		Name:            params.Name,
		Article:         params.Article,
		Brand:           params.Brand,
		CountryOfOrigin: params.CountryOfOrigin,
		Description: sql.NullString{
			String: params.Description,
			Valid:  true,
		},
		Price:   params.Price,
		Version: version,
	}

	return p.store.CreateProduct(ctx, productFromRequest)
}
