package database

import (
	"context"
	"log"
	"time"

	"fiber-app/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Config.MongoURI))
	if err != nil {
		log.Fatal(err)
	}

	DB = client.Database("fiber_db")
	log.Println("MongoDB connected",config.Config.MongoURI,"with this db",DB.Name())
}
