package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Shelter struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	City       string             `json:"city" bson:"city"`
	Address    string             `json:"address" bson:"address"`
	Rating     float32            `json:"rating" bson:"rating"`
	Complaints []Complaint        `json:"complaints" bson:"complaints"`
}

type Complaint struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	UserID     primitive.ObjectID `json:"user_id" bson:"user_id"`
	Text       string             `json:"text" bson:"text"`
	UploadedAt time.Time          `json:"uploaded_at" bson:"uploaded_at"`
	IsResolved bool               `json:"is_resolved" bson:"is_resolved"`
}
