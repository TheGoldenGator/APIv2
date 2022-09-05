package database

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/thegoldengator/APIv2/internal/config"
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
	RDB   *redis.Client
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

	// Redis Connection
	if config.Config.GetString("environment") == "dev" {
		RDB = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
	} else {
		RDB = redis.NewClient(&redis.Options{
			Addr:     "redis:6379",
			Password: "mypassword",
			DB:       0,
		})
	}

	pong, err := RDB.Ping(RDB.Context()).Result()
	if err != nil {
		log.Fatal(err)
		return err
	}

	if pong == "PONG" {
		fmt.Print("[INFO] Connected to Redis\n")
	}

	return err
}
