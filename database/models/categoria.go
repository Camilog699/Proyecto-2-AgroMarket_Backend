package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Categoria struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre        string             `json:"nombre" bson:"nombre"`
	Descripcion   string             `json:"descripcion" bson:"descripcion"`
	FechaCreacion primitive.DateTime `json:"fechaCreacion" bson:"fechaCreacion,omitempty"`
}
