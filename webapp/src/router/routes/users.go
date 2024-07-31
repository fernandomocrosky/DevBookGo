package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var usersRoutes = []Route{
	{
		URI:           "/create-user",
		Method:        http.MethodGet,
		HandleFunc:    controllers.LoadUserRegisterPage,
		Authenticated: false,
	},
	{
		URI:           "/users",
		Method:        http.MethodPost,
		HandleFunc:    controllers.CreateUser,
		Authenticated: false,
	},
}
