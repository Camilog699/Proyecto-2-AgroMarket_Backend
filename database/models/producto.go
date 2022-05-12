package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Producto struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre               string             `json:"nombre" bson:"nombre"`
	Descripcion          string             `json:"descripcion" bson:"descripcion"`
	Precio               string             `json:"precio" bson:"precio"`
	Cantidad             string             `json:"cantidad" bson:"cantidad"`
	FechaCreacion        primitive.DateTime `json:"fechaCreacion" bson:"fechaCreacion,omitempty"`
	FechaModificacion    primitive.DateTime `json:"fechaModificacion" bson:"fechaModificacion,omitempty"`
	IdCategoriaProducto  string             `json:"idCategoriaProducto" bson:"idCategoriaProducto"`
	IdUsuarioCreador     string             `json:"idUsuarioCreador" bson:"idUsuarioCreador"`
	IdUsuarioModificador string             `json:"idUsuarioModificador" bson:"idUsuarioModificador"`
}
