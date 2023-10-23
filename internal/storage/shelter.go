package storage

import (
	"context"
	"fmt"

	"github.com/Folium1/ShoPoUkryttyu/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (db DB) CreateShelter(shelter models.Shelter) (primitive.ObjectID, error) {
	const op = "mongo.CreateShelter"

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res, err := db.client.Database(db.db).Collection(db.sheltersCollection).InsertOne(ctx, shelter)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("%s: %w", op, err)
	}

	return res.InsertedID.(primitive.ObjectID), nil
}

func (db DB) AddComplaint(citiID primitive.ObjectID, complaint models.Complaint) error {
	const op = "mongo.AddComplaint"

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: citiID}, {Key: "$push", Value: bson.D{{Key: "complaints", Value: complaint}}}}

	_, err := db.client.Database(db.db).Collection(db.sheltersCollection).InsertOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (db DB) GetShelter(id primitive.ObjectID) (models.Shelter, error) {
	const op = "mongo.GetShelter"

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: id}}

	var shelter models.Shelter
	err := db.client.Database(db.db).Collection(db.sheltersCollection).FindOne(ctx, filter).Decode(&shelter)
	if err != nil {
		return models.Shelter{}, fmt.Errorf("%s: %w", op, err)
	}

	return shelter, nil
}

func (db DB) GetAllShelters() ([]models.Shelter, error) {
	const op = "mongo.GetAllShelters"

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	filter := bson.D{{}, {}}
	opts := options.Find()
	opts.SetSort(bson.D{{Key: "name", Value: 1}})

	cursor, err := db.client.Database(db.db).Collection(db.sheltersCollection).Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer cursor.Close(ctx)

	var shelters []models.Shelter
	if err = cursor.All(ctx, &shelters); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return shelters, nil
}
