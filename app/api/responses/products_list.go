package responses

//go:generate easyjson

import (
	"net/http"

	"otus-microservice-architecture/app/models"
	db "otus-microservice-architecture/app/storage/db/sqlc"
)

//easyjson:json
type ProductsListOkResponse struct {
	Paginated
	Items []ProductsGetOkResponse `json:"items"`
}

func NewProductsListOkResponse(products []db.Product, paginated models.Paginated) ProductsListOkResponse {
	response := ProductsListOkResponse{
		Paginated: Paginated{
			Page:  paginated.GetPage(),
			Limit: paginated.GetLimit(),
			Total: paginated.GetTotal(),
			Pages: paginated.GetPages(),
		},
		Items: []ProductsGetOkResponse{},
	}
	for _, product := range products {
		response.Items = append(response.Items, NewProductsGetOkResponse(product))
	}

	return response
}

func (p *ProductsListOkResponse) WriteResponse(rw http.ResponseWriter) {
	payload, _ := p.MarshalJSON()

	WriteJsonResponse(rw, http.StatusOK, payload)
}
