package services

import (
	"habilitacion_backend/database/models"
	"habilitacion_backend/database/repositories"
)

/*
Servicio que permite conectar la data obtenida de los endpoints con la base de datos
para crear una nueva categoria
*/
func CrearCategoria(body models.Categoria) (categoria models.Categoria, err error) {
	categoria, err = repositories.CrearCategoria(body)
	return
}

/*
Servicio que permite conectar la data obtenida de los endpoints con la base de datos
para actualizar una categoria
*/
func ActualizarCategoria(body models.Categoria, idCategoria string) (categoria models.Categoria, err error) {
	categoria, err = repositories.ActualizarCategoria(idCategoria, body)
	return
}

/*
Servicio que permite conectar la data obtenida de los endpoints con la base de datos
para eliminar una categoria
*/
func EliminarCategoria(idCategoria string) (err error) {
	err = repositories.EliminarCategoria(idCategoria)
	return
}

/*
Servicio que permite conectar la data obtenida de los endpoints con la base de datos
para obtener todas las categorias
*/
func ListarCategorias() (categorias []models.Categoria, err error) {
	categorias, err = repositories.ListarCategorias()
	return
}

/*
Servicio que permite conectar la data obtenida de los endpoints con la base de datos
para obtener una categoria
*/
func ObtenerCategoria(ID string) (categoria models.Categoria, err error) {
	categoria, err = repositories.ObtenerCategoria(ID)
	return
}
