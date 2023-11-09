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

type DB struct {
	timeout            time.Duration
	path               string
	dbName             string
	usersCollection    string
	sheltersCollection string
	client             *mongo.Client
}

func NewStorage(cfg config.MongoDBConfig) (*DB, error) {
	const op = "mongo.NewMongoDB"

	timeout := time.Duration(cfg.Timeout) * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	mon := &DB{
		timeout:            timeout,
		path:               cfg.MongoPath,
		dbName:             cfg.DB,
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

func (db DB) SetupMongo() {
	database := db.client.Database(db.dbName)
	collection := database.Collection(db.usersCollection)
	ctx, cancel := context.WithTimeout(context.Background(), db.timeout)
	defer cancel()

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
}
