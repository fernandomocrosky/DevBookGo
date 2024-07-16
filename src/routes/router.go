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

func GetRouter() *mux.Router {
	router := mux.NewRouter()

	return CreateRoutes(router)
}

func CreateRoutes(router *mux.Router) *mux.Router {

	routes := userRoutes
	routes = append(routes, loginRoute)

	for _, route := range routes {
		router.HandleFunc(route.URI, route.HandleFunc).Methods(route.Method)
	}

	return router
}
