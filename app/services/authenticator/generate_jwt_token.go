package authenticator

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (a *Service) GenerateJWTToken(user string) (string, error) {
	payload := jwt.MapClaims{
		"sub":  JWTAuthSubject,
		"exp":  time.Now().Add(a.config.TokenTTL).Unix(),
		"iat":  time.Now().Unix(),
		"user": user,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	t, err := token.SignedString([]byte(a.config.JWTSecretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}
