package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mohammadzaidhussain/bookshop/dto"
	"github.com/mohammadzaidhussain/bookshop/service"
	"github.com/mohammadzaidhussain/bookshop/utils"
)

type BookController struct {
	bookService     service.IBookService
	responseService utils.ResponseService
}

func (bc *BookController) CreateBook(ctx *gin.Context) {
	var bookDto dto.CreateBookDto
	if err := ctx.ShouldBindJSON(&bookDto); err != nil {
		ctx.Error(fmt.Errorf("400::%s::%s::%v", "Bad Request", "BookController_CreateBook", err))
	}

	data, err := bc.bookService.CreateBook(bookDto, nil)

	if err != nil {
		ctx.Error(err)
		return
	}

	bc.responseService.Success(ctx, 201, data, "sucessfully saved!")
}

func (bc *BookController) GetAllBooks(ctx *gin.Context) {

	var bookFilterDto dto.GetBooksFilter
	if err := ctx.ShouldBindJSON(&bookFilterDto); err != nil {
		ctx.Error(fmt.Errorf("400::%s::%s::%v", "Bad Request", "BookController_CreateBook", err))
	}

	var shopId = utils.ConvertToObjectId(ctx.GetString("user_id"))

	bookFilterDto.ShopId = shopId

	data, err := bc.bookService.GetPaginatedBooks(bookFilterDto, nil)

	if err != nil {
		ctx.Error(err)
		return
	}

	bc.responseService.Success(ctx, 200, data, "fetched sucessfully!")

}

func (bc *BookController) GetBookById(ctx *gin.Context) {

	var bookId = ctx.Param("bookId")

	if bookId == "" {
		ctx.Error(fmt.Errorf("400::%s::%s::%v", "Bad Request", "BookController_GetBookById", "Book Id is not present is request params"))
	}

	data, err := bc.bookService.GetBookById(bookId, nil)

	if err != nil {
		ctx.Error(err)
		return
	}

	bc.responseService.Success(ctx, 200, data, "fetched sucessfully!")

}

func (bc *BookController) UpdateBookById(ctx *gin.Context) {
	var bookDto map[string]interface{}
	if err := ctx.ShouldBindJSON(&bookDto); err != nil {
		ctx.Error(fmt.Errorf("400::%s::%s::%v", "Bad Request", "BookController_UpdateBookById", err))
	}

	data, err := bc.bookService.UpdateBookById(bookDto, nil)

	if err != nil {
		ctx.Error(err)
		return
	}

	bc.responseService.Success(ctx, 201, data, "updated sucessfully!")
}

func (bc *BookController) DeleteBookById(ctx *gin.Context) {

	var bookId = ctx.Param("bookId")

	if bookId == "" {
		ctx.Error(fmt.Errorf("400::%s::%s::%v", "Bad Request", "BookController_DeleteBookById", "Book Id is not present is request params"))
	}

	data, err := bc.bookService.DeleteBookById(bookId, nil)

	if err != nil {
		ctx.Error(err)
		return
	}

	bc.responseService.Success(ctx, 0, data, "Deleted sucessfully!")

}

func GetBookController(bookService service.IBookService, responseService utils.ResponseService) *BookController {
	return &BookController{
		bookService:     bookService,
		responseService: responseService,
	}
}
