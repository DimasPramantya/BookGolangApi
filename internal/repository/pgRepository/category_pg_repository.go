package pgrepository

import (
	"database/sql"

	"github.com/DimasPramantya/goMiniProject/internal/domain"
)

type categoryPGRepo struct {
	DB *sql.DB
}

func NewCategoryPGRepo(db *sql.DB) domain.CategoryRepository {
	return &categoryPGRepo{DB: db}
}

func (c *categoryPGRepo) Create(category domain.Category) (*domain.Category, error) {
	err := c.DB.QueryRow(`INSERT INTO categories (name, created_by) VALUES ($1, $2) RETURNING id, created_at`,
		category.Name, category.CreatedBy).Scan(&category.ID, &category.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (c *categoryPGRepo) Delete(id int) error {
	_, err := c.DB.Exec(`DELETE FROM categories WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (c *categoryPGRepo) FindAll() ([]domain.Category, error) {
	rows, err := c.DB.Query(`SELECT id, name, created_at, created_by, modified_at, modified_by FROM categories`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		var category domain.Category
		err := rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy, 
			&category.ModifiedAt, &category.ModifiedBy)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return categories, nil
}

func (c *categoryPGRepo) FindByID(id int) (*domain.Category, error) {
	row := c.DB.QueryRow(`SELECT id, name, created_at, created_by, modified_at, modified_by FROM categories WHERE id = $1`, id)
	category := &domain.Category{}
	err := row.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy, 
		&category.ModifiedAt, &category.ModifiedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return category, nil
}

func (c *categoryPGRepo) Update(category domain.Category) (*domain.Category, error) {
	row := c.DB.QueryRow(`
		UPDATE categories 
		SET name = $1, modified_by = $2, modified_at = NOW() 
		WHERE id = $3 
		RETURNING id, name, created_at, created_by, modified_at, modified_by
	`, category.Name, category.ModifiedBy, category.ID)

	var updatedCategory domain.Category
	err := row.Scan(
		&updatedCategory.ID,
		&updatedCategory.Name,
		&updatedCategory.CreatedAt,
		&updatedCategory.CreatedBy,
		&updatedCategory.ModifiedAt,
		&updatedCategory.ModifiedBy,
	)
	if err != nil {
		return nil, err
	}

	return &updatedCategory, nil
}

