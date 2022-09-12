package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	TodoRoutes(r)
	UserRoutes(r)
	AuthRoutes(r)
	filmRoutes(r)
	categoryRoutes(r)
}
