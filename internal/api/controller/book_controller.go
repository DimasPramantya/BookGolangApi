package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/DimasPramantya/goMiniProject/internal/domain"
	"github.com/DimasPramantya/goMiniProject/internal/dto"
	"github.com/DimasPramantya/goMiniProject/utils/helper"
	"github.com/DimasPramantya/goMiniProject/utils/validation"
	"github.com/gin-gonic/gin"
)

type bookController struct {
	bookUC    domain.BookUsecase
	validator *validation.Validator
}

func NewBookController(bookUC domain.BookUsecase, validator *validation.Validator) *bookController {
	return &bookController{
		bookUC:    bookUC,
		validator: validator,
	}
}

// CreateBook godoc
// @Summary     Create Book Entity
// @Description Create a new book
// @Tags        book
// @Param       request body dto.ReqCreateBookSwagger true "Create Book Payload"
// @Produce     json
// @Success     201 {object} dto.ResBook
// @Failure     400 {object} dto.ErrorResponse
// @Security    BearerAuth
// @Router      /books [post]
func (bc *bookController) CreateBook(ctx *gin.Context) {
	username, _ := ctx.Get("username")
	var req dto.ReqCreateBook
	err := bc.validator.ValidateRequest(ctx, &req)
	if err != nil {
		return
	}
	book, err := bc.bookUC.Create(req, username.(string))
	if err != nil {	
		if err == domain.ErrNotFound {
			helper.WriteError(ctx, http.StatusNotFound, fmt.Sprintf("Category with id %d not found!!", *req.CategoryID), nil)
			return
		}
		if err == domain.ErrTotalPageNotValid {
			helper.WriteError(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}
		if err == domain.ErrReleaseYearNotValid {
			helper.WriteError(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}
		helper.WriteError(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	helper.WriteResponse(ctx, http.StatusCreated, "Book created successfully", book)
}

// GetAllBooks godoc
// @Summary     Get All Books
// @Description Retrieves a list of books
// @Tags        book
// @Produce     json
// @Success     200 {array} dto.ResBook
// @Failure     500 {object} dto.ErrorResponse
// @Security    BearerAuth
// @Router      /books [get]
func (bc *bookController) GetAllBooks(ctx *gin.Context) {
	books, err := bc.bookUC.GetAll()
	if err != nil {
		helper.WriteError(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	helper.WriteResponse(ctx, http.StatusOK, "Success", books)
}

// GetBookByID godoc
// @Summary     Get Book By ID
// @Description Retrieves a book by ID
// @Tags        book
// @Param       id path int true "Book ID"
// @Produce     json
// @Success     200 {object} dto.ResBook
// @Failure     400 {object} dto.ErrorResponse
// @Failure     404 {object} dto.ErrorResponse
// @Security    BearerAuth
// @Router      /books/{id} [get]
func (bc *bookController) GetBookByID(ctx *gin.Context) {
	id := ctx.Param("id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		helper.WriteError(ctx, http.StatusBadRequest, "Invalid ID format", nil)
		return
	}

	book, err := bc.bookUC.GetByID(bookID)
	if err != nil {
		if err == domain.ErrNotFound {
			helper.WriteError(ctx, http.StatusNotFound, "Book not found", nil)
			return
		}
		helper.WriteError(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.WriteResponse(ctx, http.StatusOK, "Success", book)
}

// DeleteBook godoc
// @Summary     Delete Book Entity
// @Description Delete a book by ID
// @Tags        book
// @Param       id path int true "Book ID"
// @Produce     json
// @Success     200 {object} dto.BaseResponse
// @Failure     400 {object} dto.ErrorResponse
// @Failure     404 {object} dto.ErrorResponse
// @Security    BearerAuth
// @Router      /books/{id} [delete]
func (bc *bookController) DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		helper.WriteError(ctx, http.StatusBadRequest, "Invalid ID format", nil)
		return
	}

	err = bc.bookUC.Delete(bookID)
	if err != nil {
		if err == domain.ErrNotFound {
			helper.WriteError(ctx, http.StatusNotFound, "Book not found", nil)
			return
		}
		helper.WriteError(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.WriteResponse(ctx, http.StatusOK, "Book deleted successfully", nil)
}