package product

import (
	"database/sql"
	"net/http"

	"otus-microservice-architecture/app/api/parameters"
	"otus-microservice-architecture/app/api/responses"
	appErrors "otus-microservice-architecture/app/errors"
)

func (p *Product) Update() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		productsUpdateParams, err := parameters.NewProductsUpdateParamsFromRequest(request)
		if err != nil {
			responses.NewErrorResponse(http.StatusBadRequest, err).WriteErrorResponse(writer)

			return
		}

		_, err = p.processors.ProductProcessor.Get(request.Context(), productsUpdateParams.ProductID)
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

		product, err := p.processors.ProductProcessor.Update(request.Context(), productsUpdateParams)
		if err != nil {
			responses.NewErrorResponse(http.StatusInternalServerError, err).WriteErrorResponse(writer)

			return
		}

		productOkResponse := responses.NewProductUpdateOkResponse(product)
		productOkResponse.WriteResponse(writer)
	}
}
