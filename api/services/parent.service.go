package services

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"qualifighting.backend.de/models"
)

type ParentService interface {
	CreateParent(*models.Parent, context.Context) error
	GetParent(*primitive.ObjectID, context.Context) (*models.Parent, error)
	GetAllParents(context.Context) ([]*models.Parent, error)
	UpdateParent(*primitive.ObjectID, *models.UpdateParent, context.Context) error
	DeleteParent(*primitive.ObjectID, context.Context) error
}

type ParentServiceImpl struct {
	parentCollection *mongo.Collection
}

func NewParentService(parentCollection *mongo.Collection) ParentService {
	return &ParentServiceImpl{
		parentCollection: parentCollection,
	}
}

func (service *ParentServiceImpl) CreateParent(parent *models.Parent, ctx context.Context) error {

	payload := models.Parent{
		FirstName: parent.FirstName,
		LastName:  parent.LastName,
		Children:  parent.Children,
		Address:   parent.Address,
		Phone:     parent.Phone,
		Email:     parent.Email,
	}

	_, err := service.parentCollection.InsertOne(ctx, payload)
	return err
}

func (service *ParentServiceImpl) GetParent(id *primitive.ObjectID, ctx context.Context) (*models.Parent, error) {
	var parent *models.Parent
	query := bson.D{bson.E{Key: "_id", Value: id}}
	err := service.parentCollection.FindOne(ctx, query).Decode(&parent)
	return parent, err
}

func (service *ParentServiceImpl) GetAllParents(ctx context.Context) ([]*models.Parent, error) {
	var parents []*models.Parent
	cursor, err := service.parentCollection.Find(ctx, bson.D{{}})

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var parent models.Parent
		err := cursor.Decode(&parent)
		if err != nil {
			return nil, err
		}
		parents = append(parents, &parent)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	if len(parents) == 0 {
		return nil, errors.New("no parents in database")
	}

	return parents, nil
}

func (service *ParentServiceImpl) UpdateParent(id *primitive.ObjectID, parent *models.UpdateParent, ctx context.Context) error {
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	update := bson.D{bson.E{Key: "$set", Value: parent}}
	res, err := service.parentCollection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		return errors.New("parent not found")
	}

	return nil
}

func (service *ParentServiceImpl) DeleteParent(id *primitive.ObjectID, ctx context.Context) error {
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	res, err := service.parentCollection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	if res.DeletedCount != 1 {
		return errors.New("parent not found")
	}

	return nil
}
