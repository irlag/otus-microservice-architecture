package responses

import (
	"net/http"
)

type UserDeleteOkResponse struct{}

func NewUserDeleteOkResponse() UserDeleteOkResponse {
	return UserDeleteOkResponse{}
}

func (u *UserDeleteOkResponse) WriteResponse(rw http.ResponseWriter) {
	WriteJsonResponse(rw, http.StatusNoContent, []byte{})
}
