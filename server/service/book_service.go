package service

import (
	"fmt"

	"github.com/mohammadzaidhussain/bookshop/dto"
	repo "github.com/mohammadzaidhussain/bookshop/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IBookService interface {
	CreateBook(data interface{}, sessionContext mongo.SessionContext) (interface{}, error)
	GetAllBooks(sessionContext mongo.SessionContext) (interface{}, error)
	GetPaginatedBooks(filterDto dto.GetBooksFilter, sessionContext mongo.SessionContext) (interface{}, error)
	GetBookById(id string, sessionContext mongo.SessionContext) (interface{}, error)
	UpdateBookById(data map[string]interface{}, sessionContext mongo.SessionContext) (interface{}, error)
	DeleteBookById(id string, sessionContext mongo.SessionContext) (interface{}, error)
}

type BookService struct {
	repository repo.Repository
}

func (bs *BookService) CreateBook(data interface{}, sessionContext mongo.SessionContext) (interface{}, error) {
	return bs.repository.BookRepository.Create(data, sessionContext)
}

func (bs *BookService) GetAllBooks(sessionContext mongo.SessionContext) (interface{}, error) {
	return bs.repository.BookRepository.FindAll(nil, sessionContext)
}

func (bs *BookService) GetPaginatedBooks(filterDto dto.GetBooksFilter, sessionContext mongo.SessionContext) (interface{}, error) {

	filtersA := bson.M{
		"$or": bson.A{
			bson.M{"is_deleted": bson.M{"$exists": false}},
			bson.M{"is_deleted": false},
		},
	}

	filtersB := bson.M{
		"$or": bson.A{
			bson.M{"book_name": bson.M{"$regex": filterDto.BookSearch, "$options": "i"}},
			bson.M{"book_code": bson.M{"$regex": filterDto.BookSearch, "$options": "i"}},
		},
	}

	filters := bson.M{
		"$and": bson.A{
			filtersA,
			// bson.M{"created_by": filterDto.ShopId},
		},
	}

	if filterDto.BookSearch != "" {
		filters["$and"] = append(filters["$and"].(bson.A), filtersB)
	}

	page := filterDto.PageNo

	if page <= 0 {
		page = 1
	}

	limit := filterDto.Limit

	if limit <= 0 {
		limit = 10
	}

	skip := (page - 1) * limit

	pipelines := bson.A{

		// match stage
		bson.M{"$match": filters},

		// join with staff
		bson.M{
			"$lookup": bson.M{
				"from":         "staff",
				"localField":   "modified_by",
				"foreignField": "_id",
				"as":           "staffs",
			},
		},

		// take out a single item from the staffs array from specific index
		bson.M{
			"$addFields": bson.M{
				"staff": bson.M{
					"$arrayElemAt": bson.A{"$staffs", 0},
				},
			},
		},

		// take out only required data
		bson.M{
			"$project": bson.M{
				"_id":            1,
				"book_name":      1,
				"book_code":      1,
				"description":    1,
				"author":         1,
				"book_price":     1,
				"published_year": 1,
				"status":         1,
				"page_count":     1,
				"created_date":   1,
				"modified_date":  1,
				"staff":          1,
			},
		},

		// fetch total records
		bson.M{
			"$facet": bson.M{
				"data": bson.A{
					bson.M{"$skip": skip},
					bson.M{"$limit": limit},
				},
				"totalRecords": bson.A{
					bson.M{"$count": "count"},
				},
			},
		},

		// assign new key with the value of data
		bson.M{
			"$addFields": bson.M{
				"items": "$data",
			},
		},

		// format data in the expected format
		bson.M{
			"$project": bson.M{
				"items": 1,
				"total_records": bson.M{
					"$arrayElemAt": bson.A{"$totalRecords.count", 0},
				},
			},
		},
	}

	data, err := bs.repository.BookRepository.Aggregate(pipelines, sessionContext)

	if err != nil {
		return nil, fmt.Errorf("500::%s::%s::%v", "Internal server erorr", "BookService_GetPaginatedBooks", err)
	}

	return data[0], nil

}

func (bs *BookService) GetBookById(id string, sessionContext mongo.SessionContext) (interface{}, error) {
	return bs.repository.BookRepository.FindOne(id, sessionContext)
}

func (bs *BookService) UpdateBookById(data map[string]interface{}, sessionContext mongo.SessionContext) (interface{}, error) {
	bookData := data
	id := bookData["_id"]
	delete(bookData, "_id")
	return bs.repository.BookRepository.Update(id.(string), bookData, sessionContext)
}

func (bs *BookService) DeleteBookById(id string, sessionContext mongo.SessionContext) (interface{}, error) {
	return bs.repository.BookRepository.Delete(id, sessionContext)
}

func GetBookService(respository repo.Repository) IBookService {
	return &BookService{
		repository: respository,
	}
}
