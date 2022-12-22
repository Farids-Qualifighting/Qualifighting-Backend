package services

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"qualifighting.backend.de/models"
)

type SubjectService interface {
	CreateSubject(*models.Subject, context.Context) error
	GetSubject(*primitive.ObjectID, context.Context) (*models.Subject, error)
	GetAllSubjects(context.Context) ([]*models.Subject, error)
	UpdateSubject(*primitive.ObjectID, *models.UpdateSubject, context.Context) error
	DeleteSubject(*primitive.ObjectID, context.Context) error
}

type SubjectServiceImpl struct {
	subjectCollection *mongo.Collection
}

func NewSubjectService(subjectCollection *mongo.Collection) SubjectService {
	return &SubjectServiceImpl{
		subjectCollection: subjectCollection,
	}
}

func (service *SubjectServiceImpl) CreateSubject(subject *models.Subject, ctx context.Context) error {

	payload := models.Subject{
		ID:         primitive.NewObjectID(),
		Name:       subject.Name,
		Teacher:    subject.Teacher,
		SchoolID:   subject.SchoolID,
		StudentIDs: subject.StudentIDs,
	}

	_, err := service.subjectCollection.InsertOne(ctx, payload)
	return err
}

func (service *SubjectServiceImpl) GetSubject(id *primitive.ObjectID, ctx context.Context) (*models.Subject, error) {
	var subject *models.Subject
	query := bson.D{bson.E{Key: "_id", Value: id}}
	err := service.subjectCollection.FindOne(ctx, query).Decode(&subject)
	return subject, err
}

func (service *SubjectServiceImpl) GetAllSubjects(ctx context.Context) ([]*models.Subject, error) {
	var subjects []*models.Subject
	cursor, err := service.subjectCollection.Find(ctx, bson.D{{}})

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var subject models.Subject
		err := cursor.Decode(&subject)
		if err != nil {
			return nil, err
		}
		subjects = append(subjects, &subject)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	if len(subjects) == 0 {
		return nil, errors.New("no subjects in database")
	}

	return subjects, nil
}

func (service *SubjectServiceImpl) UpdateSubject(id *primitive.ObjectID, subject *models.UpdateSubject, ctx context.Context) error {
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	update := bson.D{bson.E{Key: "$set", Value: subject}}
	res, err := service.subjectCollection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		return errors.New("subject not found")
	}

	return nil
}

func (service *SubjectServiceImpl) DeleteSubject(id *primitive.ObjectID, ctx context.Context) error {
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	res, err := service.subjectCollection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	if res.DeletedCount != 1 {
		return errors.New("subject not found")
	}

	return nil
}
