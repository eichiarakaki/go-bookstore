package main

import (
	"log"
	"net/http"

	"github.com/eichiarakaki/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	socket := "localhost:9010"

	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(socket, r))
}
