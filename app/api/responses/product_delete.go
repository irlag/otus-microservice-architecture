package responses

import (
	"net/http"
)

type ProductDeleteOkResponse struct{}

func NewProductDeleteOkResponse() ProductDeleteOkResponse {
	return ProductDeleteOkResponse{}
}

func (p *ProductDeleteOkResponse) WriteResponse(rw http.ResponseWriter) {
	WriteJsonResponse(rw, http.StatusNoContent, []byte{})
}
