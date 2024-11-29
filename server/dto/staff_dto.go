package dto

import "time"

type StaffCreateDto struct {
	Name        string    `bson:"name" json:"name"`
	Email       string    `bson:"email" json:"email"`
	Password    string    `bson:"password" json:"password"`
	Status      bool      `bson:"status" json:"status"`
	Role        string    `bson:"role" json:"role"` // ['manager','staff']
	BookShopId  string    `bson:"bookshop_id" json:"bookshop_id"`
	CreatedDate time.Time `bson:"created_date" json:"created_date"`
}

type StaffFilter struct {
	BookshopId interface{} `json:"bookshop_id"`
}
