package dto

type ResBook struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	ImageURL    string     `json:"image_url"`
	ReleaseYear int        `json:"release_year"`
	Price       int        `json:"price"`
	TotalPage   int        `json:"total_page"`
	Thickness   string     `json:"thickness"`
	CategoryID  *int       `json:"category_id"`
	CreatedAt  string `json:"created_at"`
	CreatedBy  string `json:"created_by"`
	ModifiedAt *string `json:"modified_at"`
	ModifiedBy *string `json:"modified_by"`
}

type ReqCreateBook struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	ImageURL    string `json:"image_url" validate:"required"`
	ReleaseYear int    `json:"release_year" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	TotalPage   int    `json:"total_page" validate:"required"`
	Thickness   *string `json:"thickness"`
	CategoryID  *int   `json:"category_id"`
}

type ReqCreateBookSwagger struct {
    Title       string `json:"title"`
    Description string `json:"description"`
    ImageURL    string `json:"image_url"`
    ReleaseYear int    `json:"release_year"`
    Price       int    `json:"price"`
    TotalPage   int    `json:"total_page"`
    CategoryID  *int   `json:"category_id"`
}

type ReqUpdateBook struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	ImageURL    string `json:"image_url" validate:"required"`
	ReleaseYear int    `json:"release_year" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	TotalPage   int    `json:"total_page" validate:"required"`
	Thickness   *string `json:"thickness"`
	CategoryID  *int   `json:"category_id"`
}

type ReqUpdateBookSwagger struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	ImageURL    string `json:"image_url" validate:"required"`
	ReleaseYear int    `json:"release_year" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	TotalPage   int    `json:"total_page" validate:"required"`
	CategoryID  *int   `json:"category_id"`
}