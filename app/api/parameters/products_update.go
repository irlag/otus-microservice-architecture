package parameters

//go:generate easyjson

import (
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//easyjson:json
type ProductsUpdateParams struct {
	ProductID       int64  `json:"id"`
	Name            string `json:"name"`
	Article         string `json:"article"`
	Brand           string `json:"brand"`
	CountryOfOrigin string `json:"country_of_origin"`
	Description     string `json:"description"`
	Price           string `json:"price"`
	Version         string `json:"version"`
}

func NewProductsUpdateParamsFromRequest(request *http.Request) (ProductsUpdateParams, error) {
	vars := mux.Vars(request)
	productId, err := strconv.Atoi(vars["product_id"])
	if err != nil {
		return ProductsUpdateParams{}, err
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		return ProductsUpdateParams{}, err
	}

	productRequest := ProductsUpdateParams{}
	err = productRequest.UnmarshalJSON(body)
	if err != nil {
		return ProductsUpdateParams{}, err
	}

	productRequest.ProductID = int64(productId)

	return productRequest, nil
}
