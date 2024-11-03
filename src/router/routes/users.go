package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:                   "/users",
		Method:                http.MethodPost,
		Function:              controllers.Create,
		requireAuthentication: false,
	},
	{
		URI:                   "/users",
		Method:                http.MethodGet,
		Function:              controllers.Fetch,
		requireAuthentication: false,
	},
	{
		URI:                   "/users/{id}",
		Method:                http.MethodGet,
		Function:              controllers.Get,
		requireAuthentication: false,
	},
	{
		URI:                   "/users/{id}",
		Method:                http.MethodPut,
		Function:              controllers.Update,
		requireAuthentication: false,
	},
	{
		URI:                   "/users/{id}",
		Method:                http.MethodDelete,
		Function:              controllers.Delete,
		requireAuthentication: false,
	},
}
