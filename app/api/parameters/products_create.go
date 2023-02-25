package parameters

//go:generate easyjson

import (
	"io"
	"net/http"
)

//easyjson:json
type ProductsCreateParams struct {
	Name            string `json:"name"`
	Article         string `json:"article"`
	Brand           string `json:"brand"`
	CountryOfOrigin string `json:"country_of_origin"`
	Description     string `json:"description"`
	Price           string `json:"price"`
}

func NewProductsCreateParamsFromRequest(request *http.Request) (*ProductsCreateParams, error) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}

	productRequest := ProductsCreateParams{}
	err = productRequest.UnmarshalJSON(body)
	if err != nil {
		return nil, err
	}

	return &productRequest, nil
}
