package responses

//go:generate easyjson

import (
	"net/http"

	db "otus-microservice-architecture/app/storage/db/sqlc"
)

//easyjson:json
type UserCreateOkResponse struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	CreatedAt string `json:"created_at"`
}

func NewUserCreateOkResponse(userCreateRow db.CreateUserRow) UserCreateOkResponse {
	return UserCreateOkResponse{
		ID:        userCreateRow.ID,
		FirstName: userCreateRow.FirstName,
		LastName: func() string {
			if userCreateRow.LastName.Valid {
				return userCreateRow.LastName.String
			}

			return ""
		}(),
		Email:     userCreateRow.Email,
		CreatedAt: userCreateRow.CreatedAt.String(),
	}
}

func (p *UserCreateOkResponse) WriteResponse(rw http.ResponseWriter) {
	payload, _ := p.MarshalJSON()

	WriteJsonResponse(rw, http.StatusCreated, payload)
}
