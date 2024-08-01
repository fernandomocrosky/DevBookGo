package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI           string
	Method        string
	HandleFunc    func(w http.ResponseWriter, r *http.Request)
	Authenticated bool
}

func Configure(router *mux.Router) *mux.Router {
	routes := loginRoutes
	routes = append(routes, usersRoutes...)
	routes = append(routes, homeRoute)

	for _, route := range routes {
		router.HandleFunc(route.URI, route.HandleFunc).Methods(route.Method)
	}

	fileServer := http.FileServer(http.Dir("./assets/"))

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
