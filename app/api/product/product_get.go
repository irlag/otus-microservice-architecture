package product

import (
	"database/sql"
	"net/http"

	"otus-microservice-architecture/app/api/parameters"
	"otus-microservice-architecture/app/api/responses"
	appErrors "otus-microservice-architecture/app/errors"
)

func (p *Product) Get() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		productsGetParams, err := parameters.NewProductsGetParamsFromRequest(request)
		if err != nil {
			responses.NewErrorResponse(http.StatusBadRequest, err).WriteErrorResponse(writer)

			return
		}

		product, err := p.processors.ProductProcessor.Get(request.Context(), productsGetParams.ProductID)
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

		productOkResponse := responses.NewProductGetOkResponse(product)
		productOkResponse.WriteResponse(writer)
	}
}
