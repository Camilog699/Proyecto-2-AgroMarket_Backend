package handlers

import (
	"net/http"

	"habilitacion_backend/database/models"
	"habilitacion_backend/services"
	"habilitacion_backend/utils/io/request"
	"habilitacion_backend/utils/io/response"
)

func CrearProducto(w http.ResponseWriter, r *http.Request) {
	var body models.Producto
	err := request.Json(r, &body)
	if err != nil {
		response.Error("No se pudo crear el producto", http.StatusBadRequest, w)
		return
	}

	body.IdUsuarioCreador = w.Header().Get("User-Id")

	producto, err := services.CrearProducto(body)
	if err != nil {
		response.Error(err.Error(), http.StatusInternalServerError, w)
		return
	}
	response.Json(producto, http.StatusOK, w)
}

func ActualizarProducto(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()["id"]
	if len(keys) <= 0 {
		response.Error("parametro id producto es requerido", http.StatusBadRequest, w)
		return
	}

	var body models.Producto
	err := request.Json(r, &body)
	if err != nil {
		response.Error("No se pudo actualizar el producto", http.StatusBadRequest, w)
		return
	}
	body.IdUsuarioModificador = w.Header().Get("User-Id")
	producto, err := services.ActualizarProducto(body, keys[0])
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.Error("producto no encontrado", http.StatusNotFound, w)
			return
		}
		response.Error(err.Error(), http.StatusInternalServerError, w)
		return
	}
	response.Json(producto, http.StatusOK, w)
}

func EliminarProducto(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()["id"]
	if len(keys) <= 0 {
		response.Error("parametro id de producto es requerido", http.StatusBadRequest, w)
		return
	}

	err := services.EliminarProducto(keys[0])
	if err != nil {
		response.Error(err.Error(), http.StatusInternalServerError, w)
		return
	}

	response.Json("OK", http.StatusOK, w)
}

func ListarProductos(w http.ResponseWriter, r *http.Request) {
	productos, err := services.ListarProductos()
	if err != nil {
		response.Error("internal error", http.StatusInternalServerError, w)
		return
	}
	response.Json(productos, http.StatusOK, w)
}

func ListarProductosPorSocio(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()["idSocio"]
	productos, err := services.ListarProductosPorSocio(keys[0])
	if err != nil {
		response.Error("internal error", http.StatusInternalServerError, w)
		return
	}
	response.Json(productos, http.StatusOK, w)
}

func ListarProductosPorCategoria(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()["idCategoria"]
	productos, err := services.ListarProductosPorCategoria(keys[0])
	if err != nil {
		response.Error("internal error", http.StatusInternalServerError, w)
		return
	}
	response.Json(productos, http.StatusOK, w)
}

func ObtenerProducto(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()["id"]
	if len(keys) <= 0 {
		response.Error("parametro id de producto es requerido", http.StatusBadRequest, w)
		return
	}
	producto, err := services.ObtenerProducto(keys[0])
	if err != nil {
		response.Error(err.Error(), http.StatusInternalServerError, w)
		return
	}

	response.Json(producto, http.StatusOK, w)
}

func AgregarCategoriaProducto(w http.ResponseWriter, r *http.Request) {
	idPro := r.URL.Query()["producto"]
	idNewCat := r.URL.Query()["categoriaNueva"]
	if len(idPro) <= 0 || len(idNewCat) <= 0 {
		response.Error("parametro producto o categoriaNueva o categoriaAnterior es requerido", http.StatusBadRequest, w)
		return
	}
	// fmt.Println("en el back", idPro, idTop)

	err := services.AgregarCategoriaProducto(idPro[0], idNewCat[0])
	if err != nil {
		response.Error(err.Error(), http.StatusInternalServerError, w)
		return
	}

	response.Json("OK", http.StatusOK, w)
}

func EliminarCategoriaProducto(w http.ResponseWriter, r *http.Request) {
	idPro := r.URL.Query()["producto"]
	if len(idPro) <= 0 {
		response.Error("parametro producto o categoriaNueva o categoriaAnterior es requerido", http.StatusBadRequest, w)
		return
	}
	// fmt.Println("en el back", idPro, idTop)

	err := services.EliminarCategoriaProducto(idPro[0])
	if err != nil {
		response.Error(err.Error(), http.StatusInternalServerError, w)
		return
	}
	response.Json("OK", http.StatusOK, w)
}
