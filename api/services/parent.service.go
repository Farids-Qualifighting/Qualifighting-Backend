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

type ParentService interface {
	CreateParent(models.Parent, context.Context) error
	GetParent(*primitive.ObjectID, context.Context) (*models.Parent, error)
	GetAllParents(context.Context) ([]*models.Parent, error)
	UpdateParent(*primitive.ObjectID, models.UpdateParent, context.Context) error
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

func (service *ParentServiceImpl) CreateParent(parent models.Parent, ctx context.Context) error {

	encryptedParent, errEncryption := lib.Encrypt(parent)
	if errEncryption != nil {
		return errEncryption
	}

	payload := models.Parent{
		FirstName: encryptedParent.FirstName,
		LastName:  encryptedParent.LastName,
		Children:  encryptedParent.Children,
		Address:   encryptedParent.Address,
		Phone:     encryptedParent.Phone,
		Email:     encryptedParent.Email,
	}

	_, err := service.parentCollection.InsertOne(ctx, payload)
	return err
}

func (service *ParentServiceImpl) GetParent(id *primitive.ObjectID, ctx context.Context) (*models.Parent, error) {
	var parent *models.Parent
	query := bson.D{bson.E{Key: "_id", Value: id}}
	err := service.parentCollection.FindOne(ctx, query).Decode(&parent)
	if err != nil {
		return nil, err
	}

	decryptedFirstName, err := lib.DecryptString(parent.FirstName, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
	if err != nil {
		return nil, err
	}
	parent.FirstName = decryptedFirstName

	decryptedLastName, err := lib.DecryptString(parent.LastName, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
	if err != nil {
		return nil, err
	}
	parent.LastName = decryptedLastName

	decryptedPhone, err := lib.DecryptString(parent.Phone, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
	if err != nil {
		return nil, err
	}
	parent.Phone = decryptedPhone

	decryptedEmail, err := lib.DecryptString(parent.Email, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
	if err != nil {
		return nil, err
	}
	parent.Email = decryptedEmail

	return parent, err
}

func (service *ParentServiceImpl) GetAllParents(ctx context.Context) ([]*models.Parent, error) {
	var parents []*models.Parent = make([]*models.Parent, 0)
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

		decryptedFirstName, err := lib.DecryptString(parent.FirstName, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
		if err != nil {
			return nil, err
		}
		parent.FirstName = decryptedFirstName

		decryptedLastName, err := lib.DecryptString(parent.LastName, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
		if err != nil {
			return nil, err
		}
		parent.LastName = decryptedLastName

		decryptedPhone, err := lib.DecryptString(parent.Phone, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
		if err != nil {
			return nil, err
		}
		parent.Phone = decryptedPhone

		decryptedEmail, err := lib.DecryptString(parent.Email, "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u")
		if err != nil {
			return nil, err
		}
		parent.Email = decryptedEmail

		parents = append(parents, &parent)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	return parents, nil
}

func (service *ParentServiceImpl) UpdateParent(id *primitive.ObjectID, parent models.UpdateParent, ctx context.Context) error {

	encryptedParent, errEncryption := lib.Encrypt(parent)
	if errEncryption != nil {
		return errEncryption
	}

	filter := bson.D{bson.E{Key: "_id", Value: id}}
	update := bson.D{bson.E{Key: "$set", Value: encryptedParent}}
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
