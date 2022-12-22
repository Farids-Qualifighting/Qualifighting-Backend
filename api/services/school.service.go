package services

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"qualifighting.backend.de/models"
)

type SchoolService interface {
	CreateSchool(*models.School, context.Context) error
	GetSchool(*primitive.ObjectID, context.Context) (*models.School, error)
	GetAllSchools(context.Context) ([]*models.School, error)
	UpdateSchool(*primitive.ObjectID, *models.UpdateSchool, context.Context) error
	DeleteSchool(*primitive.ObjectID, context.Context) error
}

type SchoolServiceImpl struct {
	schoolCollection *mongo.Collection
}

func NewSchoolService(schoolCollection *mongo.Collection) SchoolService {
	return &SchoolServiceImpl{
		schoolCollection: schoolCollection,
	}
}

func (service *SchoolServiceImpl) CreateSchool(school *models.School, ctx context.Context) error {

	payload := models.School{
		Name:     school.Name,
		Address:  school.Address,
		Email:    school.Email,
		Phone:    school.Phone,
		Teachers: school.Teachers,
	}

	_, err := service.schoolCollection.InsertOne(ctx, payload)
	return err
}

func (service *SchoolServiceImpl) GetSchool(id *primitive.ObjectID, ctx context.Context) (*models.School, error) {
	var school *models.School
	query := bson.D{bson.E{Key: "_id", Value: id}}
	err := service.schoolCollection.FindOne(ctx, query).Decode(&school)
	return school, err
}

func (service *SchoolServiceImpl) GetAllSchools(ctx context.Context) ([]*models.School, error) {
	var schools []*models.School
	cursor, err := service.schoolCollection.Find(ctx, bson.D{{}})

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var school models.School
		err := cursor.Decode(&school)
		if err != nil {
			return nil, err
		}
		schools = append(schools, &school)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	if len(schools) == 0 {
		return nil, errors.New("no schools in database")
	}

	return schools, nil
}

func (service *SchoolServiceImpl) UpdateSchool(id *primitive.ObjectID, school *models.UpdateSchool, ctx context.Context) error {
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	update := bson.D{bson.E{Key: "$set", Value: school}}
	res, err := service.schoolCollection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		return errors.New("school not found")
	}

	return nil
}

func (service *SchoolServiceImpl) DeleteSchool(id *primitive.ObjectID, ctx context.Context) error {
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	res, err := service.schoolCollection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	if res.DeletedCount != 1 {
		return errors.New("school not found")
	}

	return nil
}
