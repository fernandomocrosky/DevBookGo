package routes

import (
	"net/http"

	"github.com/fernandomocrosky/DevBookGo/src/controllers"
)

var userRoutes = []Route{
	{
		URI:           "/users",
		Method:        http.MethodPost,
		HandleFunc:    controllers.CreateUser,
		Authenticated: false,
	},
	{
		URI:           "/users",
		Method:        http.MethodGet,
		HandleFunc:    controllers.GetAllUsers,
		Authenticated: true,
	},
	{
		URI:           "/users/{id}",
		Method:        http.MethodGet,
		HandleFunc:    controllers.GetUserById,
		Authenticated: true,
	},
	{
		URI:           "/users/{id}",
		Method:        http.MethodPut,
		HandleFunc:    controllers.UpdateUser,
		Authenticated: true,
	},
	{
		URI:           "/users/{id}",
		Method:        http.MethodDelete,
		HandleFunc:    controllers.DeleteUser,
		Authenticated: true,
	},
	{
		URI:           "/users/{id}/follow",
		Method:        http.MethodPost,
		HandleFunc:    controllers.FollowUser,
		Authenticated: true,
	},
	{
		URI:           "/users/{id}/unfollow",
		Method:        http.MethodPost,
		HandleFunc:    controllers.UnfollowUser,
		Authenticated: true,
	},
	{
		URI:           "/users/{id}/followers",
		Method:        http.MethodGet,
		HandleFunc:    controllers.GetFollowers,
		Authenticated: true,
	},
	{
		URI:           "/users/{id}/following",
		Method:        http.MethodGet,
		HandleFunc:    controllers.GetFollowing,
		Authenticated: true,
	},
	{
		URI:           "/users/{id}/update-password",
		Method:        http.MethodPost,
		HandleFunc:    controllers.UpdatePassword,
		Authenticated: true,
	},
}
