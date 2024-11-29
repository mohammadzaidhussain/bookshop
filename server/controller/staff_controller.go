package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mohammadzaidhussain/bookshop/dto"
	"github.com/mohammadzaidhussain/bookshop/model"
	"github.com/mohammadzaidhussain/bookshop/service"
	"github.com/mohammadzaidhussain/bookshop/utils"
)

type StaffController struct {
	staffService    service.IStaffService
	responseService utils.ResponseService
}

func (bc *StaffController) CreateStaff(ctx *gin.Context) {
	var staffDto dto.StaffCreateDto
	if err := ctx.ShouldBindJSON(&staffDto); err != nil {
		ctx.Error(fmt.Errorf("400::%s::%s::%v", "Bad Request", "StaffController_CreateStaff", err))
	}

	actualStaff := &model.Staff{
		Name:        staffDto.Name,
		Password:    staffDto.Password,
		Email:       staffDto.Email,
		Status:      staffDto.Status,
		Role:        staffDto.Role,
		BookShopId:  utils.ConvertToObjectId(staffDto.BookShopId),
		CreatedDate: staffDto.CreatedDate,
	}

	data, err := bc.staffService.CreateStaff(actualStaff, nil)

	if err != nil {
		ctx.Error(err)
		return
	}

	bc.responseService.Success(ctx, 201, data, "sucessfully saved!")
}

func (bc *StaffController) GetAllStaffs(ctx *gin.Context) {

	var filterDto dto.StaffFilter
	if err := ctx.ShouldBindJSON(&filterDto); err != nil {
		ctx.Error(fmt.Errorf("400::%s::%s::%v", "Bad Request", "StaffController_CreateStaff", err))
	}

	if filterDto.BookshopId != nil {
		filterDto.BookshopId = utils.ConvertToObjectId((filterDto.BookshopId).(string))
	}

	data, err := bc.staffService.GetAllStaffs(filterDto, nil)

	if err != nil {
		ctx.Error(err)
		return
	}

	bc.responseService.Success(ctx, 200, data, "fetched sucessfully!")

}

func (bc *StaffController) GetStaffById(ctx *gin.Context) {

	var staffId = ctx.Param("staffId")

	if staffId == "" {
		ctx.Error(fmt.Errorf("400::%s::%s::%v", "Bad Request", "StaffController_GetStaffById", "Staff Id is not present is request params"))
	}

	data, err := bc.staffService.GetStaffById(staffId, nil)

	if err != nil {
		ctx.Error(err)
		return
	}

	bc.responseService.Success(ctx, 200, data, "fetched sucessfully!")

}

func (bc *StaffController) UpdateStaffById(ctx *gin.Context) {
	var staffDto map[string]interface{}
	if err := ctx.ShouldBindJSON(&staffDto); err != nil {
		ctx.Error(fmt.Errorf("400::%s::%s::%v", "Bad Request", "StaffController_UpdateStaffById", err))
	}

	if staffDto["bookshop_id"] != nil {
		staffDto["bookshop_id"] = utils.ConvertToObjectId(staffDto["bookshop_id"].(string))
	}

	data, err := bc.staffService.UpdateStaffById(staffDto, nil)

	if err != nil {
		ctx.Error(err)
		return
	}

	bc.responseService.Success(ctx, 201, data, "updated sucessfully!")
}

func (bc *StaffController) DeleteStaffById(ctx *gin.Context) {

	var staffId = ctx.Param("staffId")

	if staffId == "" {
		ctx.Error(fmt.Errorf("400::%s::%s::%v", "Bad Request", "StaffController_DeleteStaffById", "Staff Id is not present is request params"))
	}

	data, err := bc.staffService.DeleteStaffById(staffId, nil)

	if err != nil {
		ctx.Error(err)
		return
	}

	bc.responseService.Success(ctx, 0, data, "Deleted sucessfully!")

}

func GetStaffController(staffService service.IStaffService, responseService utils.ResponseService) *StaffController {
	return &StaffController{
		staffService:    staffService,
		responseService: responseService,
	}
}
