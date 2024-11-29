package repository

import "github.com/mohammadzaidhussain/bookshop/config"

type Repository struct {
	UserRepository  IMongoRepository
	BookRepository  IMongoRepository
	StaffRepository IMongoRepository
}

func GetRepository() *Repository {
	dbName := config.GetEnvProperty("database_name")
	return &Repository{
		UserRepository:  GetMongoRepository(dbName, "user"),
		BookRepository:  GetMongoRepository(dbName, "book"),
		StaffRepository: GetMongoRepository(dbName, "staff"),
	}
}
