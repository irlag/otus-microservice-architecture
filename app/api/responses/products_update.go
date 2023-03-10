package responses

//go:generate easyjson

import (
	"net/http"

	db "otus-microservice-architecture/app/storage/db/sqlc"
)

//easyjson:json
type ProductsUpdateOkResponse struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	Article         string `json:"article"`
	Brand           string `json:"brand"`
	CountryOfOrigin string `json:"country_of_origin"`
	Description     string `json:"description"`
	Price           string `json:"price"`
	Rating          string `json:"rating"`
}

func NewProductsUpdateOkResponse(product db.Product) ProductsUpdateOkResponse {
	return ProductsUpdateOkResponse{
		ID:              product.ID,
		Name:            product.Name,
		Article:         product.Article,
		Brand:           product.Brand,
		CountryOfOrigin: product.CountryOfOrigin,
		Description:     product.Description.String,
		Price:           product.Price,
		Rating:          product.Rating,
	}
}

func (p *ProductsUpdateOkResponse) WriteResponse(rw http.ResponseWriter) {
	payload, _ := p.MarshalJSON()

	WriteJsonResponse(rw, http.StatusOK, payload)
}
