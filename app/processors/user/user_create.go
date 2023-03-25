package user

import (
	"context"
	"database/sql"

	"otus-microservice-architecture/app/api/parameters"
	db "otus-microservice-architecture/app/storage/db/sqlc"
)

func (u *userProcessor) Create(ctx context.Context, params *parameters.UserCreateParams) (db.CreateUserRow, error) {
	user := db.CreateUserRow{}
	createParams := db.CreateUserParams{
		Email:     params.Email,
		FirstName: params.FirstName,
		LastName: func() sql.NullString {
			lastName := sql.NullString{}

			if params != nil && params.LastName != "" {
				lastName.String = params.LastName
				lastName.Valid = true
			}

			return lastName
		}(),
	}

	passwordSalt, err := u.services.PasswordHash.GetRandomSalt()
	if err != nil {
		return user, err
	}

	passwordHash, err := u.services.PasswordHash.GeneratePasswordHashWithSalt(params.Password, passwordSalt)
	if err != nil {
		return user, err
	}

	createParams.Password = passwordHash

	user, err = u.store.CreateUser(ctx, createParams)
	if err != nil {
		return user, err
	}

	return user, nil
}
