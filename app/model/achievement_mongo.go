package model

import (
    "time"

    "go.mongodb.org/mongo-driver/bson/primitive"
)

// Achievement represents a student's achievement stored in MongoDB.
// @swagger:model Achievement
type Achievement struct {
    ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Title           string             `bson:"title" json:"title"`
    Description     string             `bson:"description" json:"description"`
    Category        string             `bson:"category" json:"category"`
    Level           string             `bson:"level" json:"level"`
    Organizer       string             `bson:"organizer" json:"organizer"`
    Year            int                `bson:"year" json:"year"`
    StudentName     string             `bson:"student_name" json:"student_name"`
    StudentID       string             `bson:"student_id" json:"student_id"`
    FileURL         string             `bson:"file_url" json:"file_url"`

    Status          string             `bson:"status" json:"status"`                       // draft | submitted | verified | rejected
    RejectionNote   *string            `bson:"rejection_note,omitempty" json:"rejection_note,omitempty"`
    VerifiedBy      *string            `bson:"verified_by,omitempty" json:"verified_by,omitempty"`
    VerifiedAt      *time.Time         `bson:"verified_at,omitempty" json:"verified_at,omitempty"`

    CreatedAt       time.Time          `bson:"created_at" json:"created_at"`
    UpdatedAt       time.Time          `bson:"updated_at" json:"updated_at"`
}
