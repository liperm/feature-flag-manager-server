package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection
var Ctx = context.TODO()

func Init() {
	clientOptions := options.Client().ApplyURI("mongodb://root:root@localhost:27017/")
	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	Collection = client.Database("feature_flag").Collection("feature_flags")
}
