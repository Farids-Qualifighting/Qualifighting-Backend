package lib

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	mongoInstance *mongo.Client
)

func newInstance(url string) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		panic(err)
	}

	err = client.Connect(context.TODO())
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		panic(err)
	}

	return client
}

func init() {
	mongoInstance = newInstance(appConfig.MongoDBURL)
}

func MongoDBSchoolCollection() *mongo.Collection {
	return mongoInstance.Database("qualifighting").Collection("school")
}
