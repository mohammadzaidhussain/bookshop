package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Staff struct {
	Id           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name         string             `bson:"name" json:"name"`
	Email        string             `bson:"email" json:"email"`
	Password     string             `bson:"password" json:"password"`
	Status       bool               `bson:"status" json:"status"`
	Role         string             `bson:"role" json:"role"` // ['manager','staff']
	BookShopId   primitive.ObjectID `bson:"bookshop_id" json:"bookshop_id"`
	ModifiedBy   primitive.ObjectID `bson:"modified_by" json:"modified_by"`
	CreatedDate  time.Time          `bson:"created_date" json:"created_date"`
	ModifiedDate time.Time          `bson:"modified_date" json:"modified_date"`
}
