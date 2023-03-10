package processors

import (
	"context"

	db "otus-microservice-architecture/app/storage/db/sqlc"
)

func (p *productsProcessor) List(ctx context.Context, limit int64, offset int64) (products []db.Product, err error) {
	listDbProductParams := db.ListProductParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	products, err = p.store.ListProduct(ctx, listDbProductParams)
	if err != nil {
		return products, err
	}

	return products, nil
}

func (p *productsProcessor) ListCount(ctx context.Context) (int64, error) {
	count, err := p.store.ListProductCount(ctx)
	if err != nil {
		return 0, err
	}

	return count, nil
}
