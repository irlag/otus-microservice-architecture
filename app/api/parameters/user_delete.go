package parameters

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserDeleteParams struct {
	UserID int64
}

func NewUserDeleteParamsFromRequest(request *http.Request) (UserDeleteParams, error) {
	vars := mux.Vars(request)
	userId, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		return UserDeleteParams{}, err
	}

	return UserDeleteParams{
		UserID: int64(userId),
	}, nil
}
