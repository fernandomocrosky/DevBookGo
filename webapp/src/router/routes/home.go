package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var homeRoute = Route{
	URI:           "/home",
	Method:        http.MethodGet,
	HandleFunc:    controllers.LoadHomePage,
	Authenticated: true,
}
