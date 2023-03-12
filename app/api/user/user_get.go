package product

import (
	"database/sql"
	"net/http"

	"otus-microservice-architecture/app/api/parameters"
	"otus-microservice-architecture/app/api/responses"
	appErrors "otus-microservice-architecture/app/errors"
)

func (u *User) Get() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		userCreateParams, err := parameters.NewUserGetParamsFromRequest(request)
		if err != nil {
			responses.NewErrorResponse(http.StatusBadRequest, err).WriteErrorResponse(writer)

			return
		}

		user, err := u.processors.UserProcessor.Get(request.Context(), userCreateParams.UserID)
		if err != nil {
			response := &responses.ErrorResponse{}
			switch err {
			case sql.ErrNoRows:
				response = responses.NewErrorResponse(http.StatusNotFound, appErrors.UserNotFoundError)
			default:
				response = responses.NewErrorResponse(http.StatusInternalServerError, err)
			}
			response.WriteErrorResponse(writer)

			return
		}

		userOkResponse := responses.NewUserGetOkResponse(user)
		userOkResponse.WriteResponse(writer)
	}
}
