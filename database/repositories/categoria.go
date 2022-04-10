package repositories

import (
	"context"
	"time"

	"habilitacion_backend/app"
	"habilitacion_backend/database/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Se define una constante con la colección de categorias
const COLECCION_CATEGORIAS = "categoria"

/*
Funciones para crear una nueva categoria
se realiza la conexión a la base de datos creada en el archivo database.go
se crea una nueva instancia de la colección de categorias
se inserta el nuevo registro en la colección
*/
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

/*
Funciones para actualizar una categoria
se realiza la conexión a la base de datos creada en el archivo database.go
se crea una nueva instancia de la colección de categorias con las nuevas actualizaciones
se actualiza el registro en la colección
*/
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

/*
Funciones para eliminar una categoria
se realiza la conexión a la base de datos creada en el archivo database.go
se elimina el registro en la colección buscandolo por el ID
*/
func EliminarCategoria(ID string) (err error) {
	db := app.GetConnection()
	objId, _ := primitive.ObjectIDFromHex(ID)
	_, err = db.Collection(COLECCION_CATEGORIAS).DeleteOne(
		context.TODO(),
		bson.M{"_id": objId},
	)
	return
}

/*
Funciones para obtener una categoria
se realiza la conexión a la base de datos creada en el archivo database.go
se busca la categoria por medio del ID
*/
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

/*
Funciones para obtener todas las categorias
se realiza la conexión a la base de datos creada en el archivo database.go
se buscan todas las categorias y se asignan en una nueva lista
*/
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
