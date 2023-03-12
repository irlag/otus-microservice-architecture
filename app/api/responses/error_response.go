package responses

//go:generate easyjson

import (
	"net/http"
)

//easyjson:json
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewErrorResponse(code int, err error) *ErrorResponse {
	return &ErrorResponse{
		Code:    code,
		Message: err.Error(),
	}
}

func (e *ErrorResponse) WriteErrorResponse(rw http.ResponseWriter) {
	result, _ := e.MarshalJSON()

	WriteJsonResponse(rw, e.Code, result)
}
