package services

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"qualifighting.backend.de/models"
)

type ExamService interface {
	CreateExam(*models.Exam, context.Context) error
	GetExam(*primitive.ObjectID, context.Context) (*models.Exam, error)
	GetAllExams(context.Context) ([]*models.Exam, error)
	UpdateExam(*primitive.ObjectID, *models.UpdateExam, context.Context) error
	DeleteExam(*primitive.ObjectID, context.Context) error
}

type ExamServiceImpl struct {
	examCollection *mongo.Collection
}

func NewExamService(examCollection *mongo.Collection) ExamService {
	return &ExamServiceImpl{
		examCollection: examCollection,
	}
}

func (service *ExamServiceImpl) CreateExam(exam *models.Exam, ctx context.Context) error {

	payload := models.Exam{
		Topics:   exam.Topics,
		Date:     exam.Date,
		Students: exam.Students,
	}

	_, err := service.examCollection.InsertOne(ctx, payload)
	return err
}

func (service *ExamServiceImpl) GetExam(id *primitive.ObjectID, ctx context.Context) (*models.Exam, error) {
	var exam *models.Exam
	query := bson.D{bson.E{Key: "_id", Value: id}}
	err := service.examCollection.FindOne(ctx, query).Decode(&exam)
	return exam, err
}

func (service *ExamServiceImpl) GetAllExams(ctx context.Context) ([]*models.Exam, error) {
	var exams []*models.Exam = make([]*models.Exam, 0)
	cursor, err := service.examCollection.Find(ctx, bson.D{{}})

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var exam models.Exam
		err := cursor.Decode(&exam)
		if err != nil {
			return nil, err
		}
		exams = append(exams, &exam)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	return exams, nil
}

func (service *ExamServiceImpl) UpdateExam(id *primitive.ObjectID, exam *models.UpdateExam, ctx context.Context) error {
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	update := bson.D{bson.E{Key: "$set", Value: exam}}
	res, err := service.examCollection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		return errors.New("exam not found")
	}

	return nil
}

func (service *ExamServiceImpl) DeleteExam(id *primitive.ObjectID, ctx context.Context) error {
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	res, err := service.examCollection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	if res.DeletedCount != 1 {
		return errors.New("exam not found")
	}

	return nil
}
