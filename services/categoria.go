package services

import (
	"habilitacion_backend/database/models"
	"habilitacion_backend/database/repositories"
)

func CrearCategoria(body models.Categoria) (categoria models.Categoria, err error) {
	categoria, err = repositories.CrearCategoria(body)
	return
}

func ActualizarCategoria(body models.Categoria, idCategoria string) (categoria models.Categoria, err error) {
	categoria, err = repositories.ActualizarCategoria(idCategoria, body)
	return
}

func EliminarCategoria(idCategoria string) (err error) {
	err = repositories.EliminarCategoria(idCategoria)
	return
}

func ListarCategorias() (categorias []models.Categoria, err error) {
	categorias, err = repositories.ListarCategorias()
	return
}

func ObtenerCategoria(ID string) (categoria models.Categoria, err error) {
	categoria, err = repositories.ObtenerCategoria(ID)
	return
}
