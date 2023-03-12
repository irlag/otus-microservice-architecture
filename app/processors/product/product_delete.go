package product

import (
	"context"

	"otus-microservice-architecture/app/api/parameters"
	appErrors "otus-microservice-architecture/app/errors"
)

func (p *productProcessor) Delete(ctx context.Context, params parameters.ProductsDeleteParams) error {
	rowsAffected, err := p.store.DeleteProduct(ctx, params.ProductID)
	if err != nil {
		return err
	}

	if rowsAffected <= 0 {
		return appErrors.ProductNotDeletedError
	}

	return nil
}
