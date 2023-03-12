package errors

import "errors"

var (
	ProductNotFoundError      = errors.New("product not found")
	ProductNotDeletedError    = errors.New("product not deleted")
	UserNotFoundError         = errors.New("user not found")
	UserNotDeletedError       = errors.New("user not deleted")
	InvalidCredentionalsError = errors.New("invalid user or password")
)
