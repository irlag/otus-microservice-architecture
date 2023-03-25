package authenticator

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"

	"otus-microservice-architecture/app/models"
)

func (a *Service) Authenticate(request *http.Request) (int, error) {

	tokenString, err := a.authHeaderTokenExtractor(request)
	if err != nil {
		return http.StatusBadRequest, err
	}

	token, err := a.parse(tokenString)
	if err != nil {
		return http.StatusUnauthorized, err
	}

	request = request.Clone(context.WithValue(request.Context(), models.RequestAuthToken, token))

	return http.StatusOK, nil
}

func (a *Service) authHeaderTokenExtractor(request *http.Request) (string, error) {
	authHeader := request.Header.Get("Authorization")
	if authHeader == "" {
		return "", AuthorizationHeaderNotSetError
	}

	authHeaderParts := strings.Fields(authHeader)
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", TokenBearerFormatError
	}
	if authHeaderParts[1] == "" {
		return "", TokenBearerEmptyError
	}

	return authHeaderParts[1], nil
}

func (a *Service) parse(tokenString string) (interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, UnexpectedSigningMethod
		}

		return []byte(a.config.JWTSecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if token.Valid {
		return token, nil
	}

	return nil, InvalidTokenError
}
