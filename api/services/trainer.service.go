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

type TrainerService interface {
	CreateTrainer(models.Trainer, context.Context) error
	GetTrainer(*primitive.ObjectID, context.Context) (*models.Trainer, error)
	GetAllTrainers(context.Context) ([]*models.Trainer, error)
	UpdateTrainer(*primitive.ObjectID, models.UpdateTrainer, context.Context) error
	DeleteTrainer(*primitive.ObjectID, context.Context) error
}

type TrainerServiceImpl struct {
	trainerCollection *mongo.Collection
}

func NewTrainerService(trainerCollection *mongo.Collection) TrainerService {
	return &TrainerServiceImpl{
		trainerCollection: trainerCollection,
	}
}

func (service *TrainerServiceImpl) CreateTrainer(trainer models.Trainer, ctx context.Context) error {

	encryptedTrainer, errEncryption := lib.Encrypt(trainer)
	if errEncryption != nil {
		return errEncryption
	}

	payload := models.Trainer{
		FirstName:  encryptedTrainer.FirstName,
		LastName:   encryptedTrainer.LastName,
		StudentIDs: encryptedTrainer.Sport,
		Sport:      encryptedTrainer.Sport,
	}

	_, err := service.trainerCollection.InsertOne(ctx, payload)
	return err
}

func (service *TrainerServiceImpl) GetTrainer(id *primitive.ObjectID, ctx context.Context) (*models.Trainer, error) {
	var trainer *models.Trainer
	query := bson.D{bson.E{Key: "_id", Value: id}}
	err := service.trainerCollection.FindOne(ctx, query).Decode(&trainer)
	if err != nil {
		return nil, err
	}

	decryptedFirstName, err := lib.DecryptString(trainer.FirstName, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
	if err != nil {
		return nil, err
	}
	trainer.FirstName = decryptedFirstName

	decryptedLastName, err := lib.DecryptString(trainer.LastName, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
	if err != nil {
		return nil, err
	}
	trainer.LastName = decryptedLastName

	return trainer, err
}

func (service *TrainerServiceImpl) GetAllTrainers(ctx context.Context) ([]*models.Trainer, error) {
	var trainers []*models.Trainer = make([]*models.Trainer, 0)
	cursor, err := service.trainerCollection.Find(ctx, bson.D{{}})

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var trainer models.Trainer
		err := cursor.Decode(&trainer)
		if err != nil {
			return nil, err
		}

		decryptedFirstName, err := lib.DecryptString(trainer.FirstName, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
		if err != nil {
			return nil, err
		}
		trainer.FirstName = decryptedFirstName

		decryptedLastName, err := lib.DecryptString(trainer.LastName, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
		if err != nil {
			return nil, err
		}
		trainer.LastName = decryptedLastName

		trainers = append(trainers, &trainer)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	return trainers, nil
}

func (service *TrainerServiceImpl) UpdateTrainer(id *primitive.ObjectID, trainer models.UpdateTrainer, ctx context.Context) error {

	encryptedTrainer, errEncryption := lib.Encrypt(trainer)
	if errEncryption != nil {
		return errEncryption
	}

	filter := bson.D{bson.E{Key: "_id", Value: id}}
	update := bson.D{bson.E{Key: "$set", Value: encryptedTrainer}}
	res, err := service.trainerCollection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		return errors.New("trainer not found")
	}

	return nil
}

func (service *TrainerServiceImpl) DeleteTrainer(id *primitive.ObjectID, ctx context.Context) error {
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	res, err := service.trainerCollection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	if res.DeletedCount != 1 {
		return errors.New("trainer not found")
	}

	return nil
}
