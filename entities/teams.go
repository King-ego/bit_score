package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Teams struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name         string             `bson:"name" json:"name"`
	PrimaryColor string             `bson:"primary_color" json:"primary_color"`
	SecondColor  string             `bson:"second_color" json:"second_color"`
}
