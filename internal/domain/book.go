package domain

import (
	"time"

	"github.com/DimasPramantya/goMiniProject/internal/dto"
)

type Book struct {
	ID          int       `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	ImageURL    string    `db:"image_url" json:"image_url"`
	ReleaseYear int       `db:"release_year" json:"release_year"`
	Price       int       `db:"price" json:"price"`
	TotalPage   int       `db:"total_page" json:"total_page"`
	Thickness   string    `db:"thickness" json:"thickness"`
	CategoryID  *int      `db:"category_id" json:"category_id"` // nullable
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	CreatedBy   string    `db:"created_by" json:"created_by"`
	ModifiedAt  *time.Time `db:"modified_at" json:"modified_at"`
	ModifiedBy  *string    `db:"modified_by" json:"modified_by"`
}

type BookRepository interface {
	FindByID(id int) (*Book, error)
	FindAll() ([]Book, error)
	FindAllByCategoryID(categoryID int) ([]Book, error)
	Insert(book *Book) error
	Update(book *Book) error
	Delete(id int) error
}

type BookUsecase interface {
	GetByID(id int) (dto.ResBook, error)
	GetAll() ([]dto.ResBook, error)
	Create(book dto.ReqCreateBook, username string) (*dto.ResBook, error)
	Delete(id int) error
}
