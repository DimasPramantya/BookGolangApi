package pgrepository

import (
	"database/sql"

	"github.com/DimasPramantya/goMiniProject/internal/domain"
)

type bookPGRepo struct {
	DB *sql.DB
}

func NewBookPGRepo(db *sql.DB) domain.BookRepository {
	return &bookPGRepo{DB: db}
}

func (r *bookPGRepo) FindAllByCategoryID(categoryID int) ([]domain.Book, error) {
	var books []domain.Book
	row, err := r.DB.Query(`SELECT id, title, description, image_url, release_year, price, total_page,
	thickness, category_id, created_at, created_by, modified_at, modified_by FROM books 
	WHERE category_id = $1`, categoryID)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		var book domain.Book
		err := row.Scan(
			&book.ID, &book.Title, &book.Description, &book.ImageURL, 
			&book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness,
			&book.CategoryID, &book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	if err := row.Err(); err != nil {
		return nil, err
	}
	return books, nil
}

func (r *bookPGRepo) FindByID(id int) (*domain.Book, error) {
	row := r.DB.QueryRow(`SELECT id, title, description, category_id FROM books WHERE id = $1`, id)
	book := &domain.Book{}
	err := row.Scan(&book.ID, &book.Title, &book.Description, &book.CategoryID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return book, nil
}

func (r *bookPGRepo) Delete(id int) error {
	_, err := r.DB.Exec(`DELETE FROM books WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *bookPGRepo) FindAll() ([]domain.Book, error) {
	rows, err := r.DB.Query(`SELECT id, title, description, image_url, release_year, price, total_page,
	thickness, category_id, created_at, created_by, modified_at, modified_by FROM books`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []domain.Book
	for rows.Next() {
		var book domain.Book
		err := rows.Scan(
			&book.ID, &book.Title, &book.Description, &book.ImageURL,
			&book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness,
			&book.CategoryID, &book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return books, nil
}

func (r *bookPGRepo) Insert(book *domain.Book) error {
	err := r.DB.QueryRow(`INSERT INTO books (title, description, image_url, release_year, price, total_page,
	thickness, category_id, created_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id, created_at`,
		book.Title, book.Description, book.ImageURL, book.ReleaseYear,
		book.Price, book.TotalPage, book.Thickness, book.CategoryID,
		book.CreatedBy).Scan(&book.ID, &book.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *bookPGRepo) Update(book *domain.Book) error {
	err := r.DB.QueryRow(`UPDATE books SET title = $1, description = $2, image_url = $3, release_year = $4,
	price = $5, total_page = $6, thickness = $7, category_id = $8, modified_by = $9, modified_at = NOW()
	WHERE id = $10 returning modified_at, modified_by`, book.Title, book.Description, book.ImageURL,
		book.ReleaseYear, book.Price, book.TotalPage, book.Thickness,
		book.CategoryID, book.ModifiedBy, book.ID).Scan(&book.ModifiedAt, &book.ModifiedBy)
	if err != nil {
		return err
	}
	return nil
}