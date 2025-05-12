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

type categoryController struct {
	categoryUC domain.CategoryUsecase
	validator *validation.Validator
}

func NewCategoryController(categoryUC domain.CategoryUsecase, validator *validation.Validator) *categoryController {
	return &categoryController{
		categoryUC: categoryUC,
		validator: validator,
	}
}

// CreateCategory godoc
// @Summary     Create Category Entity
// @Description Create a new category
// @Tags        category
// @Param       request body dto.CreateCategoryRequest true "Create Category Payload"
// @Produce     json
// @Success     201 {object} dto.ResCategory
// @Failure     400 {object} dto.ErrorResponse
// @Security    BearerAuth
// @Router      /categories [post]
func (cc *categoryController) CreateCategory(ctx *gin.Context) {
	username, _ := ctx.Get("username")
	var req dto.CreateCategoryRequest
	err := cc.validator.ValidateRequest(ctx, &req)
	if err != nil {
		return
	}

	category, err := cc.categoryUC.Create(req, username.(string))
	if err != nil {
		helper.WriteError(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.WriteResponse(ctx, http.StatusCreated, "Category created successfully", category)
}

// GetAllCategory godoc
// @Summary     Get All Categories
// @Description Retrieves a list of categories
// @Tags        category
// @Produce     json
// @Success     200 {array} dto.ResCategory
// @Failure     500 {object} dto.ErrorResponse
// @Security    BearerAuth
// @Router      /categories [get]
func (cc *categoryController) GetAllCategory(ctx *gin.Context) {
	categories, err := cc.categoryUC.FindAll()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	helper.WriteResponse(ctx, http.StatusOK, "Success", categories)
}

// GetCategoryById godoc
// @Summary     Get Category By ID
// @Description Retrieves a category by ID
// @Tags        category
// @Param       id path int true "Category ID"
// @Produce     json
// @Success     200 {object} dto.ResCategory
// @Failure     400 {object} dto.ErrorResponse
// @Security    BearerAuth
// @Router      /categories/{id} [get]
func (cc *categoryController) GetCategoryByID(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		helper.WriteError(ctx, http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	category, err := cc.categoryUC.FindByID(intID)
	if err != nil {
		if err == domain.ErrNotFound {
			helper.WriteError(ctx, http.StatusNotFound, fmt.Sprintf("Category with id %d not found!!", intID), nil)
			return
		}
		helper.WriteError(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.WriteResponse(ctx, http.StatusOK, "Success", category)
}

// DeleteCategory godoc
// @Summary     Delete Category Entity
// @Description Delete a category by ID
// @Tags        category
// @Param       id path int true "Category ID"
// @Produce     json
// @Success     200 {object} dto.BaseResponse
// @Failure     400 {object} dto.ErrorResponse
// @Security    BearerAuth
// @Router      /categories/{id} [delete]
func (cc *categoryController) DeleteCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		helper.WriteError(ctx, http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	err = cc.categoryUC.Delete(intID)
	if err != nil {
		if err == domain.ErrNotFound {
			helper.WriteError(ctx, http.StatusNotFound, fmt.Sprintf("Category with id %d not found!!", intID), nil)
			return
		}
		helper.WriteError(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.WriteResponse(ctx, http.StatusOK, "Category deleted successfully", nil)
}

// UpdateCategory godoc
// @Summary     Update Category Entity
// @Description Update category
// @Tags        category
// @Param       id path int true "Category ID"
// @Param       request body dto.UpdateCategoryRequest true "Update Category Payload"
// @Produce     json
// @Success     200 {object} dto.ResCategory
// @Failure     400 {object} dto.ErrorResponse
// @Security    BearerAuth
// @Router      /categories/{id} [put]
func (cc *categoryController) UpdateCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	username, _ := ctx.Get("username")
	intID, err := strconv.Atoi(id)
	if err != nil {
		helper.WriteError(ctx, http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	var req dto.UpdateCategoryRequest
	err = cc.validator.ValidateRequest(ctx, &req)
	if err != nil {
		return
	}

	category, err := cc.categoryUC.Update(req, intID, username.(string))
	if err != nil {
		if err == domain.ErrNotFound {
			helper.WriteError(ctx, http.StatusNotFound, fmt.Sprintf("Category with id %d not found!!", intID), nil)
			return
		}
		helper.WriteError(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.WriteResponse(ctx, http.StatusOK, "Category updated successfully", category)
}


// GetBooksByCategoryID godoc
// @Summary     Get books by category ID
// @Description Retrieves a list of books in a specific category
// @Tags        book
// @Param       id path int true "Category ID"
// @Produce     json
// @Success     200 {array} dto.ResBook
// @Failure     400 {object} dto.ErrorResponse
// @Security    BearerAuth
// @Router      /categories/{id}/books [get]
func (cc *categoryController) GetAllBooksByCategoryID(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		helper.WriteError(ctx, http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	books, err := cc.categoryUC.FindAllBooksByCategoryID(intID)
	if err != nil {
		if err == domain.ErrNotFound {
			helper.WriteError(ctx, http.StatusNotFound, fmt.Sprintf("Category with id %d not found!!", intID), nil)
			return
		}
		helper.WriteError(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.WriteResponse(ctx, http.StatusOK, "Success", books)
}

