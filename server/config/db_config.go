package config

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbClient *mongo.Client
var DEFAULT_DB_NAME = "test_db"
var syncOnce sync.Once

func init() {
	LoadEnvironmentVariable()
}

func InitDatabase() (*mongo.Client, error) {

	db_url := GetEnvProperty("database_url")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db_url))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database  connected succesfully !")

	return client, nil
}

func GetDatabaseCollection(dbName *string, collectionName string) *mongo.Collection {
	if *dbName == "" || dbName == nil {
		dbName = &DEFAULT_DB_NAME
	}
	dbClient, _ := InitDatabase()
	collection := dbClient.Database(*dbName).Collection(collectionName)
	return collection
}

func InitializeDB() (*mongo.Client, error) {
	db, err := InitDatabase()

	if err != nil {
		log.Fatalf("Failed to connect to db %v", err)
	}

	return db, nil
}
