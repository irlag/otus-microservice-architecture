package product

import (
	"github.com/gorilla/mux"

	"otus-microservice-architecture/app/api"
	"otus-microservice-architecture/app/processors"
)

type User struct {
	processors *processors.Processors
}

func NewUserApi(processors *processors.Processors) *User {
	return &User{
		processors: processors,
	}
}

func (u *User) HandleMethods(router *mux.Router) {
	router.HandleFunc(api.AppRoutes["user_create"].Path, u.Create()).
		Methods(api.AppRoutes["user_create"].Method).
		Name(api.AppRoutes["user_create"].Name)

	router.HandleFunc(api.AppRoutes["user_detail"].Path, u.Get()).
		Methods(api.AppRoutes["user_detail"].Method).
		Name(api.AppRoutes["user_detail"].Name)

	router.HandleFunc(api.AppRoutes["user_update"].Path, u.Update()).
		Methods(api.AppRoutes["user_update"].Method).
		Name(api.AppRoutes["user_update"].Name)

	router.HandleFunc(api.AppRoutes["user_delete"].Path, u.Delete()).
		Methods(api.AppRoutes["user_delete"].Method).
		Name(api.AppRoutes["user_delete"].Name)
}
