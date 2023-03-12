package product

import (
	"context"

	db "otus-microservice-architecture/app/storage/db/sqlc"
)

func (p *productProcessor) Get(ctx context.Context, id int64) (product db.Product, err error) {
	product, err = p.store.GetProduct(ctx, id)
	if err != nil {
		return product, err
	}

	return product, nil
}
