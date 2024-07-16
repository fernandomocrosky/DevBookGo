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
		Authenticated: false,
	},
	{
		URI:           "/users/{id}",
		Method:        http.MethodGet,
		HandleFunc:    controllers.GetUserById,
		Authenticated: false,
	},
	{
		URI:           "/users/{id}",
		Method:        http.MethodPut,
		HandleFunc:    controllers.UpdateUser,
		Authenticated: false,
	},
	{
		URI:           "/users/{id}",
		Method:        http.MethodDelete,
		HandleFunc:    controllers.DeleteUser,
		Authenticated: false,
	},
}
