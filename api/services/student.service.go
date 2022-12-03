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
	CreateStudent(*models.Student) error
	GetStudent(*primitive.ObjectID) (*models.Student, error)
	GetAll() ([]*models.Student, error)
	UpdateStudent(*primitive.ObjectID, *models.Student) error
	DeleteStudent(*primitive.ObjectID) error
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

	payload := models.Student{
		ID:           primitive.NewObjectID(),
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

	_, err := service.studentCollection.InsertOne(service.ctx, payload)
	return err
}

func (service *StudentServiceImpl) GetStudent(id *primitive.ObjectID) (*models.Student, error) {
	var student *models.Student
	query := bson.D{bson.E{Key: "_id", Value: id}}
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

func (service *StudentServiceImpl) UpdateStudent(id *primitive.ObjectID, student *models.Student) error {
	filter := bson.D{bson.E{Key: "_id", Value: id}}
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

func (service *StudentServiceImpl) DeleteStudent(id *primitive.ObjectID) error {
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	res, _ := service.studentCollection.DeleteOne(service.ctx, filter)

	if res.DeletedCount != 1 {
		return errors.New("student not found")
	}

	return nil
}
