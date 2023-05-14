package product

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"otus-microservice-architecture/app/api/parameters"
	appErrors "otus-microservice-architecture/app/errors"
	db "otus-microservice-architecture/app/storage/db/sqlc"
)

func (p *productProcessor) Update(ctx context.Context, params parameters.ProductsUpdateParams) (db.Product, error) {
	version, err := uuid.NewUUID()
	if err != nil {
		return db.Product{}, err
	}

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
		Price:   params.Price,
		Version: version,
	}
	tx, err := p.store.Begin()
	if err != nil {
		return db.Product{}, err
	}
	defer tx.Rollback()

	qtx := p.store.WithTx(tx)

	currentProduct, err := qtx.GetProductForUpdate(ctx, productFromRequest.ID)
	if err != nil {
		return db.Product{}, err
	}

	if currentProduct.Version.String() != params.Version {
		return db.Product{}, appErrors.ProductVersionError
	}

	product, err := qtx.UpdateProduct(ctx, productFromRequest)
	if err != nil {
		return db.Product{}, err
	}

	err = tx.Commit()
	if err != nil {
		return db.Product{}, err
	}

	return product, nil
}
