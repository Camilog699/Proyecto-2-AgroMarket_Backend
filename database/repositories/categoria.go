package repositories

import (
	"context"
	"time"

	"habilitacion_backend/app"
	"habilitacion_backend/database/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const COLECCION_CATEGORIAS = "categoria"

func CrearCategoria(body models.Categoria) (categoria models.Categoria, err error) {
	db := app.GetConnection()
	body.FechaCreacion = primitive.NewDateTimeFromTime(time.Now())
	result, err := db.Collection(COLECCION_CATEGORIAS).InsertOne(
		context.TODO(),
		body,
	)
	if err != nil {
		return
	}
	ID, _ := result.InsertedID.(primitive.ObjectID)
	err = db.Collection(COLECCION_CATEGORIAS).FindOne(
		context.TODO(),
		bson.M{"_id": ID},
	).Decode(&categoria)
	return
}

func ActualizarCategoria(ID string, body models.Categoria) (categoria models.Categoria, err error) {
	db := app.GetConnection()
	objId, _ := primitive.ObjectIDFromHex(ID)
	_, err = db.Collection(COLECCION_CATEGORIAS).UpdateOne(
		context.TODO(),
		bson.M{"_id": objId},
		bson.M{
			"$set": bson.M{
				"nombre":        body.Nombre,
				"descripcion":   body.Descripcion,
				"fechaCreacion": time.Now(),
			},
		},
	)
	if err != nil {
		return
	}

	err = db.Collection(COLECCION_CATEGORIAS).FindOne(
		context.TODO(),
		bson.M{"_id": objId},
	).Decode(&categoria)
	return
}

func EliminarCategoria(ID string) (err error) {
	db := app.GetConnection()
	objId, _ := primitive.ObjectIDFromHex(ID)
	_, err = db.Collection(COLECCION_CATEGORIAS).DeleteOne(
		context.TODO(),
		bson.M{"_id": objId},
	)
	return
}

func ObtenerCategoria(ID string) (categoria models.Categoria, err error) {
	db := app.GetConnection()
	objId, _ := primitive.ObjectIDFromHex(ID)
	err = db.Collection(COLECCION_CATEGORIAS).FindOne(
		context.TODO(),
		bson.M{"_id": objId},
	).Decode(&categoria)
	if err != nil {
		return
	}
	return
}

func ListarCategorias() (categorias []models.Categoria, err error) {
	categorias = make([]models.Categoria, 0) // inicializar

	db := app.GetConnection()

	results, err := db.Collection(COLECCION_CATEGORIAS).Find(
		context.TODO(),
		bson.M{},
	)
	if err != nil {
		return
	}

	for results.Next(context.TODO()) {
		var categoria models.Categoria
		err = results.Decode(&categoria)
		if err != nil {
			return
		}
		categorias = append(categorias, categoria)
	}

	return
}

func ObtenerCategoriaPorId(ID string) (categoria models.Categoria, err error) {
	db := app.GetConnection()
	objId, _ := primitive.ObjectIDFromHex(ID)
	err = db.Collection(COLECCION_CATEGORIAS).FindOne(context.TODO(), bson.M{"_id": objId}).Decode(&categoria)
	return
}
