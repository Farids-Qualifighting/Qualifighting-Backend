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

type CompetitionService interface {
	CreateCompetition(models.Competition, context.Context) error
	GetCompetition(*primitive.ObjectID, context.Context) (*models.Competition, error)
	GetAllCompetitions(context.Context) ([]*models.Competition, error)
	UpdateCompetition(*primitive.ObjectID, *models.UpdateCompetition, context.Context) error
	DeleteCompetition(*primitive.ObjectID, context.Context) error
}

type CompetitionServiceImpl struct {
	competitionCollection *mongo.Collection
}

func NewCompetitionService(competitionCollection *mongo.Collection) CompetitionService {
	return &CompetitionServiceImpl{
		competitionCollection: competitionCollection,
	}
}

func (service *CompetitionServiceImpl) CreateCompetition(competition models.Competition, ctx context.Context) error {

	encryptedCompetition, errEncryption := lib.Encrypt(competition)
	if errEncryption != nil {
		return errEncryption
	}

	// encryptedName, errEncryption := lib.EncryptString(competition.Name, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
	// if errEncryption != nil {
	// 	return errEncryption
	// }

	payload := models.Competition{
		Name:           encryptedCompetition.Name,
		Date:           competition.Date,
		Rank:           competition.Rank,
		WonCompetition: competition.WonCompetition,
		StudentID:      competition.StudentID,
	}

	_, err := service.competitionCollection.InsertOne(ctx, payload)
	return err
}

func (service *CompetitionServiceImpl) GetCompetition(id *primitive.ObjectID, ctx context.Context) (*models.Competition, error) {
	var competition *models.Competition
	query := bson.D{bson.E{Key: "_id", Value: id}}
	err := service.competitionCollection.FindOne(ctx, query).Decode(&competition)
	if err != nil {
		return nil, err
	}
	decryptedName, err := lib.DecryptString(competition.Name, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
	competition.Name = decryptedName
	return competition, err
}

func (service *CompetitionServiceImpl) GetAllCompetitions(ctx context.Context) ([]*models.Competition, error) {
	var competitions []*models.Competition = make([]*models.Competition, 0)
	cursor, err := service.competitionCollection.Find(ctx, bson.D{{}})

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var competition models.Competition
		err := cursor.Decode(&competition)
		if err != nil {
			return nil, err
		}
		decryptedName, err := lib.DecryptString(competition.Name, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
		if err != nil {
			return nil, err
		}
		competition.Name = decryptedName
		competitions = append(competitions, &competition)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	return competitions, nil
}

func (service *CompetitionServiceImpl) UpdateCompetition(id *primitive.ObjectID, competition *models.UpdateCompetition, ctx context.Context) error {
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	update := bson.D{bson.E{Key: "$set", Value: competition}}
	res, err := service.competitionCollection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		return errors.New("competition not found")
	}

	return nil
}

func (service *CompetitionServiceImpl) DeleteCompetition(id *primitive.ObjectID, ctx context.Context) error {
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	res, err := service.competitionCollection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	if res.DeletedCount != 1 {
		return errors.New("competition not found")
	}

	return nil
}
