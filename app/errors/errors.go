package errors

import "errors"

var (
	ProductNotFound   = errors.New("product not found")
	ProductNotDeleted = errors.New("product not deleted")
)
