package auth

import (
	"context"

	"otus-microservice-architecture/app/api/parameters"
	appErrors "otus-microservice-architecture/app/errors"
	"otus-microservice-architecture/app/models"
)

func (a *authProcessor) Login(ctx context.Context, params *parameters.AuthTokenParams) (models.Token, error) {
	user, err := a.store.GetUserByEmail(ctx, params.Email)
	if err != nil {
		return models.Token{}, err
	}

	success, err := a.services.PasswordHash.CheckPassword(params.Password, user.Password)
	if err != nil {
		return models.Token{}, err
	}

	tokenString := ""
	if success {
		tokenString, err = a.services.Authenticator.GenerateJWTToken(params.Email)
		if err != nil {
			return models.Token{}, err
		}
	} else {
		return models.Token{}, appErrors.InvalidCredentionalsError
	}

	return models.Token{
		Token: tokenString,
	}, nil
}
