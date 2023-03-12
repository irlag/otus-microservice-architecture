package product

import (
	"context"
	"database/sql"

	"otus-microservice-architecture/app/api/parameters"
	db "otus-microservice-architecture/app/storage/db/sqlc"
)

func (p *productProcessor) Create(ctx context.Context, params parameters.ProductsCreateParams) (db.Product, error) {
	productFromRequest := db.CreateProductParams{
		Name:            params.Name,
		Article:         params.Article,
		Brand:           params.Brand,
		CountryOfOrigin: params.CountryOfOrigin,
		Description: sql.NullString{
			String: params.Description,
			Valid:  true,
		},
		Price: params.Price,
	}

	return p.store.CreateProduct(ctx, productFromRequest)
}
