package parameters

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductsGetParams struct {
	ProductID int64
}

func NewProductsGetParamsFromRequest(request *http.Request) (*ProductsGetParams, error) {
	vars := mux.Vars(request)
	productId, err := strconv.Atoi(vars["product_id"])
	if err != nil {
		return nil, err
	}

	return &ProductsGetParams{
		ProductID: int64(productId),
	}, nil
}
