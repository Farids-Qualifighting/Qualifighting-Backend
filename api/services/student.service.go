package services

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"qualifighting.backend.de/models"
)

type StudentService interface {
	CreateStudent(*models.Student) error
	GetStudent(*string) (*models.Student, error)
	GetAll() ([]*models.Student, error)
	UpdateStudent(*models.Student) error
	DeleteStudent(*string) error
}

type StudentServiceImpl struct {
	studentCollection *mongo.Collection
	ctx               context.Context
}

func NewStudentService(studentCollection *mongo.Collection, ctx context.Context) StudentService {
	return &StudentServiceImpl{
		studentCollection: studentCollection,
		ctx:               ctx,
	}
}

func (service *StudentServiceImpl) CreateStudent(student *models.Student) error {
	_, err := service.studentCollection.InsertOne(service.ctx, student)
	return err
}

func (service *StudentServiceImpl) GetStudent(name *string) (*models.Student, error) {
	var student *models.Student
	query := bson.D{bson.E{Key: "first_name", Value: name}}
	err := service.studentCollection.FindOne(service.ctx, query).Decode(&student)
	return student, err
}

func (service *StudentServiceImpl) GetAll() ([]*models.Student, error) {
	var students []*models.Student
	cursor, err := service.studentCollection.Find(service.ctx, bson.D{{}})

	if err != nil {
		return nil, err
	}

	for cursor.Next(service.ctx) {
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

	cursor.Close(service.ctx)

	if len(students) == 0 {
		return nil, errors.New("no students in database")
	}

	return students, nil
}

func (service *StudentServiceImpl) UpdateStudent(student *models.Student) error {
	filter := bson.D{bson.E{Key: "first_name", Value: student.FirstName}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "first_name", Value: student.FirstName},
		bson.E{Key: "last_name", Value: student.LastName},
		bson.E{Key: "class_teacher", Value: student.ClassTeacher},
		bson.E{Key: "birthday", Value: student.Birthday},
		bson.E{Key: "gender", Value: student.Gender},
		bson.E{Key: "address", Value: student.Address},
		bson.E{Key: "phone", Value: student.Phone},
		bson.E{Key: "email", Value: student.Email},
		bson.E{Key: "social_media", Value: student.SocialMedia},
		bson.E{Key: "certificate", Value: student.Certificate}}}}
	res, _ := service.studentCollection.UpdateOne(service.ctx, filter, update)

	if res.MatchedCount != 1 {
		return errors.New("student not found")
	}

	return nil
}

func (service *StudentServiceImpl) DeleteStudent(name *string) error {
	filter := bson.D{bson.E{Key: "first_name", Value: name}}
	res, _ := service.studentCollection.DeleteOne(service.ctx, filter)

	if res.DeletedCount != 1 {
		return errors.New("student not found")
	}

	return nil
}
