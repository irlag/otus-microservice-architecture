package product

import (
	"database/sql"
	"net/http"

	"otus-microservice-architecture/app/api/parameters"
	"otus-microservice-architecture/app/api/responses"
	appErrors "otus-microservice-architecture/app/errors"
)

func (p *Product) Delete() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		productsDeleteParams, err := parameters.NewProductsDeleteParamsFromRequest(request)
		if err != nil {
			responses.NewErrorResponse(http.StatusBadRequest, err).WriteErrorResponse(writer)

			return
		}

		_, err = p.processors.ProductProcessor.Get(request.Context(), productsDeleteParams.ProductID)
		if err != nil {
			response := &responses.ErrorResponse{}
			switch err {
			case sql.ErrNoRows:
				response = responses.NewErrorResponse(http.StatusNotFound, appErrors.ProductNotFoundError)
			default:
				response = responses.NewErrorResponse(http.StatusInternalServerError, err)
			}
			response.WriteErrorResponse(writer)

			return
		}

		err = p.processors.ProductProcessor.Delete(request.Context(), productsDeleteParams)
		if err != nil {
			responses.NewErrorResponse(http.StatusInternalServerError, err).WriteErrorResponse(writer)

			return
		}

		productDeleteOkResponse := responses.NewProductDeleteOkResponse()
		productDeleteOkResponse.WriteResponse(writer)
	}
}
