package storage

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Folium1/ShoPoUkryttyu/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	timeout = 10 * time.Second
)

type DB struct {
	path               string
	db                 string
	usersCollection    string
	sheltersCollection string
	client             *mongo.Client
}

func NewStorage(cfg config.MongoDBConfig) (*DB, error) {
	const op = "mongo.NewMongoDB"

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	mon := &DB{
		path:               cfg.MongoPath,
		db:                 cfg.DB,
		usersCollection:    cfg.UsersCollection,
		sheltersCollection: cfg.SheltersCollection,
	}
	var err error

	mon.client, err = mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoPath))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return mon, nil
}

// SetupMongo creates unique indexes on the "email" and "address" fields
func (db DB) SetupMongo() {

	// Get the database and collection objects
	database := db.client.Database(db.db)
	collection := database.Collection(db.usersCollection)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Create a unique index on the "email" field
	indexModel := mongo.IndexModel{
		Keys: bson.M{
			"email": 1,
		},
		Options: options.Index().SetUnique(true),
	}
	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		log.Fatalf("Failed to create index: %v", err)
	}

	// Get the database and collection objects
	database = db.client.Database(db.db)
	collection = database.Collection(db.sheltersCollection)
	ctx, cancel = context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Create a unique index on the "email" field
	indexModel = mongo.IndexModel{
		Keys: bson.M{
			"address": 1,
		},
		Options: options.Index().SetUnique(true),
	}
	_, err = collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		log.Fatalf("Failed to create index: %v", err)
	}
}
