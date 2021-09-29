package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Tipo das rotas
type Route struct {
	URI          string
	Method       string
	Function     func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}

// Coloca todas as rotas dentro do router
func Config(r *mux.Router) *mux.Router {
	routes := usersRoutes
	routes = append(routes, loginRoutes)
	routes = append(routes, publicationsRoutes...)

	for _, route := range routes {

		if route.AuthRequired {
			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Auth(route.Function)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}

	}

	return r
}
