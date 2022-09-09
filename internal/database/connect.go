package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoS struct {
	Client   *mongo.Client
	Stream   *mongo.Collection
	Members  *mongo.Collection
	Requests *mongo.Collection
}

var (
	Mongo *mongoS
)

func init() {
	Mongo = new(mongoS)
}

func Connect(mongoUri string) error {
	var err error

	mdbClientOptions := options.Client().ApplyURI(mongoUri)
	Mongo.Client, err = mongo.Connect(context.TODO(), mdbClientOptions)

	if err != nil {
		log.Fatal(err)
		return err
	}

	err = Mongo.Client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return err
	}

	ggdb := Mongo.Client.Database("tgg")
	Mongo.Stream = ggdb.Collection("streams")
	Mongo.Members = ggdb.Collection("members")
	Mongo.Requests = ggdb.Collection("requests")

	fmt.Println("[INFO] Connected to MongoDB")

	return err
}
