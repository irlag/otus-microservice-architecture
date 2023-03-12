package product

import (
	"net/http"

	"otus-microservice-architecture/app/api/parameters"
	"otus-microservice-architecture/app/api/responses"
)

func (p *Product) Create() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		productsCreateParams, err := parameters.NewProductsCreateParamsFromRequest(request)
		if err != nil {
			responses.NewErrorResponse(http.StatusBadRequest, err).WriteErrorResponse(writer)

			return
		}

		product, err := p.processors.ProductProcessor.Create(request.Context(), *productsCreateParams)
		if err != nil {
			responses.NewErrorResponse(http.StatusInternalServerError, err).WriteErrorResponse(writer)

			return
		}

		productOkResponse := responses.NewProductCreateOkResponse(product)
		productOkResponse.WriteResponse(writer)
	}
}
