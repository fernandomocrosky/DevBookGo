package routes

import (
	"net/http"

	"github.com/fernandomocrosky/DevBookGo/src/controllers"
)

var loginRoute = Route{
	URI:           "/login",
	Method:        http.MethodPost,
	HandleFunc:    controllers.Login,
	Authenticated: false,
}
