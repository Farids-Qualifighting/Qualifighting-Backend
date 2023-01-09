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

type DailyNoteService interface {
	CreateDailyNote(models.DailyNote, context.Context) error
	GetDailyNote(*primitive.ObjectID, context.Context) (*models.DailyNote, error)
	GetAllDailyNotes(context.Context) ([]*models.DailyNote, error)
	UpdateDailyNote(*primitive.ObjectID, models.UpdateDailyNote, context.Context) error
	DeleteDailyNote(*primitive.ObjectID, context.Context) error
}

type DailyNoteServiceImpl struct {
	dailyNoteCollection *mongo.Collection
}

func NewDailyNoteService(dailyNoteCollection *mongo.Collection) DailyNoteService {
	return &DailyNoteServiceImpl{
		dailyNoteCollection: dailyNoteCollection,
	}
}

func (service *DailyNoteServiceImpl) CreateDailyNote(dailyNote models.DailyNote, ctx context.Context) error {

	encryptedDailyNote, errEncryption := lib.Encrypt(dailyNote)
	if errEncryption != nil {
		return errEncryption
	}

	payload := models.DailyNote{
		StudentID: encryptedDailyNote.StudentID,
		CreatedAt: encryptedDailyNote.CreatedAt,
		Subject:   encryptedDailyNote.Subject,
		Rating:    encryptedDailyNote.Rating,
		Note:      encryptedDailyNote.Note,
		CreatorID: encryptedDailyNote.CreatorID,
		UpdatedAt: encryptedDailyNote.UpdatedAt,
	}

	_, err := service.dailyNoteCollection.InsertOne(ctx, payload)
	return err
}

func (service *DailyNoteServiceImpl) GetDailyNote(id *primitive.ObjectID, ctx context.Context) (*models.DailyNote, error) {
	var dailyNote *models.DailyNote
	query := bson.D{bson.E{Key: "_id", Value: id}}
	err := service.dailyNoteCollection.FindOne(ctx, query).Decode(&dailyNote)
	if err != nil {
		return nil, err
	}
	decryptedNote, err := lib.DecryptString(dailyNote.Note, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
	dailyNote.Note = decryptedNote
	return dailyNote, err
}

func (service *DailyNoteServiceImpl) GetAllDailyNotes(ctx context.Context) ([]*models.DailyNote, error) {
	var dailyNotes []*models.DailyNote = make([]*models.DailyNote, 0)
	cursor, err := service.dailyNoteCollection.Find(ctx, bson.D{{}})

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var dailyNote models.DailyNote
		err := cursor.Decode(&dailyNote)
		if err != nil {
			return nil, err
		}
		decryptedNote, err := lib.DecryptString(dailyNote.Note, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
		if err != nil {
			return nil, err
		}
		dailyNote.Note = decryptedNote
		dailyNotes = append(dailyNotes, &dailyNote)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	return dailyNotes, nil
}

func (service *DailyNoteServiceImpl) UpdateDailyNote(id *primitive.ObjectID, dailyNote models.UpdateDailyNote, ctx context.Context) error {

	encryptedDailyNote, errEncryption := lib.Encrypt(dailyNote)
	if errEncryption != nil {
		return errEncryption
	}

	filter := bson.D{bson.E{Key: "_id", Value: id}}
	update := bson.D{bson.E{Key: "$set", Value: encryptedDailyNote}}
	res, err := service.dailyNoteCollection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		return errors.New("dailyNote not found")
	}

	return nil
}

func (service *DailyNoteServiceImpl) DeleteDailyNote(id *primitive.ObjectID, ctx context.Context) error {
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	res, err := service.dailyNoteCollection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	if res.DeletedCount != 1 {
		return errors.New("dailyNote not found")
	}

	return nil
}
