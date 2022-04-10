package routes

import (
	"habilitacion_backend/http/handlers"
)

/*
creación de los diferentes endopints para la categoria de productos
*/
func registrarRutasCategoria() {
	router.HandleFunc("/categoria/crear", handlers.CrearCategoria).Methods("POST")
	router.HandleFunc("/categoria/actualizar", handlers.ActualizarCategoria).Methods("PUT")
	router.HandleFunc("/categoria/eliminar", handlers.EliminarCategoria).Methods("DELETE")
	router.HandleFunc("/categoria/listar", handlers.ListarCategorias).Methods("GET")
	router.HandleFunc("/categoria/buscar", handlers.ObtenerCategoria).Methods("GET")
}
