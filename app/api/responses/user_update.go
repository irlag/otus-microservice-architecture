package responses

//go:generate easyjson

import (
	"net/http"

	db "otus-microservice-architecture/app/storage/db/sqlc"
)

//easyjson:json
type UserUpdateOkResponse struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewUserUpdateOkResponse(user db.User) UserUpdateOkResponse {
	return UserUpdateOkResponse{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName: func() string {
			if user.LastName.Valid {
				return user.LastName.String
			}

			return ""
		}(),
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.Time.String(),
	}
}

func (u *UserUpdateOkResponse) WriteResponse(rw http.ResponseWriter) {
	payload, _ := u.MarshalJSON()

	WriteJsonResponse(rw, http.StatusOK, payload)
}
