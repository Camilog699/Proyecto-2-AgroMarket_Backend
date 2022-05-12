package repositories

import (
	"context"
	"errors"
	"time"

	"habilitacion_backend/app"
	"habilitacion_backend/database/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const COLECCION_PRODUCTOS = "productos"

func CrearProducto(body models.Producto) (producto models.Producto, err error) {
	db := app.GetConnection()
	body.FechaCreacion = primitive.NewDateTimeFromTime(time.Now())
	result, err := db.Collection(COLECCION_PRODUCTOS).InsertOne(
		context.TODO(),
		body,
	)
	if err != nil {
		return
	}
	ID, _ := result.InsertedID.(primitive.ObjectID)
	err = db.Collection(COLECCION_PRODUCTOS).FindOne(
		context.TODO(),
		bson.M{"_id": ID},
	).Decode(&producto)
	return
}

func ActualizarProducto(ID string, body models.Producto) (producto models.Producto, err error) {
	db := app.GetConnection()
	objId, _ := primitive.ObjectIDFromHex(ID)
	_, err = db.Collection(COLECCION_PRODUCTOS).UpdateOne(
		context.TODO(),
		bson.M{"_id": objId},
		bson.M{
			"$set": bson.M{
				"nombre":              body.Nombre,
				"descripcion":         body.Descripcion,
				"precio":              body.Precio,
				"cantidad":            body.Cantidad,
				"idCategoriaProducto": body.IdCategoriaProducto,
				"fechaActualizacion":  time.Now(),
			},
		},
	)
	if err != nil {
		return
	}

	err = db.Collection(COLECCION_PRODUCTOS).FindOne(
		context.TODO(),
		bson.M{"_id": objId},
	).Decode(&producto)
	return
}

func EliminarProducto(ID string) (err error) {
	db := app.GetConnection()
	objId, _ := primitive.ObjectIDFromHex(ID)
	_, err = db.Collection(COLECCION_PRODUCTOS).DeleteOne(
		context.TODO(),
		bson.M{"_id": objId},
	)
	if err != nil {
		err = errors.New("no se pudo eliminar el producto")
		return
	}
	return
}

func ListarProductos() (productos []models.Producto, err error) {
	productos = make([]models.Producto, 0) // inicializar

	db := app.GetConnection()

	results, err := db.Collection(COLECCION_PRODUCTOS).Find(
		context.TODO(),
		bson.M{},
	)
	if err != nil {
		return
	}

	for results.Next(context.TODO()) {
		var producto models.Producto
		err = results.Decode(&producto)
		if err != nil {
			return
		}
		productos = append(productos, producto)
	}

	return
}

func ListarProductosPorSocio(idSocio string) (productos []models.Producto, err error) {
	productos = make([]models.Producto, 0) // inicializar

	db := app.GetConnection()

	results, err := db.Collection(COLECCION_PRODUCTOS).Find(
		context.TODO(),
		bson.M{
			"idUsuarioCreador": idSocio,
		},
	)
	if err != nil {
		return
	}

	for results.Next(context.TODO()) {
		var producto models.Producto
		err = results.Decode(&producto)
		if err != nil {
			return
		}
		productos = append(productos, producto)
	}

	return
}

func ListarProductosPorCategoria(idCategoria string) (productos []models.Producto, err error) {
	productos = make([]models.Producto, 0) // inicializar

	db := app.GetConnection()

	results, err := db.Collection(COLECCION_PRODUCTOS).Find(
		context.TODO(),
		bson.M{
			"idCategoriaProducto": idCategoria,
		},
	)
	if err != nil {
		return
	}

	for results.Next(context.TODO()) {
		var producto models.Producto
		err = results.Decode(&producto)
		if err != nil {
			return
		}
		productos = append(productos, producto)
	}

	return
}

func ObtenerProducto(ID string) (producto models.Producto, err error) {
	db := app.GetConnection()
	objId, _ := primitive.ObjectIDFromHex(ID)
	err = db.Collection(COLECCION_PRODUCTOS).FindOne(
		context.TODO(),
		bson.M{"_id": objId},
	).Decode(&producto)
	if err != nil {
		return
	}
	return
}

func AgregarCategoriaProducto(idPro, idNewCat string) (err error) {
	db := app.GetConnection()
	ObjId, _ := primitive.ObjectIDFromHex(idPro)

	_, err = db.Collection(COLECCION_PRODUCTOS).UpdateOne(
		context.TODO(),
		bson.M{"_id": ObjId},
		bson.M{
			"$set": bson.M{
				"idCategoriaProducto": idNewCat,
			},
		},
	)
	if err != nil {
		err = errors.New("no se pudo cambiar la categoria del producto")
		return
	}
	return
}

func EliminarCategoriaProducto(idPro string) (err error) {
	db := app.GetConnection()
	ObjId, _ := primitive.ObjectIDFromHex(idPro)

	_, err = db.Collection(COLECCION_PRODUCTOS).UpdateOne(
		context.TODO(),
		bson.M{"_id": ObjId},
		bson.M{
			"$set": bson.M{
				"idCategoriaProducto": "",
			},
		},
	)
	if err != nil {
		err = errors.New("no se pudo cambiar la categoria del producto")
		return
	}
	return
}
