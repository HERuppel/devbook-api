package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:                   "/users",
		Method:                http.MethodPost,
		Handler:               controllers.Create,
		requireAuthentication: false,
	},
	{
		URI:                   "/users",
		Method:                http.MethodGet,
		Handler:               controllers.Fetch,
		requireAuthentication: false,
	},
	{
		URI:                   "/users/{id}",
		Method:                http.MethodGet,
		Handler:               controllers.Get,
		requireAuthentication: false,
	},
	{
		URI:                   "/users/{id}",
		Method:                http.MethodPut,
		Handler:               controllers.Update,
		requireAuthentication: false,
	},
	{
		URI:                   "/users/{id}",
		Method:                http.MethodDelete,
		Handler:               controllers.Delete,
		requireAuthentication: false,
	},
}
