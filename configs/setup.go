package configs

import (
	"context"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type DBSingleton struct {
	client *mongo.Client
}

var dbInstance *DBSingleton
var once sync.Once

func ConnectDB() *mongo.Client {
	once.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client, err := mongo.Connect(options.Client().ApplyURI("mongodb+srv://workshop_user:ccqXXxl7ahiE4RTZ@graphql.snfilnf.mongodb.net/"))
		if err != nil {
			log.Fatalf("Unable to create MongoDB client")
		}

		if err = client.Ping(ctx, nil); err != nil {
			log.Fatalf("Unable to ping database")
		}

		log.Printf("Connected to MongoDB")
		dbInstance = &DBSingleton{client: client}
	})

	return dbInstance.client
}

func GetCollection(collectionName string) *mongo.Collection {
	return ConnectDB().Database("workshop_db").Collection(collectionName)
}
