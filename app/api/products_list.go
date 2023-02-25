package api

import (
	"net/http"

	"otus-microservice-architecture/app/api/responses"
	"otus-microservice-architecture/app/models"
)

func (p *Products) List() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		paginated := models.NewPaginatedFromRequest(request)

		products, err := p.processors.ProductsProcessor.List(request.Context(), paginated.GetLimit(), paginated.GetOffset())
		if err != nil {
			responses.NewErrorResponse(http.StatusBadRequest, err).WriteErrorResponse(writer)

			return
		}
		total, err := p.processors.ProductsProcessor.ListCount(request.Context())
		if err != nil {
			responses.NewErrorResponse(http.StatusInternalServerError, err).WriteErrorResponse(writer)

			return
		}

		paginated.SetTotal(total)

		productListResponse := responses.NewProductsListOkResponse(products, paginated)
		productListResponse.WriteResponse(writer)
	}
}
