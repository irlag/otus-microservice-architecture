package parameters

//go:generate easyjson

import (
	"io"
	"net/http"
)

//easyjson:json
type AuthTokenParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewAuthTokenParamsFromRequest(request *http.Request) (*AuthTokenParams, error) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}

	authTokenParams := AuthTokenParams{}
	err = authTokenParams.UnmarshalJSON(body)
	if err != nil {
		return nil, err
	}

	return &authTokenParams, nil
}
