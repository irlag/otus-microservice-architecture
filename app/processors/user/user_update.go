package user

import (
	"context"
	"database/sql"

	"otus-microservice-architecture/app/api/parameters"
	db "otus-microservice-architecture/app/storage/db/sqlc"
)

func (u *userProcessor) Update(ctx context.Context, params parameters.UserUpdateParams) (db.User, error) {
	userFromRequest := db.UpdateUserParams{
		ID:        params.UserID,
		Email:     params.Email,
		FirstName: params.FirstName,
		LastName: sql.NullString{
			String: params.LastName,
			Valid:  true,
		},
	}

	return u.store.UpdateUser(ctx, userFromRequest)
}
