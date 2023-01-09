package services

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"qualifighting.backend.de/lib"
	"qualifighting.backend.de/models"
)

type StudentService interface {
	CreateStudent(models.Student, context.Context) error
	GetStudent(*primitive.ObjectID, context.Context) (*models.Student, error)
	GetAllStudents(context.Context) ([]*models.Student, error)
	UpdateStudent(*primitive.ObjectID, models.UpdateStudent, context.Context) error
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

func (service *StudentServiceImpl) CreateStudent(student models.Student, ctx context.Context) error {

	encryptedStudent, errEncryption := lib.Encrypt(student)
	if errEncryption != nil {
		return errEncryption
	}

	payload := models.Student{
		FirstName:    encryptedStudent.FirstName,
		LastName:     encryptedStudent.LastName,
		ClassTeacher: encryptedStudent.ClassTeacher,
		Birthday:     encryptedStudent.Birthday,
		Gender:       encryptedStudent.Gender,
		Address:      encryptedStudent.Address,
		Phone:        encryptedStudent.Phone,
		Email:        encryptedStudent.Email,
		SocialMedia:  encryptedStudent.SocialMedia,
		Certificate:  encryptedStudent.Certificate,
	}

	_, err := service.studentCollection.InsertOne(ctx, payload)
	return err
}

func (service *StudentServiceImpl) GetStudent(id *primitive.ObjectID, ctx context.Context) (*models.Student, error) {
	var student *models.Student
	query := bson.D{bson.E{Key: "_id", Value: id}}
	err := service.studentCollection.FindOne(ctx, query).Decode(&student)
	if err != nil {
		return nil, err
	}

	decryptedFirstName, err := lib.DecryptString(student.FirstName, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
	if err != nil {
		return nil, err
	}
	student.FirstName = decryptedFirstName

	decryptedLastName, err := lib.DecryptString(student.LastName, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
	if err != nil {
		return nil, err
	}
	student.LastName = decryptedLastName

	decryptedClassTeacher, err := lib.DecryptString(student.ClassTeacher, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
	if err != nil {
		return nil, err
	}
	student.ClassTeacher = decryptedClassTeacher

	decryptedGender, err := lib.DecryptString(student.Gender, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
	if err != nil {
		return nil, err
	}
	student.Gender = decryptedGender

	decryptedPhone, err := lib.DecryptString(student.Phone, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
	if err != nil {
		return nil, err
	}
	student.Phone = decryptedPhone

	decryptedEmail, err := lib.DecryptString(student.Email, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
	if err != nil {
		return nil, err
	}
	student.Email = decryptedEmail

	decryptedSocialMedia, err := lib.DecryptString(student.SocialMedia, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
	if err != nil {
		return nil, err
	}
	student.SocialMedia = decryptedSocialMedia

	return student, err
}

func (service *StudentServiceImpl) GetAllStudents(ctx context.Context) ([]*models.Student, error) {
	var students []*models.Student = make([]*models.Student, 0)
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

		decryptedFirstName, err := lib.DecryptString(student.FirstName, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
		if err != nil {
			return nil, err
		}
		student.FirstName = decryptedFirstName

		decryptedLastName, err := lib.DecryptString(student.LastName, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
		if err != nil {
			return nil, err
		}
		student.LastName = decryptedLastName

		decryptedClassTeacher, err := lib.DecryptString(student.ClassTeacher, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
		if err != nil {
			return nil, err
		}
		student.ClassTeacher = decryptedClassTeacher

		decryptedGender, err := lib.DecryptString(student.Gender, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
		if err != nil {
			return nil, err
		}
		student.Gender = decryptedGender

		decryptedPhone, err := lib.DecryptString(student.Phone, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
		if err != nil {
			return nil, err
		}
		student.Phone = decryptedPhone

		decryptedEmail, err := lib.DecryptString(student.Email, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
		if err != nil {
			return nil, err
		}
		student.Email = decryptedEmail

		decryptedSocialMedia, err := lib.DecryptString(student.SocialMedia, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
		if err != nil {
			return nil, err
		}
		student.SocialMedia = decryptedSocialMedia

		students = append(students, &student)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	return students, nil
}

func (service *StudentServiceImpl) UpdateStudent(id *primitive.ObjectID, student models.UpdateStudent, ctx context.Context) error {

	encryptedStudent, errEncryption := lib.Encrypt(student)
	if errEncryption != nil {
		return errEncryption
	}

	filter := bson.D{bson.E{Key: "_id", Value: id}}
	update := bson.D{bson.E{Key: "$set", Value: encryptedStudent}}
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
