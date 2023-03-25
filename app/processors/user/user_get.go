package user

import (
	"context"

	db "otus-microservice-architecture/app/storage/db/sqlc"
)

func (u *userProcessor) Get(ctx context.Context, id int64) (user db.User, err error) {
	user, err = u.store.GetUser(ctx, id)
	if err != nil {
		return db.User{}, err
	}

	return user, nil
}
