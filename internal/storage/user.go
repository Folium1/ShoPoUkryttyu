package storage

import (
	"context"
	"fmt"

	"github.com/Folium1/ShoPoUkryttyu/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db DB) CreateUser(user models.User) (primitive.ObjectID, error) {
	const op = "mongo.CreateUser"

	ctx, cancel := context.WithTimeout(context.Background(), db.timeout)
	defer cancel()
	user.ID = primitive.NewObjectID()

	res, err := db.client.Database(db.dbName).Collection(db.usersCollection).InsertOne(ctx, user)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("%s: %w", op, err)
	}

	return res.InsertedID.(primitive.ObjectID), nil
}

func (db DB) UserByID(id primitive.ObjectID) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.timeout)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: id}}
	var user models.User

	err := db.client.Database(db.dbName).Collection(db.usersCollection).FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return models.User{}, fmt.Errorf("mongo.GetUser: %w", err)
	}

	return user, nil
}

func (db DB) AllUsers() ([]models.User, error) {
	const op = "mongo.GetAllUsers"

	ctx, cancel := context.WithTimeout(context.Background(), db.timeout)
	defer cancel()

	filter := bson.D{{}}
	cursor, err := db.client.Database(db.dbName).Collection(db.usersCollection).Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer cursor.Close(ctx)

	var users []models.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return users, nil
}

func (db DB) UserByMail(mail string) (models.User, error) {
	const op = "mongo.LoginUser"

	ctx, cancel := context.WithTimeout(context.Background(), db.timeout)
	defer cancel()

	filter := bson.D{{Key: "email", Value: mail}}

	var userDB models.User
	err := db.client.Database(db.dbName).Collection(db.usersCollection).FindOne(ctx, filter).Decode(&userDB)
	if err != nil {
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return userDB, nil
}
