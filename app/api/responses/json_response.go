package responses

import (
	"net/http"
)

func WriteJsonResponse(rw http.ResponseWriter, code int, data []byte) {
	rw.WriteHeader(code)

	rw.Write(data)
}
