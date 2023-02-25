package api

type Route struct {
	Name   string
	Method string
	Path   string
}

type Routes struct {
	Healthcheck    Route
	ProductsList   Route
	ProductsDetail Route
	ProductsCreate Route
	ProductsDelete Route
	ProductsUpdate Route
}

var routes = Routes{
	Healthcheck: Route{
		Name:   "healthcheck",
		Method: "GET",
		Path:   "/health",
	},
	ProductsList: Route{
		Name:   "products_list",
		Method: "GET",
		Path:   "/products",
	},
	ProductsDetail: Route{
		Name:   "products_detail",
		Method: "GET",
		Path:   "/products/{product_id:[0-9]+}",
	},
	ProductsCreate: Route{
		Name:   "products_create",
		Method: "POST",
		Path:   "/products",
	},
	ProductsUpdate: Route{
		Name:   "products_update",
		Method: "PUT",
		Path:   "/products/{product_id:[0-9]+}",
	},
	ProductsDelete: Route{
		Name:   "products_delete",
		Method: "DELETE",
		Path:   "/products/{product_id:[0-9]+}",
	},
}
