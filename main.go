package main

import (
	"log"
	"net/http"

	"github.com/eichiarakaki/go-bookstore/pkg/middleware"
	"github.com/eichiarakaki/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	socket := "localhost:9010"

	routes.RegisterAPIRoutes(r)
	routes.RegisterWebRoutes(r)

	// Crea un servidor de archivos que sirve archivos del directorio ./static.
	fs := http.FileServer(http.Dir("./static"))

	// Asocia todas las URLs que empiezan con /static con el manejador de archivos estáticos.
	// http.StripPrefix("/static/", fs) elimina el prefijo /static/ de la URL antes de pasarla al servidor de archivos.
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// Adding the Middleware
	r.Use(middleware.CheckAllowedPaths)

	log.Println("web server running on", socket)
	log.Fatal(http.ListenAndServe(socket, r))
}
