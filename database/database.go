package database

import (
	"context"
	"log"
	"time"

	"graphyy/internal/envvar"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// GetDatabase returns a database instance.
func InitDatabase() *mongo.Database {
	url := envvar.MongoURL()
	dbName := envvar.DBName()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database(dbName)
	collection := db.Collection(envvar.PointCollection())
	models := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{Key: "location", Value: bsonx.String("2dsphere")}},
		},
	}

	// Declare an options object
	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	_, err = collection.Indexes().CreateMany(ctx, models, opts)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
