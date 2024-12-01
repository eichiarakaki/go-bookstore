package middleware

import (
	"fmt"
	"net/http"
)

// es un mapa que contiene las rutas permitidas. Solo las rutas listadas aquí estarán accesibles
var allowedPaths = map[string]bool{
	"/":     true,
	"/book": true,
	"/api/": false, // User will have NO access to backend API (uncessesary to put the forbidden paths)
}

// checkAllowedPaths es un middleware que verifica si la ruta solicitada está en la lista de rutas permitidas. Si no lo está, responde con un error 403 Forbidden
func CheckAllowedPaths(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		if allowedPaths[r.URL.Path] {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}
