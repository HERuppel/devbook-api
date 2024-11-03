package routes

import (
	"api/src/controllers"
	"net/http"
)

var authRoutes = []Route{
	{
		URI:                   "/auth/login",
		Method:                http.MethodPost,
		Handler:               controllers.Login,
		requireAuthentication: false,
	},
	{
		URI:                   "/auth/{id}/updatePassword",
		Method:                http.MethodPost,
		Handler:               controllers.UpdatePassword,
		requireAuthentication: true,
	},
}
