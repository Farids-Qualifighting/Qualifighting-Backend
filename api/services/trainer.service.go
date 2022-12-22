package services

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"qualifighting.backend.de/models"
)

type TrainerService interface {
	CreateTrainer(*models.Trainer, context.Context) error
	GetTrainer(*primitive.ObjectID, context.Context) (*models.Trainer, error)
	GetAllTrainers(context.Context) ([]*models.Trainer, error)
	UpdateTrainer(*primitive.ObjectID, *models.UpdateTrainer, context.Context) error
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

func (service *TrainerServiceImpl) CreateTrainer(trainer *models.Trainer, ctx context.Context) error {

	payload := models.Trainer{
		ID:         primitive.NewObjectID(),
		FirstName:  trainer.FirstName,
		LastName:   trainer.LastName,
		StudentIDs: trainer.Sport,
		Sport:      trainer.Sport,
	}

	_, err := service.trainerCollection.InsertOne(ctx, payload)
	return err
}

func (service *TrainerServiceImpl) GetTrainer(id *primitive.ObjectID, ctx context.Context) (*models.Trainer, error) {
	var trainer *models.Trainer
	query := bson.D{bson.E{Key: "_id", Value: id}}
	err := service.trainerCollection.FindOne(ctx, query).Decode(&trainer)
	return trainer, err
}

func (service *TrainerServiceImpl) GetAllTrainers(ctx context.Context) ([]*models.Trainer, error) {
	var trainers []*models.Trainer
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
		trainers = append(trainers, &trainer)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	if len(trainers) == 0 {
		return nil, errors.New("no trainers in database")
	}

	return trainers, nil
}

func (service *TrainerServiceImpl) UpdateTrainer(id *primitive.ObjectID, trainer *models.UpdateTrainer, ctx context.Context) error {
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	update := bson.D{bson.E{Key: "$set", Value: trainer}}
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
