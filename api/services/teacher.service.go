package services

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"qualifighting.backend.de/models"
)

type TeacherService interface {
	CreateTeacher(*models.Teacher, context.Context) error
	GetTeacher(*primitive.ObjectID, context.Context) (*models.Teacher, error)
	GetAllTeachers(context.Context) ([]*models.Teacher, error)
	UpdateTeacher(*primitive.ObjectID, *models.UpdateTeacher, context.Context) error
	DeleteTeacher(*primitive.ObjectID, context.Context) error
}

type TeacherServiceImpl struct {
	teacherCollection *mongo.Collection
}

func NewTeacherService(teacherCollection *mongo.Collection) TeacherService {
	return &TeacherServiceImpl{
		teacherCollection: teacherCollection,
	}
}

func (service *TeacherServiceImpl) CreateTeacher(teacher *models.Teacher, ctx context.Context) error {

	payload := models.Teacher{
		FirstName: teacher.FirstName,
		LastName:  teacher.LastName,
		SchoolID:  teacher.SchoolID,
		Email:     teacher.Email,
		Phone:     teacher.Phone,
	}

	_, err := service.teacherCollection.InsertOne(ctx, payload)
	return err
}

func (service *TeacherServiceImpl) GetTeacher(id *primitive.ObjectID, ctx context.Context) (*models.Teacher, error) {
	var teacher *models.Teacher
	query := bson.D{bson.E{Key: "_id", Value: id}}
	err := service.teacherCollection.FindOne(ctx, query).Decode(&teacher)
	return teacher, err
}

func (service *TeacherServiceImpl) GetAllTeachers(ctx context.Context) ([]*models.Teacher, error) {
	var teachers []*models.Teacher
	cursor, err := service.teacherCollection.Find(ctx, bson.D{{}})

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var teacher models.Teacher
		err := cursor.Decode(&teacher)
		if err != nil {
			return nil, err
		}
		teachers = append(teachers, &teacher)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	if len(teachers) == 0 {
		return nil, errors.New("no teachers in database")
	}

	return teachers, nil
}

func (service *TeacherServiceImpl) UpdateTeacher(id *primitive.ObjectID, teacher *models.UpdateTeacher, ctx context.Context) error {
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	update := bson.D{bson.E{Key: "$set", Value: teacher}}
	res, err := service.teacherCollection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		return errors.New("teacher not found")
	}

	return nil
}

func (service *TeacherServiceImpl) DeleteTeacher(id *primitive.ObjectID, ctx context.Context) error {
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	res, err := service.teacherCollection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	if res.DeletedCount != 1 {
		return errors.New("teacher not found")
	}

	return nil
}
