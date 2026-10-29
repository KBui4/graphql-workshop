package configs

import (
	"context"
	"log"
	"os"
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

		client, err := mongo.Connect(options.Client().ApplyURI(GetEnvMongoURI()))
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
	return ConnectDB().Database("combinedDB").Collection(collectionName)
}

func GetEnvMongoURI() string {
	uri, exists := os.LookupEnv("MONGODB_URI")
	if !exists {
		log.Fatalf("Error loading 'MONGODB_URI' from the .env file")
	}

	return uri
}
