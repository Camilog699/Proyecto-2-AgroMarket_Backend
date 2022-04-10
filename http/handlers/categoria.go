package handlers

import (
	"net/http"

	"habilitacion_backend/database/models"

	"habilitacion_backend/services"
	"habilitacion_backend/utils/io/request"
	"habilitacion_backend/utils/io/response"
)

/*
funcion creada para obtener la informacion generada en el endpoint
para crear una categoria
*/
func CrearCategoria(w http.ResponseWriter, r *http.Request) {
	var body models.Categoria
	err := request.Json(r, &body)
	if err != nil {
		response.Error("No se pudo crear la categoria", http.StatusBadRequest, w)
		return
	}
	categoria, err := services.CrearCategoria(body)
	if err != nil {
		response.Error(err.Error(), http.StatusInternalServerError, w)
		return
	}
	response.Json(categoria, http.StatusOK, w)
}

/*
funcion creada para obtener la informacion generada en el endpoint
para actualizar una categoria
*/
func ActualizarCategoria(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()["id"]
	if len(keys) <= 0 {
		response.Error("parametro id categoria es requerido", http.StatusBadRequest, w)
		return
	}

	var body models.Categoria
	err := request.Json(r, &body)
	if err != nil {
		response.Error("No se pudo actualizar la categoria", http.StatusBadRequest, w)
		return
	}
	categoria, err := services.ActualizarCategoria(body, keys[0])
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.Error("Categoria no encontrada", http.StatusNotFound, w)
			return
		}
		response.Error(err.Error(), http.StatusInternalServerError, w)
		return
	}
	response.Json(categoria, http.StatusOK, w)
}

/*
funcion creada para obtener la informacion generada en el endpoint
para eliminar una categoria
*/
func EliminarCategoria(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()["id"]
	if len(keys) <= 0 {
		response.Error("parametro id de categoria es requerido", http.StatusBadRequest, w)
		return
	}
	err := services.EliminarCategoria(keys[0])
	if err != nil {
		response.Error(err.Error(), http.StatusInternalServerError, w)
		return
	}
	response.Json("OK", http.StatusOK, w)
}

/*
funcion creada para obtener la informacion generada en el endpoint
para todas las categorias
*/
func ListarCategorias(w http.ResponseWriter, r *http.Request) {
	categorias, err := services.ListarCategorias()
	if err != nil {
		response.Error("internal error", http.StatusInternalServerError, w)
		return
	}
	response.Json(categorias, http.StatusOK, w)
}

/*
funcion creada para obtener la informacion generada en el endpoint
para obtener una categoria
*/
func ObtenerCategoria(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()["id"]
	if len(keys) <= 0 {
		response.Error("parametro id de topping es requerido", http.StatusBadRequest, w)
		return
	}
	topping, err := services.ObtenerCategoria(keys[0])
	if err != nil {
		response.Error(err.Error(), http.StatusInternalServerError, w)
		return
	}

	response.Json(topping, http.StatusOK, w)
}
