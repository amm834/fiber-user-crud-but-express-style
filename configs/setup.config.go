package configs

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// ConnectDB connects to the database and returns a client
// if it cannot connect to the database, it will timeout after 10 seconds
// and log an error message
func ConnectDB() *mongo.Client {
	log.Println("Trying to connect to MongoDB...")
	// connect to mongo db
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvGoURI()))
	if err != nil {
		log.Fatal(err)
	}
	// create timeout for mongodb connection
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// ping the mongodb server
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")
	return client
}

// Client mongodb client instance
var Client *mongo.Client = ConnectDB()

// GetCollection returns a collection from the database
func GetCollection(client *mongo.Client, collection string) *mongo.Collection {
	return client.Database(EnvGoDB()).Collection(collection)
}
