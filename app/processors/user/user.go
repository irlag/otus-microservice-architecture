package user

import (
	"context"

	"otus-microservice-architecture/app/api/parameters"
	"otus-microservice-architecture/app/services"
	db "otus-microservice-architecture/app/storage/db/sqlc"
)

type User interface {
	Create(ctx context.Context, params *parameters.UserCreateParams) (db.CreateUserRow, error)
	Get(ctx context.Context, id int64) (db.User, error)
	Update(ctx context.Context, params parameters.UserUpdateParams) (db.User, error)
	Delete(ctx context.Context, params parameters.UserDeleteParams) error
}

type userProcessor struct {
	store    db.Store
	services *services.Services
}

func NewUserProcessor(store db.Store, services *services.Services) User {
	return &userProcessor{
		store:    store,
		services: services,
	}
}
