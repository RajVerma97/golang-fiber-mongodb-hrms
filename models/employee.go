package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)



type Employee struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FullName  string             `json:"fullName" bson:"fullName"`
	Email     string             `json:"email" bson:"email"`
	Position  string             `json:"position" bson:"position"`
	Salary    float64            `json:"salary" bson:"salary"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}
