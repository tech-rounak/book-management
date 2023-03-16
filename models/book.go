package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      *string            `bson:"name" validate:"required"`
	Publisher *string            `bson:"publisher" validate:"required"`
	Price     *int64             `bson:"price" validate:"required"`
	Author    *string            `bson:"author" validate:"required"`
	ISBN      *string            `bson:"isbn" validate:"required,min=12,max=12"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}
