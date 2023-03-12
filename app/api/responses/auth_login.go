package responses

//go:generate easyjson

import (
	"net/http"

	"otus-microservice-architecture/app/models"
)

//easyjson:json
type AuthLoginOkResponse struct {
	Token string `json:"token"`
}

func NewAuthLoginOkResponse(token models.Token) AuthLoginOkResponse {
	return AuthLoginOkResponse{
		Token: token.Token,
	}
}

func (a *AuthLoginOkResponse) WriteResponse(rw http.ResponseWriter) {
	payload, _ := a.MarshalJSON()

	WriteJsonResponse(rw, http.StatusOK, payload)
}
