package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var loginRoutes = []Route{
	{
		URI:           "/",
		Method:        http.MethodGet,
		HandleFunc:    controllers.LoadLoginPage,
		Authenticated: false,
	},
	{
		URI:           "/login",
		Method:        http.MethodGet,
		HandleFunc:    controllers.LoadLoginPage,
		Authenticated: false,
	},
	{
		URI:           "/login",
		Method:        http.MethodPost,
		HandleFunc:    controllers.Login,
		Authenticated: false,
	},
}
