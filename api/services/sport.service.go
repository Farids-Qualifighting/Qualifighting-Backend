package services

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"qualifighting.backend.de/models"
)

type SportService interface {
	CreateSport(*models.Sport, context.Context) error
	GetSport(*primitive.ObjectID, context.Context) (*models.Sport, error)
	GetAllSports(context.Context) ([]*models.Sport, error)
	UpdateSport(*primitive.ObjectID, *models.UpdateSport, context.Context) error
	DeleteSport(*primitive.ObjectID, context.Context) error
}

type SportServiceImpl struct {
	sportCollection *mongo.Collection
}

func NewSportService(sportCollection *mongo.Collection) SportService {
	return &SportServiceImpl{
		sportCollection: sportCollection,
	}
}

func (service *SportServiceImpl) CreateSport(sport *models.Sport, ctx context.Context) error {

	payload := models.Sport{
		Name: sport.Name,
	}

	_, err := service.sportCollection.InsertOne(ctx, payload)
	return err
}

func (service *SportServiceImpl) GetSport(id *primitive.ObjectID, ctx context.Context) (*models.Sport, error) {
	var sport *models.Sport
	query := bson.D{bson.E{Key: "_id", Value: id}}
	err := service.sportCollection.FindOne(ctx, query).Decode(&sport)
	return sport, err
}

func (service *SportServiceImpl) GetAllSports(ctx context.Context) ([]*models.Sport, error) {
	var sports []*models.Sport
	cursor, err := service.sportCollection.Find(ctx, bson.D{{}})

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var sport models.Sport
		err := cursor.Decode(&sport)
		if err != nil {
			return nil, err
		}
		sports = append(sports, &sport)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	if len(sports) == 0 {
		return nil, errors.New("no sports in database")
	}

	return sports, nil
}

func (service *SportServiceImpl) UpdateSport(id *primitive.ObjectID, sport *models.UpdateSport, ctx context.Context) error {
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	update := bson.D{bson.E{Key: "$set", Value: sport}}
	res, err := service.sportCollection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		return errors.New("sport not found")
	}

	return nil
}

func (service *SportServiceImpl) DeleteSport(id *primitive.ObjectID, ctx context.Context) error {
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	res, err := service.sportCollection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	if res.DeletedCount != 1 {
		return errors.New("sport not found")
	}

	return nil
}
