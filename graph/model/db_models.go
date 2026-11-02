package model

import "go.mongodb.org/mongo-driver/v2/bson"

// Database models for MongoDB

// GraphQL generates the ID fields as strings (see models_gen.go)
// For any objects that reference another with ID, when retrieving data
// we need to convert from the database model to the GraphQL model.

type DBUser struct {
	ID            bson.ObjectID            `bson:"_id"`
	Name          string                   `bson:"name"`
	Hobbies       []string                 `bson:"hobbies"`
	Products      []bson.ObjectID          `bson:"products"`
}

type DBProduct struct {
	ID            bson.ObjectID            `bson:"_id"`
	Seller        bson.ObjectID            `bson:"seller"`
	Price         float64                  `bson:"price"`
	Name          string                   `bson:"name"`
	Description   string                   `bson:"description"`
}

// Challenge 2 Part A


// Transform functions
// These functions convert from the database model to the GraphQL model.

func TransformUser(dbUser *DBUser) *User {
	if dbUser == nil {
		return nil
	}

	products := make([]*Product, len(dbUser.Products))
	for i, productID := range dbUser.Products {
		products[i] = &Product{ID: productID.Hex()}
	}

	return &User{
		ID:       dbUser.ID.Hex(),
		Name:     dbUser.Name,
		Hobbies:  dbUser.Hobbies,
		Products: products,
	}
}

func TransformProduct(dbProduct *DBProduct) *Product {
	if dbProduct == nil {
		return nil
	}

	userRef := &User{
		ID: dbProduct.Seller.Hex(),
	}

	return &Product{
		ID:          dbProduct.ID.Hex(),
		Seller:      userRef,
		Price:       dbProduct.Price,
		Name:        dbProduct.Name,
		Description: dbProduct.Description,
	}
}

// Challenge 2 Part B

