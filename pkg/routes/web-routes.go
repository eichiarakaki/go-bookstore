package routes

import (
	"github.com/eichiarakaki/go-bookstore/pkg/pages"
	"github.com/gorilla/mux"
)

var RegisterWebRoutes = func(router *mux.Router) {
	router.HandleFunc("/", pages.Homepage).Methods("GET")
	router.HandleFunc("/book", pages.Book).Methods("GET")
}
