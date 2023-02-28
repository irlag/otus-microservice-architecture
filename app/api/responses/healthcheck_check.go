package responses

//go:generate easyjson

import (
	"net/http"
)

type Status string

const StatusOK Status = "OK"

//easyjson:json
type HealthcheckOkResponse struct {
	Status `json:"status"`
}

func NewHealthcheckOkResponse(status Status) HealthcheckOkResponse {
	return HealthcheckOkResponse{
		Status: status,
	}
}

func (h *HealthcheckOkResponse) WriteResponse(rw http.ResponseWriter) {
	result, _ := h.MarshalJSON()

	WriteJsonResponse(rw, http.StatusOK, result)
}
