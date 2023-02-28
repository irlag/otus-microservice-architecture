package api

type Route struct {
	Name   string
	Method string
	Path   string
}

type Routes struct {
	Healthcheck Route
}

var routes = Routes{
	Healthcheck: Route{
		Name:   "healthcheck",
		Method: "GET",
		Path:   "/health",
	},
}
