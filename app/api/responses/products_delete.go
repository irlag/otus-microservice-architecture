package responses

import (
	"net/http"
)

type ProductsDeleteOkResponse struct{}

func NewProductsDeleteOkResponse() ProductsDeleteOkResponse {
	return ProductsDeleteOkResponse{}
}

func (p *ProductsDeleteOkResponse) WriteResponse(rw http.ResponseWriter) {
	WriteJsonResponse(rw, http.StatusNoContent, []byte{})
}
