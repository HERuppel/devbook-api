package routes

import (
	"api/src/controllers"
	"net/http"
)

var authRoutes = []Route{
	{
		URI:                   "/login",
		Method:                http.MethodPost,
		Handler:               controllers.Login,
		requireAuthentication: false,
	},
}
