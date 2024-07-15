package main

import (
	"net/http"

	"github.com/fernandomocrosky/DevBookGo/src/routes"
)

func main() {
	router := routes.GetRouter()

	http.ListenAndServe(":8080", router)
}
