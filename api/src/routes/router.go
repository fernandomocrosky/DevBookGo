package routes

import (
	"net/http"

	"github.com/fernandomocrosky/DevBookGo/src/middlewares"
	"github.com/gorilla/mux"
)

type Route struct {
	URI           string
	Method        string
	HandleFunc    func(w http.ResponseWriter, r *http.Request)
	Authenticated bool
}

func GetRouter() *mux.Router {
	router := mux.NewRouter()

	return CreateRoutes(router)
}

func CreateRoutes(router *mux.Router) *mux.Router {

	routes := userRoutes
	routes = append(routes, loginRoute)
	routes = append(routes, postsRoutes...)

	for _, route := range routes {
		if route.Authenticated {
			router.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authenticate(route.HandleFunc)),
			).Methods(route.Method)
		} else {
			router.HandleFunc(route.URI, middlewares.Logger(route.HandleFunc)).Methods(route.Method)
		}
	}

	return router
}
