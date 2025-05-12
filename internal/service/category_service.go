package service

import (
	"github.com/DimasPramantya/goMiniProject/internal/domain"
	"github.com/DimasPramantya/goMiniProject/internal/dto"
	"github.com/DimasPramantya/goMiniProject/utils/helper"
)

type categoryService struct {
	categoryRepo domain.CategoryRepository
	bookRepo    domain.BookRepository
}

func NewCategoryService(
	categoryRepo domain.CategoryRepository, 
	bookRepo domain.BookRepository,
) domain.CategoryUsecase {
	return &categoryService{
		categoryRepo: categoryRepo,
		bookRepo:   bookRepo,
	}
}

func (c *categoryService) FindAllBooksByCategoryID(categoryID int) ([]dto.ResBook, error) {
	category , err := c.categoryRepo.FindByID(categoryID)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, domain.ErrNotFound
	}
	books, err := c.bookRepo.FindAllByCategoryID(categoryID)
	if err != nil {
		return nil, err
	}
	var resBooks []dto.ResBook
	for _, book := range books {
		resBook := dto.ResBook{
			ID:          book.ID,
			Title:       book.Title,
			Description: book.Description,
			ImageURL:    book.ImageURL,
			ReleaseYear: book.ReleaseYear,
			Price:       book.Price,
			TotalPage:   book.TotalPage,
			Thickness:   book.Thickness,
			CategoryID:  book.CategoryID,
			CreatedAt:   *helper.TimeToString(&book.CreatedAt),
			CreatedBy:   book.CreatedBy,
			ModifiedAt:  helper.TimeToString(book.ModifiedAt),
			ModifiedBy:  book.ModifiedBy,
		}
		resBooks = append(resBooks, resBook)
	}
	return resBooks, nil
}

func (c *categoryService) Create(req dto.CreateCategoryRequest, createdBy string) (dto.ResCategory, error) {
	category := domain.Category{
		Name: 	 req.Name,
		CreatedBy: createdBy,
	}
	newCategory, err := c.categoryRepo.Create(category)
	if err != nil {
		return dto.ResCategory{}, err
	}
	return dto.ResCategory{
		ID:        newCategory.ID,
		Name: 	newCategory.Name,
		CreatedAt: *helper.TimeToString(&newCategory.CreatedAt),
		CreatedBy: newCategory.CreatedBy,
		ModifiedAt: helper.TimeToString(newCategory.ModifiedAt),
		ModifiedBy: newCategory.ModifiedBy,
	}, nil
}


func (c *categoryService) Delete(id int) error {
	category , err := c.categoryRepo.FindByID(id)
	if err != nil {
		return err
	}
	if category == nil {
		return domain.ErrNotFound
	}
	err = c.categoryRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (c *categoryService) FindAll() ([]dto.ResCategory, error) {
	categories, err := c.categoryRepo.FindAll()
	if err != nil {
		return nil, err
	}
	var resCategories []dto.ResCategory
	for _, category := range categories {
		resCategory := dto.ResCategory{
			ID:        category.ID,
			Name:      category.Name,
			CreatedAt: *helper.TimeToString(&category.CreatedAt),
			CreatedBy: category.CreatedBy,
			ModifiedAt: helper.TimeToString(category.ModifiedAt),
			ModifiedBy: category.ModifiedBy,
		}
		resCategories = append(resCategories, resCategory)
	}
	return resCategories, nil
}

func (c *categoryService) FindByID(id int) (*dto.ResCategory, error) {
	category, err := c.categoryRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, domain.ErrNotFound
	}
	resCategory := dto.ResCategory{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: *helper.TimeToString(&category.CreatedAt),
		CreatedBy: category.CreatedBy,
		ModifiedAt: helper.TimeToString(category.ModifiedAt),
		ModifiedBy: category.ModifiedBy,
	}
	return &resCategory, nil
}

func (c *categoryService) Update(req dto.UpdateCategoryRequest, id int, modifiedBy string) (dto.ResCategory, error) {
	category, err := c.categoryRepo.FindByID(id)
	if err != nil {
		return dto.ResCategory{}, err
	}
	if category == nil {
		return dto.ResCategory{}, domain.ErrNotFound
	}
	category.Name = req.Name
	category.ModifiedBy = &modifiedBy
	updatedCategory, err := c.categoryRepo.Update(*category)
	if err != nil {
		return dto.ResCategory{}, err
	}
	return dto.ResCategory{
		ID:        updatedCategory.ID,
		Name:      updatedCategory.Name,
		CreatedAt: *helper.TimeToString(&updatedCategory.CreatedAt),
		CreatedBy: updatedCategory.CreatedBy,
		ModifiedAt: helper.TimeToString(updatedCategory.ModifiedAt),
		ModifiedBy: updatedCategory.ModifiedBy,
	}, nil
}
