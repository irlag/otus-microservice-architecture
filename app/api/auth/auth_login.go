package product

import (
	"database/sql"
	"net/http"

	"otus-microservice-architecture/app/api/parameters"
	"otus-microservice-architecture/app/api/responses"
	appErrors "otus-microservice-architecture/app/errors"
)

func (a *Auth) Login() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		authTokenParams, err := parameters.NewAuthTokenParamsFromRequest(request)
		if err != nil {
			responses.NewErrorResponse(http.StatusBadRequest, err).WriteErrorResponse(writer)

			return
		}

		token, err := a.processors.AuthProcessor.Login(request.Context(), authTokenParams)
		if err != nil {
			response := &responses.ErrorResponse{}
			switch err {
			case sql.ErrNoRows:
				response = responses.NewErrorResponse(http.StatusBadRequest, appErrors.UserNotFoundError)
			case appErrors.InvalidCredentionalsError:
				response = responses.NewErrorResponse(http.StatusBadRequest, appErrors.InvalidCredentionalsError)
			default:
				response = responses.NewErrorResponse(http.StatusInternalServerError, err)
			}
			response.WriteErrorResponse(writer)

			return
		}

		loginOkResponse := responses.NewAuthLoginOkResponse(token)
		loginOkResponse.WriteResponse(writer)

	}
}
