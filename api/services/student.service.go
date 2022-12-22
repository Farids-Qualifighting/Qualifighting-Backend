package services

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"qualifighting.backend.de/models"
)

type StudentService interface {
	CreateStudent(*models.Student, context.Context) error
	GetStudent(*primitive.ObjectID, context.Context) (*models.Student, error)
	GetAllStudents(context.Context) ([]*models.Student, error)
	UpdateStudent(*primitive.ObjectID, *models.UpdateStudent, context.Context) error
	DeleteStudent(*primitive.ObjectID, context.Context) error
}

type StudentServiceImpl struct {
	studentCollection *mongo.Collection
}

func NewStudentService(studentCollection *mongo.Collection) StudentService {
	return &StudentServiceImpl{
		studentCollection: studentCollection,
	}
}

func (service *StudentServiceImpl) CreateStudent(student *models.Student, ctx context.Context) error {

	payload := models.Student{
		FirstName:    student.FirstName,
		LastName:     student.LastName,
		ClassTeacher: student.ClassTeacher,
		Birthday:     student.Birthday,
		Gender:       student.Gender,
		Address:      student.Address,
		Phone:        student.Phone,
		Email:        student.Email,
		SocialMedia:  student.SocialMedia,
		Certificate:  student.Certificate,
	}

	_, err := service.studentCollection.InsertOne(ctx, payload)
	return err
}

func (service *StudentServiceImpl) GetStudent(id *primitive.ObjectID, ctx context.Context) (*models.Student, error) {
	var student *models.Student
	query := bson.D{bson.E{Key: "_id", Value: id}}
	err := service.studentCollection.FindOne(ctx, query).Decode(&student)
	return student, err
}

func (service *StudentServiceImpl) GetAllStudents(ctx context.Context) ([]*models.Student, error) {
	var students []*models.Student
	cursor, err := service.studentCollection.Find(ctx, bson.D{{}})

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var student models.Student
		err := cursor.Decode(&student)
		if err != nil {
			return nil, err
		}
		students = append(students, &student)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	if len(students) == 0 {
		return nil, errors.New("no students in database")
	}

	return students, nil
}

func (service *StudentServiceImpl) UpdateStudent(id *primitive.ObjectID, student *models.UpdateStudent, ctx context.Context) error {
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	update := bson.D{bson.E{Key: "$set", Value: student}}
	res, err := service.studentCollection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		return errors.New("student not found")
	}

	return nil
}

func (service *StudentServiceImpl) DeleteStudent(id *primitive.ObjectID, ctx context.Context) error {
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	res, err := service.studentCollection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	if res.DeletedCount != 1 {
		return errors.New("student not found")
	}

	return nil
}
