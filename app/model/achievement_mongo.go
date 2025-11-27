package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Achievement struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Category    string             `bson:"category"`
	Level       string             `bson:"level"` 
	Organizer   string             `bson:"organizer"`
	Year        int                `bson:"year"`
	StudentName string             `bson:"student_name"`
	FileURL     string             `bson:"file_url"`
	CreatedAt   int64              `bson:"created_at"`
	UpdatedAt   int64              `bson:"updated_at"`
}
