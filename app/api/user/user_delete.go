package product

import (
	"database/sql"
	"net/http"

	"otus-microservice-architecture/app/api/parameters"
	"otus-microservice-architecture/app/api/responses"
	appErrors "otus-microservice-architecture/app/errors"
)

func (u *User) Delete() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		userDeleteParams, err := parameters.NewUserDeleteParamsFromRequest(request)
		if err != nil {
			responses.NewErrorResponse(http.StatusBadRequest, err).WriteErrorResponse(writer)

			return
		}

		_, err = u.processors.UserProcessor.Get(request.Context(), userDeleteParams.UserID)
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

		err = u.processors.UserProcessor.Delete(request.Context(), userDeleteParams)
		if err != nil {
			responses.NewErrorResponse(http.StatusInternalServerError, err).WriteErrorResponse(writer)

			return
		}

		userDeleteOkResponse := responses.NewUserDeleteOkResponse()
		userDeleteOkResponse.WriteResponse(writer)
	}
}
