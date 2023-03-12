package product

import (
	"github.com/gorilla/mux"

	"otus-microservice-architecture/app/api"
	"otus-microservice-architecture/app/processors"
)

type Auth struct {
	processors *processors.Processors
}

func NewAuthApi(processors *processors.Processors) *Auth {
	return &Auth{
		processors: processors,
	}
}

func (a *Auth) HandleMethods(router *mux.Router) {
	router.HandleFunc(api.AppRoutes["auth_login"].Path, a.Login()).
		Methods(api.AppRoutes["auth_login"].Method).
		Name(api.AppRoutes["auth_login"].Name)

}
