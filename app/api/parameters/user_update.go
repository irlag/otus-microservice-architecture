package parameters

import (
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//go:generate easyjson

//easyjson:json
type UserUpdateParams struct {
	UserID    int64  `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

func NewUserUpdateParamsFromRequest(request *http.Request) (UserUpdateParams, error) {
	vars := mux.Vars(request)
	userId, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		return UserUpdateParams{}, err
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		return UserUpdateParams{}, err
	}

	userRequest := UserUpdateParams{}
	err = userRequest.UnmarshalJSON(body)
	if err != nil {
		return UserUpdateParams{}, err
	}

	userRequest.UserID = int64(userId)

	return userRequest, nil
}
