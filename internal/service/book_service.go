package service

import (
	"github.com/DimasPramantya/goMiniProject/internal/domain"
	"github.com/DimasPramantya/goMiniProject/internal/dto"
	"github.com/DimasPramantya/goMiniProject/utils/helper"
)

type bookService struct {
	bookRepo     domain.BookRepository
	categoryRepo domain.CategoryRepository
}

func NewBookService(
	bookRepo domain.BookRepository,
	categoryRepo domain.CategoryRepository,
) domain.BookUsecase {
	return &bookService{
		bookRepo:     bookRepo,
		categoryRepo: categoryRepo,
	}
}

func (b *bookService) Create(book dto.ReqCreateBook, username string) (*dto.ResBook, error) {

	if book.TotalPage >= 100 {
		thickness := "thick"
		book.Thickness = &thickness
	} else if book.TotalPage < 100 && book.TotalPage >= 0 {
		thickness := "thin"
		book.Thickness = &thickness
	} else {
		return nil, domain.ErrTotalPageNotValid
	}

	if book.ReleaseYear > 2024 || book.ReleaseYear < 1980 {
		return nil, domain.ErrReleaseYearNotValid
	}

	if book.CategoryID != nil {
		category, err := b.categoryRepo.FindByID(*book.CategoryID)
		if err != nil {
			return nil, err
		}
		if category == nil {
			return nil, domain.ErrNotFound
		}
	}

	newBook := domain.Book{
		Title:       book.Title,
		Description: book.Description,
		ImageURL:    book.ImageURL,
		ReleaseYear: book.ReleaseYear,
		Price:       book.Price,
		TotalPage:   book.TotalPage,
		Thickness:   *book.Thickness,
		CategoryID:  book.CategoryID,
		CreatedBy:   username,
	}

	err := b.bookRepo.Insert(&newBook)
	if err != nil {
		return nil, err
	}

	result := mapToResBook(newBook)

	return &result, nil
}

func (b *bookService) Update(id int, book dto.ReqUpdateBook, username string) (*dto.ResBook, error) {
	bookFromDB, err := b.bookRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if bookFromDB == nil {
		return nil, domain.ErrNotFound
	}

	if book.TotalPage >= 100 {
		thickness := "thick"
		book.Thickness = &thickness
	} else if book.TotalPage < 100 && book.TotalPage >= 0 {
		thickness := "thin"
		book.Thickness = &thickness
	} else {
		return nil, domain.ErrTotalPageNotValid
	}

	if book.ReleaseYear > 2024 || book.ReleaseYear < 1980 {
		return nil, domain.ErrReleaseYearNotValid
	}

	if book.CategoryID != nil {
		category, err := b.categoryRepo.FindByID(*book.CategoryID)
		if err != nil {
			return nil, err
		}
		if category == nil {
			return nil, domain.ErrNotFound
		}
	}

	bookFromDB.Title = book.Title
	bookFromDB.Description = book.Description
	bookFromDB.ImageURL = book.ImageURL
	bookFromDB.ReleaseYear = book.ReleaseYear
	bookFromDB.Price = book.Price
	bookFromDB.TotalPage = book.TotalPage
	bookFromDB.Thickness = *book.Thickness
	bookFromDB.CategoryID = book.CategoryID
	bookFromDB.ModifiedBy = &username

	err = b.bookRepo.Update(bookFromDB)
	if err != nil {
		return nil, err
	}

	result := mapToResBook(*bookFromDB)

	return &result, nil
}

func (b *bookService) Delete(id int) error {
	book, err := b.bookRepo.FindByID(id)
	if err != nil {
		return err
	}
	if book == nil {
		return domain.ErrNotFound
	}

	err = b.bookRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (b *bookService) GetAll() ([]dto.ResBook, error) {
	books, err := b.bookRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var resBooks []dto.ResBook
	for _, book := range books {
		resBook := mapToResBook(book)
		resBooks = append(resBooks, resBook)
	}

	return resBooks, nil
}

func (b *bookService) GetByID(id int) (dto.ResBook, error) {
	book, err := b.bookRepo.FindByID(id)
	if err != nil {
		return dto.ResBook{}, err
	}
	if book == nil {
		return dto.ResBook{}, domain.ErrNotFound
	}

	resBook := mapToResBook(*book)

	return resBook, nil
}

func mapToResBook(book domain.Book) dto.ResBook {
	return dto.ResBook{
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
}
