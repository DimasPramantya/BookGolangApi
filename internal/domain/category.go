package domain

import (
	"time"

	"github.com/DimasPramantya/goMiniProject/internal/dto"
)

type Category struct {
	ID         int       `db:"id" json:"id"`
	Name       string    `db:"name" json:"name"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	CreatedBy  string    `db:"created_by" json:"created_by"`
	ModifiedAt *time.Time `db:"modified_at" json:"modified_at"`
	ModifiedBy *string    `db:"modified_by" json:"modified_by"`
}

type CategoryRepository interface {
	FindByID(id int) (*Category, error)
	Create(category Category) (*Category, error)
	Update(category Category) (*Category, error)
	Delete(id int) error
	FindAll() ([]Category, error)
}

type CategoryUsecase interface {
	FindByID(id int) (*dto.ResCategory, error)
	Create(req dto.CreateCategoryRequest, createdBy string) (dto.ResCategory, error)
	Update(req dto.UpdateCategoryRequest, id int, modifiedBy string) (dto.ResCategory, error)
	Delete(id int) error
	FindAll() ([]dto.ResCategory, error)
	FindAllBooksByCategoryID(categoryID int) ([]dto.ResBook, error)
}