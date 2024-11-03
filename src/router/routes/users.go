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
		requireAuthentication: true,
	},
	{
		URI:                   "/users/{id}",
		Method:                http.MethodGet,
		Handler:               controllers.Get,
		requireAuthentication: true,
	},
	{
		URI:                   "/users/{id}",
		Method:                http.MethodPut,
		Handler:               controllers.Update,
		requireAuthentication: true,
	},
	{
		URI:                   "/users/{id}",
		Method:                http.MethodDelete,
		Handler:               controllers.Delete,
		requireAuthentication: true,
	},
	{
		URI:                   "/users/{id}/follow",
		Method:                http.MethodPost,
		Handler:               controllers.FollowUser,
		requireAuthentication: true,
	},
	{
		URI:                   "/users/{id}/unfollow",
		Method:                http.MethodPost,
		Handler:               controllers.UnfollowUser,
		requireAuthentication: true,
	},
	{
		URI:                   "/users/{id}/followers",
		Method:                http.MethodGet,
		Handler:               controllers.FetchFollowers,
		requireAuthentication: true,
	},
	{
		URI:                   "/users/{id}/following",
		Method:                http.MethodGet,
		Handler:               controllers.FetchFollowing,
		requireAuthentication: true,
	},
}
