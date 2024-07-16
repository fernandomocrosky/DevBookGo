package main

import (
	"fmt"
	"net/http"

	"github.com/fernandomocrosky/DevBookGo/src/config"
	"github.com/fernandomocrosky/DevBookGo/src/routes"
)

func main() {
	router := routes.GetRouter()

	config.Load()

	fmt.Printf("Server Listening on port %d...\n", config.Port)

	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router)
}
