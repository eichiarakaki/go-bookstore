package routes

import (
	"github.com/eichiarakaki/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterAPIRoutes = func(router *mux.Router) {
	apiRouter := router.PathPrefix("/api").Subrouter()

	apiRouter.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	apiRouter.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	apiRouter.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	apiRouter.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	apiRouter.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
	// APIs URL: /api/book/{id} - /api/book
}
