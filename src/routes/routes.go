package routes

import (
	"net/http"

	"github.com/fernandomocrosky/DevBookGo/src/controllers"
)

type Routes struct {
	URI           string
	Method        string
	HandleFunc    func(w http.ResponseWriter, r *http.Request)
	Authenticated bool
}

var bookRoutes = []Routes{
	{
		URI:           "/books",
		Method:        http.MethodGet,
		HandleFunc:    controllers.GetBooks,
		Authenticated: false,
	},
}
