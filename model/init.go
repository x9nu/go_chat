package model

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Mongo = InitDB()

func InitDB() *mongo.Database {
	// 1.initialize
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{
		Username: "admin",
		Password: "admin",
	}).ApplyURI("mongodb://10.0.2.15:27017"))
	if err != nil {
		log.Println("Mongo connection:", err)
		return nil
	}
	// 2.connect
	return client.Database("im_gin_mongodb")
}
