package storage

import (
	"context"
	"fmt"

	"github.com/Folium1/ShoPoUkryttyu/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db DB) AddComplaint(complaint models.Complaint) (primitive.ObjectID, error) {
	const op = "mongo.AddComplaint"

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	complaint.ID = primitive.NewObjectID()
	filter := bson.D{{Key: "complaint", Value: complaint}}

	res, err := db.client.Database(db.dbName).Collection(db.sheltersCollection).InsertOne(ctx, filter)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("%s: %w", op, err)
	}

	return res.InsertedID.(primitive.ObjectID), nil
}

func (db DB) GetComplaintByID(ID primitive.ObjectID) (models.Complaint, error) {
	const op = "mongo.GetComplaint"

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: ID}}

	cursor, err := db.client.Database(db.dbName).Collection(db.sheltersCollection).Find(ctx, filter)
	if err != nil {
		return models.Complaint{}, fmt.Errorf("%s: %w", op, err)
	}
	defer cursor.Close(ctx)

	var complaint models.Complaint
	err = cursor.Decode(&complaint)
	if err != nil {
		return models.Complaint{}, fmt.Errorf("%s: %w", op, err)
	}

	return complaint, nil
}

func (db DB) GetComplaintByAddress(shelter models.Shelter) ([]models.Complaint, error) {
	op := "mongo.GetComplaint"

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	filter := bson.D{{Key: "city", Value: shelter.City}, {Key: "address", Value: shelter.Address}}
	cursor, err := db.client.Database(db.dbName).Collection(db.sheltersCollection).Find(ctx, filter)
	if err != nil {
		return []models.Complaint{}, fmt.Errorf("%s: %w", op, err)
	}
	defer cursor.Close(ctx)

	var complaints []models.Complaint
	if err = cursor.All(ctx, &complaints); err != nil {
		return []models.Complaint{}, fmt.Errorf("%s: %w", op, err)
	}

	return complaints, nil
}

func (db DB) GetComplaintByUserID(ID primitive.ObjectID) ([]models.Complaint, error) {
	op := "mongo.GetComplaintByUserID"

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	filter := bson.D{{Key: "userID", Value: ID}}
	cursor, err := db.client.Database(db.dbName).Collection(db.sheltersCollection).Find(ctx, filter)
	if err != nil {
		return []models.Complaint{}, fmt.Errorf("%s: %w", op, err)
	}
	defer cursor.Close(ctx)

	var complaints []models.Complaint
	if err = cursor.All(ctx, &complaints); err != nil {
		return []models.Complaint{}, fmt.Errorf("%s: %w", op, err)
	}

	return complaints, nil
}

func (db DB) GetAllComplaints() ([]models.Complaint, error) {
	const op = "mongo.GetAllComplaints"

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	filter := bson.D{{}}

	cursor, err := db.client.Database(db.dbName).Collection(db.sheltersCollection).Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var complaints []models.Complaint
	if err = cursor.All(ctx, &complaints); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return complaints, nil
}
