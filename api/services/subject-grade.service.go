package services

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"qualifighting.backend.de/models"
)

type SubjectGradeService interface {
	CreateSubjectGrade(*models.SubjectGrade, context.Context) error
	GetSubjectGrade(*primitive.ObjectID, context.Context) (*models.SubjectGrade, error)
	GetAllSubjectGrades(context.Context) ([]*models.SubjectGrade, error)
	UpdateSubjectGrade(*primitive.ObjectID, *models.UpdateSubjectGrade, context.Context) error
	DeleteSubjectGrade(*primitive.ObjectID, context.Context) error
}

type SubjectGradeServiceImpl struct {
	subjectGradeCollection *mongo.Collection
}

func NewSubjectGradeService(subjectGradeCollection *mongo.Collection) SubjectGradeService {
	return &SubjectGradeServiceImpl{
		subjectGradeCollection: subjectGradeCollection,
	}
}

func (service *SubjectGradeServiceImpl) CreateSubjectGrade(subjectGrade *models.SubjectGrade, ctx context.Context) error {

	payload := models.SubjectGrade{
		Grade:             subjectGrade.Grade,
		IntermediateGrade: subjectGrade.IntermediateGrade,
		StudentID:         subjectGrade.StudentID,
	}

	_, err := service.subjectGradeCollection.InsertOne(ctx, payload)
	return err
}

func (service *SubjectGradeServiceImpl) GetSubjectGrade(id *primitive.ObjectID, ctx context.Context) (*models.SubjectGrade, error) {
	var subjectGrade *models.SubjectGrade
	query := bson.D{bson.E{Key: "_id", Value: id}}
	err := service.subjectGradeCollection.FindOne(ctx, query).Decode(&subjectGrade)
	return subjectGrade, err
}

func (service *SubjectGradeServiceImpl) GetAllSubjectGrades(ctx context.Context) ([]*models.SubjectGrade, error) {
	var subjectGrades []*models.SubjectGrade
	cursor, err := service.subjectGradeCollection.Find(ctx, bson.D{{}})

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var subjectGrade models.SubjectGrade
		err := cursor.Decode(&subjectGrade)
		if err != nil {
			return nil, err
		}
		subjectGrades = append(subjectGrades, &subjectGrade)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	if len(subjectGrades) == 0 {
		return nil, errors.New("no subjectGrades in database")
	}

	return subjectGrades, nil
}

func (service *SubjectGradeServiceImpl) UpdateSubjectGrade(id *primitive.ObjectID, subjectGrade *models.UpdateSubjectGrade, ctx context.Context) error {
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	update := bson.D{bson.E{Key: "$set", Value: subjectGrade}}
	res, err := service.subjectGradeCollection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		return errors.New("subjectGrade not found")
	}

	return nil
}

func (service *SubjectGradeServiceImpl) DeleteSubjectGrade(id *primitive.ObjectID, ctx context.Context) error {
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	res, err := service.subjectGradeCollection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	if res.DeletedCount != 1 {
		return errors.New("subjectGrade not found")
	}

	return nil
}
