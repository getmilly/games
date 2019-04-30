package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Game ...
type Game struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	URLImage    string             `bson:"url_image"`
	StartAt     time.Time          `bson:"start_at"`
	EndAt       time.Time          `bson:"end_at"`
	Private     bool               `bson:"private"`
	Active      bool               `bson:"active"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}
