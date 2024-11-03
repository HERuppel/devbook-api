package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI                   string
	Method                string
	Handler               func(http.ResponseWriter, *http.Request)
	requireAuthentication bool
}

func Configure(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, authRoutes...)

	for _, route := range routes {

		if route.requireAuthentication {
			r.HandleFunc(route.URI, middlewares.Logger(middlewares.Authenticate(route.Handler))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Handler)).Methods(route.Method)
		}
	}

	return r
}
