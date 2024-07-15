package controllers

import (
	"net/http"

	"github.com/fernandomocrosky/DevBookGo/src/responses"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")

	responses.JSON(w, struct {
		Hello string
	}{
		Hello: "Hello World!",
	})
}
