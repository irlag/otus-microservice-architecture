package user

import (
	"context"

	"otus-microservice-architecture/app/api/parameters"
	appErrors "otus-microservice-architecture/app/errors"
)

func (u *userProcessor) Delete(ctx context.Context, params parameters.UserDeleteParams) error {
	rowsAffected, err := u.store.DeleteUser(ctx, params.UserID)
	if err != nil {
		return err
	}

	if rowsAffected <= 0 {
		return appErrors.UserNotDeletedError
	}

	return nil
}
