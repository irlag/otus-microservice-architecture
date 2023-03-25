package product

import (
	"context"

	"otus-microservice-architecture/app/api/parameters"
	db "otus-microservice-architecture/app/storage/db/sqlc"
)

type Product interface {
	Get(ctx context.Context, id int64) (product db.Product, err error)
	List(ctx context.Context, limit int64, offset int64) (products []db.Product, err error)
	ListCount(ctx context.Context) (int64, error)
	Create(ctx context.Context, params parameters.ProductsCreateParams) (db.Product, error)
	Update(ctx context.Context, params parameters.ProductsUpdateParams) (db.Product, error)
	Delete(ctx context.Context, params parameters.ProductsDeleteParams) error
}

type productProcessor struct {
	store db.Store
}

func NewProductProcessor(store db.Store) Product {
	return &productProcessor{
		store: store,
	}
}
