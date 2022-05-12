package routes

import (
	"habilitacion_backend/http/handlers"
)

func registrarRutasProducto() {
	router.HandleFunc("/producto/crear", handlers.CrearProducto).Methods("POST")
	router.HandleFunc("/producto/actualizar", handlers.ActualizarProducto).Methods("PUT")
	router.HandleFunc("/producto/eliminar", handlers.EliminarProducto).Methods("DELETE")
	router.HandleFunc("/producto/listar", handlers.ListarProductos).Methods("GET")
	router.HandleFunc("/producto/socio/listar", handlers.ListarProductosPorSocio).Methods("GET")
	router.HandleFunc("/producto/listar/categoria", handlers.ListarProductosPorCategoria).Methods("GET")
	router.HandleFunc("/producto/buscar", handlers.ObtenerProducto).Methods("GET")
	router.HandleFunc("/producto/categoria/agregar", handlers.AgregarCategoriaProducto).Methods("PATCH")
	router.HandleFunc("/producto/categoria/eliminar", handlers.EliminarCategoriaProducto).Methods("PATCH")
}
