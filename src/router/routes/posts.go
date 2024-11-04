package routes

import (
	"api/src/controllers"
	"net/http"
)

var postRoutes = []Route{
	{
		URI:                   "/posts",
		Method:                http.MethodPost,
		Handler:               controllers.CreatePost,
		requireAuthentication: true,
	},
	{
		URI:                   "/posts",
		Method:                http.MethodGet,
		Handler:               controllers.FetchPosts,
		requireAuthentication: true,
	},
	{
		URI:                   "/posts/{id}",
		Method:                http.MethodGet,
		Handler:               controllers.GetPost,
		requireAuthentication: true,
	},
	{
		URI:                   "/posts",
		Method:                http.MethodPut,
		Handler:               controllers.UpdatePost,
		requireAuthentication: true,
	},
	{
		URI:                   "/posts",
		Method:                http.MethodDelete,
		Handler:               controllers.DeletePost,
		requireAuthentication: true,
	},
}
