package parameters

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserGetParams struct {
	UserID int64
}

func NewUserGetParamsFromRequest(request *http.Request) (*UserGetParams, error) {
	vars := mux.Vars(request)
	userId, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		return nil, err
	}

	return &UserGetParams{
		UserID: int64(userId),
	}, nil
}
