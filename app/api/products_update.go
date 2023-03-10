package api

import (
	"database/sql"
	"net/http"

	"otus-microservice-architecture/app/api/parameters"
	"otus-microservice-architecture/app/api/responses"
	appErrors "otus-microservice-architecture/app/errors"
)

func (p *Products) Update() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		productsUpdateParams, err := parameters.NewProductsUpdateParamsFromRequest(request)
		if err != nil {
			responses.NewErrorResponse(http.StatusBadRequest, err).WriteErrorResponse(writer)

			return
		}

		_, err = p.processors.ProductsProcessor.Get(request.Context(), productsUpdateParams.ProductID)
		if err != nil {
			response := &responses.ErrorResponse{}
			switch err {
			case sql.ErrNoRows:
				response = responses.NewErrorResponse(http.StatusNotFound, appErrors.ProductNotFound)
			default:
				response = responses.NewErrorResponse(http.StatusInternalServerError, err)
			}
			response.WriteErrorResponse(writer)

			return
		}

		product, err := p.processors.ProductsProcessor.Update(request.Context(), productsUpdateParams)
		if err != nil {
			responses.NewErrorResponse(http.StatusInternalServerError, err).WriteErrorResponse(writer)

			return
		}

		productOkResponse := responses.NewProductsUpdateOkResponse(product)
		productOkResponse.WriteResponse(writer)
	}
}
