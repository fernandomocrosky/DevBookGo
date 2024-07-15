package routes

import "github.com/gorilla/mux"

func GetRouter() *mux.Router {
	router := mux.NewRouter()

	return CreateRoutes(router)
}

func CreateRoutes(router *mux.Router) *mux.Router {

	bookRoutes := bookRoutes

	for _, route := range bookRoutes {
		router.HandleFunc(route.URI, route.HandleFunc).Methods(route.Method)
	}

	return router
}
