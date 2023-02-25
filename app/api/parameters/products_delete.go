package parameters

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductsDeleteParams struct {
	ProductID int64
}

func NewProductsDeleteParamsFromRequest(request *http.Request) (ProductsDeleteParams, error) {
	vars := mux.Vars(request)
	productId, err := strconv.Atoi(vars["product_id"])
	if err != nil {
		return ProductsDeleteParams{}, err
	}

	return ProductsDeleteParams{
		ProductID: int64(productId),
	}, nil
}
