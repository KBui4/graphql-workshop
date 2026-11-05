package model

import "go.mongodb.org/mongo-driver/v2/bson"

type Product struct {
	ID          bson.ObjectID `json:"id" bson:"_id"`
	Seller      bson.ObjectID `json:"seller" bson:"seller"`
	Price       float64       `json:"price" bson:"price"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
}

type User struct {
	ID       bson.ObjectID    `json:"id" bson:"_id"`
	Name     string           `json:"name" bson:"name"`
	Hobbies  []string         `json:"hobbies" bson:"hobbies"`
	Products []bson.ObjectID  `json:"products" bson:"products"`
}
