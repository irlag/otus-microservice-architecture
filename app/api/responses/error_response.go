package responses

//go:generate easyjson

import (
	"net/http"
)

//easyjson:json
type ErrorResponse struct {
	Code    int
	Message string
}

func NewErrorResponse(code int, message string) ErrorResponse {
	return ErrorResponse{
		Code:    code,
		Message: message,
	}
}

func (e *ErrorResponse) WriteErrorResponse(rw http.ResponseWriter) {
	result, _ := e.MarshalJSON()
	WriteJsonResponse(rw, e.Code, result)
}
