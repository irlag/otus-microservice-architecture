package parameters

//go:generate easyjson

import (
	"io"
	"net/http"
)

//easyjson:json
type UserCreateParams struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

func NewUserCreateParamsFromRequest(request *http.Request) (*UserCreateParams, error) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}

	userRequest := UserCreateParams{}
	err = userRequest.UnmarshalJSON(body)
	if err != nil {
		return nil, err
	}

	return &userRequest, nil
}
