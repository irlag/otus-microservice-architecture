package product

import (
	"github.com/gorilla/mux"

	"otus-microservice-architecture/app/api"
	"otus-microservice-architecture/app/processors"
)

type Product struct {
	processors *processors.Processors
}

func NewProductApi(processors *processors.Processors) *Product {
	return &Product{
		processors: processors,
	}
}

func (p *Product) HandleMethods(router *mux.Router) {
	router.HandleFunc(api.AppRoutes["product_list"].Path, p.List()).
		Methods(api.AppRoutes["product_list"].Method).
		Name(api.AppRoutes["product_list"].Name)

	router.HandleFunc(api.AppRoutes["product_detail"].Path, p.Get()).
		Methods(api.AppRoutes["product_detail"].Method).
		Name(api.AppRoutes["product_detail"].Name)

	router.HandleFunc(api.AppRoutes["product_create"].Path, p.Create()).
		Methods(api.AppRoutes["product_create"].Method).
		Name(api.AppRoutes["product_create"].Name)

	router.HandleFunc(api.AppRoutes["product_delete"].Path, p.Delete()).
		Methods(api.AppRoutes["product_delete"].Method).
		Name(api.AppRoutes["product_delete"].Name)

	router.HandleFunc(api.AppRoutes["product_update"].Path, p.Update()).
		Methods(api.AppRoutes["product_update"].Method).
		Name(api.AppRoutes["product_update"].Name)
}
