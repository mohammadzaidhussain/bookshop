package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateBookDto struct {
	BookName      string    `bson:"book_name" json:"book_name"`
	BookCode      string    `bson:"book_code" json:"book_code"`
	Description   string    `bson:"description,omitempty" json:"description,emitempty"`
	Author        string    `bson:"author" json:"author"`
	BookPrice     float32   `bson:"book_price" json:"book_price"`
	PageCount     int       `bson:"page_count" json:"page_count"`
	PublishedYear int       `bson:"published_year" json:"published_year"`
	Status        bool      `bson:"status" json:"status"`
	CreatedBy     string    `bson:"created_by" json:"created_by"`
	ModifiedBy    string    `bson:"modified_by" json:"modified_by"`
	CreatedDate   time.Time `bson:"created_date" json:"created_date"`
	ModifiedDate  time.Time `bson:"modified_date" json:"modified_date"`
}

type UpdateBookDto struct {
	Id           string    `bson:"_id" json:"_id"`
	BookName     string    `bson:"book_name" json:"book_name"`
	BookCode     string    `bson:"book_code" json:"book_code"`
	Description  string    `bson:"description,omitempty" json:"description,emitempty"`
	BookPrice    float32   `bson:"book_price" json:"book_price"`
	Status       bool      `bson:"status" json:"status"`
	CreatedBy    string    `bson:"created_by" json:"created_by"`
	ModifiedBy   string    `bson:"modified_by" json:"modified_by"`
	ModifiedDate time.Time `bson:"modified_date" json:"modified_date"`
}

type GetBooksFilter struct {
	PageNo     int                `json:"page_no"`
	Limit      int                `json:"limit"`
	BookSearch string             `json:"book_search"`
	ShopId     primitive.ObjectID `json:"shop_id"`
}
