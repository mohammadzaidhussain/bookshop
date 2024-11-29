package service

import (
	"fmt"

	"github.com/mohammadzaidhussain/bookshop/dto"
	"github.com/mohammadzaidhussain/bookshop/model"
	repo "github.com/mohammadzaidhussain/bookshop/repository"
	"github.com/mohammadzaidhussain/bookshop/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type IAuthService interface {
	Login(loginDto dto.LoginDto, sessionContext mongo.SessionContext) (map[string]interface{}, error)
	LoginStaff(loginDto dto.LoginDto, sessionContext mongo.SessionContext) (map[string]interface{}, error)
}

type AuthService struct {
	repository   repo.Repository
	userService  IUserService
	staffService IStaffService
}

func (as *AuthService) Login(loginDto dto.LoginDto, sessionContext mongo.SessionContext) (map[string]interface{}, error) {
	res, err := as.userService.FindOneUserByEmail(loginDto.Email, sessionContext)
	if err != nil {
		return nil, err
	}
	user := res.(model.User)
	token, err := utils.CreateToken(map[string]any{
		"user_id": user.Id.Hex(),
		"role":    "owner",
	})
	if err != nil {
		return nil, fmt.Errorf("500::%s::%s::%v", "Internal server erorr", "UserService_Login", err)
	}
	if user.Password == loginDto.Password {
		return map[string]any{
			"user":  user,
			"token": token,
			"role":  "owner",
		}, nil
	}
	return nil, fmt.Errorf("401::%s::%s::%v", "Invalid credentials", "UserService_Login", "Password Mismatch")
}

func (as *AuthService) LoginStaff(loginDto dto.LoginDto, sessionContext mongo.SessionContext) (map[string]interface{}, error) {
	res, err := as.staffService.FindOneStaffByEmail(loginDto.Email, sessionContext)
	if err != nil {
		return nil, err
	}
	staff := res.(model.Staff)
	token, err := utils.CreateToken(map[string]any{
		"user_id": staff.Id.Hex(),
		"role":    staff.Role,
	})
	if err != nil {
		return nil, fmt.Errorf("500::%s::%s::%v", "Internal server erorr", "UserService_Login", err)
	}
	if staff.Password == loginDto.Password {
		return map[string]any{
			"user":  staff,
			"token": token,
			"role":  staff.Role,
		}, nil
	}
	return nil, fmt.Errorf("401::%s::%s::%v", "Invalid credentials", "UserService_Login", "Password Mismatch")
}

func GetAuthService(repository repo.Repository, userService IUserService, staffService IStaffService) IAuthService {
	return &AuthService{
		repository:   repository,
		userService:  userService,
		staffService: staffService,
	}
}
