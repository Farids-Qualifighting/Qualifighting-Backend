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

func MongoDBStudentCollection() *mongo.Collection {
	return mongoInstance.Database("qualifighting").Collection("student")
}

func MongoDBDailyNoteCollection() *mongo.Collection {
	return mongoInstance.Database("qualifighting").Collection("daily_note")
}

func MongoDBExamCollection() *mongo.Collection {
	return mongoInstance.Database("qualifighting").Collection("exam")
}

func MongoDBParentCollection() *mongo.Collection {
	return mongoInstance.Database("qualifighting").Collection("parent")
}

func MongoDBSchoolCollection() *mongo.Collection {
	return mongoInstance.Database("qualifighting").Collection("school")
}

func MongoDBSportCollection() *mongo.Collection {
	return mongoInstance.Database("qualifighting").Collection("sport")
}

func MongoDBSubjectGradeCollection() *mongo.Collection {
	return mongoInstance.Database("qualifighting").Collection("subject_grade")
}

func MongoDBSubjectCollection() *mongo.Collection {
	return mongoInstance.Database("qualifighting").Collection("subject")
}

func MongoDBTeacherCollection() *mongo.Collection {
	return mongoInstance.Database("qualifighting").Collection("teacher")
}

func MongoDBTrainerCollection() *mongo.Collection {
	return mongoInstance.Database("qualifighting").Collection("trainer")
}

func MongoDBTutorCollection() *mongo.Collection {
	return mongoInstance.Database("qualifighting").Collection("tutor")
}

func MongoDBCompetitionCollection() *mongo.Collection {
	return mongoInstance.Database("qualifighting").Collection("competition")
}
