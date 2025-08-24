package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Players struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string             `bson:"name" json:"name"`
}
