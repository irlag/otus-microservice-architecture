package api

import (
	"github.com/gorilla/mux"

	"otus-microservice-architecture/app/processors"
)

type Products struct {
	processors *processors.Processors
}

func NewProductsApi(processors *processors.Processors) *Products {
	return &Products{
		processors: processors,
	}
}

func (p *Products) HandleMethods(router *mux.Router) {
	router.HandleFunc(routes.ProductsList.Path, p.List()).Methods(routes.ProductsList.Method)
	router.HandleFunc(routes.ProductsDetail.Path, p.Get()).Methods(routes.ProductsDetail.Method)
	router.HandleFunc(routes.ProductsCreate.Path, p.Create()).Methods(routes.ProductsCreate.Method)
	router.HandleFunc(routes.ProductsDelete.Path, p.Delete()).Methods(routes.ProductsDelete.Method)
	router.HandleFunc(routes.ProductsUpdate.Path, p.Update()).Methods(routes.ProductsUpdate.Method)
}
