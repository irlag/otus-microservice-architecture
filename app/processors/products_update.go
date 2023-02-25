package processors

import (
	"context"
	"database/sql"

	"otus-microservice-architecture/app/api/parameters"
	db "otus-microservice-architecture/app/storage/db/sqlc"
)

func (p *productsProcessor) Update(ctx context.Context, params parameters.ProductsUpdateParams) (db.Product, error) {
	productFromRequest := db.UpdateProductParams{
		ID:              params.ProductID,
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

	return p.store.UpdateProduct(ctx, productFromRequest)
}
