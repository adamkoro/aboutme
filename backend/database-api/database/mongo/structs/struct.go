package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type Picture struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name,omitempty" bson:"name,omitempty"`
	Base64 string             `json:"base64,omitempty" bson:"base64,omitempty"`
}
