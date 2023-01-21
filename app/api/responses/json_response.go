package responses

import (
	"net/http"
)

func WriteJsonResponse(rw http.ResponseWriter, code int, data []byte) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")

	rw.WriteHeader(code)

	rw.Write(data)
}
