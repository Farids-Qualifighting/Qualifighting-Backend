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

type TutorService interface {
	CreateTutor(models.Tutor, context.Context) error
	GetTutor(*primitive.ObjectID, context.Context) (*models.Tutor, error)
	GetAllTutors(context.Context) ([]*models.Tutor, error)
	UpdateTutor(*primitive.ObjectID, models.UpdateTutor, context.Context) error
	DeleteTutor(*primitive.ObjectID, context.Context) error
}

type TutorServiceImpl struct {
	tutorCollection *mongo.Collection
}

func NewTutorService(tutorCollection *mongo.Collection) TutorService {
	return &TutorServiceImpl{
		tutorCollection: tutorCollection,
	}
}

func (service *TutorServiceImpl) CreateTutor(tutor models.Tutor, ctx context.Context) error {

	encryptedTutor, errEncryption := lib.Encrypt(tutor)
	if errEncryption != nil {
		return errEncryption
	}

	payload := models.Tutor{
		FirstName: encryptedTutor.FirstName,
		LastName:  encryptedTutor.LastName,
		Email:     encryptedTutor.Email,
		Phone:     encryptedTutor.Phone,
		Subjects:  encryptedTutor.Subjects,
		Students:  encryptedTutor.Students,
	}

	_, err := service.tutorCollection.InsertOne(ctx, payload)
	return err
}

func (service *TutorServiceImpl) GetTutor(id *primitive.ObjectID, ctx context.Context) (*models.Tutor, error) {
	var tutor *models.Tutor
	query := bson.D{bson.E{Key: "_id", Value: id}}
	err := service.tutorCollection.FindOne(ctx, query).Decode(&tutor)
	if err != nil {
		return nil, err
	}

	decryptedFirstName, err := lib.DecryptString(tutor.FirstName, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
	if err != nil {
		return nil, err
	}
	tutor.FirstName = decryptedFirstName

	decryptedLastName, err := lib.DecryptString(tutor.LastName, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
	if err != nil {
		return nil, err
	}
	tutor.LastName = decryptedLastName

	decryptedPhone, err := lib.DecryptString(tutor.Phone, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
	if err != nil {
		return nil, err
	}
	tutor.Phone = decryptedPhone

	decryptedEmail, err := lib.DecryptString(tutor.Email, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
	if err != nil {
		return nil, err
	}
	tutor.Email = decryptedEmail

	return tutor, err
}

func (service *TutorServiceImpl) GetAllTutors(ctx context.Context) ([]*models.Tutor, error) {
	var tutors []*models.Tutor = make([]*models.Tutor, 0)
	cursor, err := service.tutorCollection.Find(ctx, bson.D{{}})

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var tutor models.Tutor
		err := cursor.Decode(&tutor)
		if err != nil {
			return nil, err
		}

		decryptedFirstName, err := lib.DecryptString(tutor.FirstName, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
		if err != nil {
			return nil, err
		}
		tutor.FirstName = decryptedFirstName

		decryptedLastName, err := lib.DecryptString(tutor.LastName, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
		if err != nil {
			return nil, err
		}
		tutor.LastName = decryptedLastName

		decryptedPhone, err := lib.DecryptString(tutor.Phone, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
		if err != nil {
			return nil, err
		}
		tutor.Phone = decryptedPhone

		decryptedEmail, err := lib.DecryptString(tutor.Email, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
		if err != nil {
			return nil, err
		}
		tutor.Email = decryptedEmail

		tutors = append(tutors, &tutor)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	return tutors, nil
}

func (service *TutorServiceImpl) UpdateTutor(id *primitive.ObjectID, tutor models.UpdateTutor, ctx context.Context) error {

	encryptedTutor, errEncryption := lib.Encrypt(tutor)
	if errEncryption != nil {
		return errEncryption
	}

	filter := bson.D{bson.E{Key: "_id", Value: id}}
	update := bson.D{bson.E{Key: "$set", Value: encryptedTutor}}
	res, err := service.tutorCollection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		return errors.New("tutor not found")
	}

	return nil
}

func (service *TutorServiceImpl) DeleteTutor(id *primitive.ObjectID, ctx context.Context) error {
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	res, err := service.tutorCollection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	if res.DeletedCount != 1 {
		return errors.New("tutor not found")
	}

	return nil
}
