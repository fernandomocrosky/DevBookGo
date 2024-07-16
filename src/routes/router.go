package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Routes struct {
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

	userRoutes := userRoutes

	for _, route := range userRoutes {
		router.HandleFunc(route.URI, route.HandleFunc).Methods(route.Method)
	}

	return router
}
