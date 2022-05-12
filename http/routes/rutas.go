package routes

import (
	"github.com/gorilla/mux"
)

var router *mux.Router

func CreateRouter() *mux.Router {
	router = mux.NewRouter()
	registrarRutasCategoria()
	registrarRutasProducto()

	return router
}
