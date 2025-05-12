package dto

type ResUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	CreatedAt string `json:"created_at"`
	CreatedBy string `json:"created_by"`
	ModifiedAt *string `json:"modified_at"`
	ModifiedBy *string `json:"modified_by"`
}