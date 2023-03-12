package api

import "errors"

var (
	RouteNotExistError = errors.New("route not found")
)

type Route struct {
	Name   string
	Method string
	Path   string
	Secure bool
}

type Routes map[string]Route

func (r *Routes) Find(name string) (Route, error) {
	if appRoute, ok := AppRoutes[name]; ok {
		return appRoute, nil
	}
	return Route{}, RouteNotExistError
}

var AppRoutes = Routes{
	"healthcheck": Route{
		Name:   "healthcheck",
		Method: "GET",
		Path:   "/health",
		Secure: false,
	},

	"auth_login": Route{
		Name:   "auth_login",
		Method: "POST",
		Path:   "/auth/login",
		Secure: false,
	},

	"product_create": Route{
		Name:   "product_create",
		Method: "POST",
		Path:   "/product",
		Secure: false,
	},
	"product_list": Route{
		Name:   "product_list",
		Method: "GET",
		Path:   "/product",
		Secure: false,
	},
	"product_detail": Route{
		Name:   "product_detail",
		Method: "GET",
		Path:   "/product/{product_id:[0-9]+}",
		Secure: false,
	},
	"product_update": Route{
		Name:   "product_update",
		Method: "PUT",
		Path:   "/product/{product_id:[0-9]+}",
		Secure: false,
	},
	"product_delete": Route{
		Name:   "product_delete",
		Method: "DELETE",
		Path:   "/product/{product_id:[0-9]+}",
		Secure: false,
	},

	"user_create": Route{
		Name:   "user_create",
		Method: "POST",
		Path:   "/user",
		Secure: false,
	},
	"user_detail": Route{
		Name:   "user_detail",
		Method: "GET",
		Path:   "/user/{user_id:[0-9]+}",
		Secure: true,
	},
	"user_update": Route{
		Name:   "user_update",
		Method: "PUT",
		Path:   "/user/{user_id:[0-9]+}",
		Secure: true,
	},
	"user_delete": Route{
		Name:   "user_delete",
		Method: "DELETE",
		Path:   "/user/{user_id:[0-9]+}",
		Secure: true,
	},
}
