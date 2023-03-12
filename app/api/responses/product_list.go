package responses

//go:generate easyjson

import (
	"net/http"

	"otus-microservice-architecture/app/models"
	db "otus-microservice-architecture/app/storage/db/sqlc"
)

//easyjson:json
type ProductListOkResponse struct {
	Paginated
	Items []ProductGetOkResponse `json:"items"`
}

func NewProductListOkResponse(products []db.Product, paginated models.Paginated) ProductListOkResponse {
	response := ProductListOkResponse{
		Paginated: Paginated{
			Page:  paginated.GetPage(),
			Limit: paginated.GetLimit(),
			Total: paginated.GetTotal(),
			Pages: paginated.GetPages(),
		},
		Items: []ProductGetOkResponse{},
	}
	for _, product := range products {
		response.Items = append(response.Items, NewProductGetOkResponse(product))
	}

	return response
}

func (p *ProductListOkResponse) WriteResponse(rw http.ResponseWriter) {
	payload, _ := p.MarshalJSON()

	WriteJsonResponse(rw, http.StatusOK, payload)
}
