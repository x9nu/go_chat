package test

import (
	"context"
	"fmt"
	"go_chat/model"
	"log"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestFindOne(t *testing.T) {
	// 1.initialize
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{
		Username: "admin",
		Password: "admin",
	}).ApplyURI("mongodb://10.0.2.15:27017"))
	if err != nil {
		log.Fatal("Mongo connection:", err)
	}
	// 2.connect
	db := client.Database("im_gin_mongodb")

	// operation
	coll := db.Collection("user_basic") //connect with collection(table)
	ub := new(model.UserBasic)
	err = coll.FindOne(context.Background(), bson.D{}).Decode(ub)
	if err != nil {
		log.Fatal("FindOne:", err)
	}

	fmt.Println("ub=>", ub)
}
