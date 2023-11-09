package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Shelter struct {
	City    string `json:"city" bson:"city"`
	Address string `json:"address" bson:"address"`
}

type Complaint struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	UserID      primitive.ObjectID `json:"user_id" bson:"user_id"`
	ShelterData Shelter            `json:"shelter_data" bson:"shelter_data"`
	Text        string             `json:"text" bson:"text"`
	UploadedAt  time.Time          `json:"uploaded_at" bson:"uploaded_at"`
	IsResolved  bool               `json:"is_resolved" bson:"is_resolved"`
}
