package services

import (
	"habilitacion_backend/database/models"
	"habilitacion_backend/database/repositories"
)

func CrearProducto(body models.Producto) (producto models.Producto, err error) {
	producto, err = repositories.CrearProducto(body)
	return
}

func ActualizarProducto(body models.Producto, idProducto string) (producto models.Producto, err error) {
	producto, err = repositories.ActualizarProducto(idProducto, body)
	return
}

func EliminarProducto(idProducto string) (err error) {
	err = repositories.EliminarProducto(idProducto)
	return
}

func ListarProductos() (productos []models.Producto, err error) {
	productos, err = repositories.ListarProductos()
	return
}

func ListarProductosPorSocio(idSocio string) (productos []models.Producto, err error) {
	productos, err = repositories.ListarProductosPorSocio(idSocio)
	return
}

func ListarProductosPorCategoria(idCategoria string) (productos []models.Producto, err error) {
	productos, err = repositories.ListarProductosPorCategoria(idCategoria)
	return
}

func ObtenerProducto(ID string) (producto models.Producto, err error) {
	producto, err = repositories.ObtenerProducto(ID)
	return
}

func AgregarCategoriaProducto(idPro, idNewCat string) (err error) {
	err = repositories.AgregarCategoriaProducto(idPro, idNewCat)
	return
}

func EliminarCategoriaProducto(idPro string) (err error) {
	err = repositories.EliminarCategoriaProducto(idPro)
	return
}
