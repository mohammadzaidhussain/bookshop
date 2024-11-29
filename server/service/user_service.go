package service

import (
	"github.com/mohammadzaidhussain/bookshop/model"
	repo "github.com/mohammadzaidhussain/bookshop/repository"
	"github.com/mohammadzaidhussain/bookshop/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserService interface {
	CreateUser(data interface{}, sessionContext mongo.SessionContext) (interface{}, error)
	FindOneUserByEmail(email string, sessionContext mongo.SessionContext) (interface{}, error)
}

type UserService struct {
	repository repo.Repository
}

func (us *UserService) CreateUser(data interface{}, sessionContext mongo.SessionContext) (interface{}, error) {
	return us.repository.UserRepository.Create(data, sessionContext)
}

func (us *UserService) FindOneUserByEmail(email string, sessionContext mongo.SessionContext) (interface{}, error) {
	res, err := us.repository.UserRepository.FindOneByKey("email", email, sessionContext)
	if err != nil {
		return nil, err
	}
	var user model.User
	if err := utils.MapToStruct(res.(map[string]interface{}), &user); err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserService(repository repo.Repository) IUserService {
	return &UserService{
		repository: repository,
	}
}
