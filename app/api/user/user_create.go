package product

import (
	"net/http"

	"otus-microservice-architecture/app/api/parameters"
	"otus-microservice-architecture/app/api/responses"
)

func (u *User) Create() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		userCreateParams, err := parameters.NewUserCreateParamsFromRequest(request)
		if err != nil {
			responses.NewErrorResponse(http.StatusBadRequest, err).WriteErrorResponse(writer)

			return
		}

		createUserRow, err := u.processors.UserProcessor.Create(request.Context(), userCreateParams)
		if err != nil {
			responses.NewErrorResponse(http.StatusInternalServerError, err).WriteErrorResponse(writer)

			return
		}

		userOkResponse := responses.NewUserCreateOkResponse(createUserRow)
		userOkResponse.WriteResponse(writer)
	}
}
