package service

import (
	"github.com/mohammadzaidhussain/bookshop/dto"
	"github.com/mohammadzaidhussain/bookshop/model"
	repo "github.com/mohammadzaidhussain/bookshop/repository"
	"github.com/mohammadzaidhussain/bookshop/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type IStaffService interface {
	CreateStaff(data interface{}, sessionContext mongo.SessionContext) (interface{}, error)
	GetAllStaffs(filter dto.StaffFilter, sessionContext mongo.SessionContext) (interface{}, error)
	GetStaffById(id string, sessionContext mongo.SessionContext) (interface{}, error)
	UpdateStaffById(data map[string]interface{}, sessionContext mongo.SessionContext) (interface{}, error)
	DeleteStaffById(id string, sessionContext mongo.SessionContext) (interface{}, error)
	FindOneStaffByEmail(email string, sessionContext mongo.SessionContext) (interface{}, error)
}

type StaffService struct {
	repository repo.Repository
}

func (bs *StaffService) CreateStaff(data interface{}, sessionContext mongo.SessionContext) (interface{}, error) {
	return bs.repository.StaffRepository.Create(data, sessionContext)
}

func (bs *StaffService) GetAllStaffs(filter dto.StaffFilter, sessionContext mongo.SessionContext) (interface{}, error) {
	return bs.repository.StaffRepository.FindAll(filter, sessionContext)
}

func (us *StaffService) FindOneStaffByEmail(email string, sessionContext mongo.SessionContext) (interface{}, error) {
	res, err := us.repository.StaffRepository.FindOneByKey("email", email, sessionContext)
	if err != nil {
		return nil, err
	}
	var staff model.Staff
	if err := utils.MapToStruct(res.(map[string]interface{}), &staff); err != nil {
		return nil, err
	}
	return staff, nil
}

func (bs *StaffService) GetStaffById(id string, sessionContext mongo.SessionContext) (interface{}, error) {
	return bs.repository.StaffRepository.FindOne(id, sessionContext)
}

func (bs *StaffService) UpdateStaffById(data map[string]interface{}, sessionContext mongo.SessionContext) (interface{}, error) {
	bookData := data
	id := bookData["_id"]
	delete(bookData, "_id")
	return bs.repository.StaffRepository.Update(id.(string), bookData, sessionContext)
}

func (bs *StaffService) DeleteStaffById(id string, sessionContext mongo.SessionContext) (interface{}, error) {
	return bs.repository.StaffRepository.Delete(id, sessionContext)
}

func GetStaffService(respository repo.Repository) IStaffService {
	return &StaffService{
		repository: respository,
	}
}
