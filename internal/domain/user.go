package domain

import (
	"time"

	"github.com/DimasPramantya/goMiniProject/internal/dto"
)

type User struct {
	ID         int       `db:"id" json:"id"`
	Username   string    `db:"username" json:"username"`
	Password   string    `db:"password" json:"password"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	CreatedBy  string    `db:"created_by" json:"created_by"`
	ModifiedAt *time.Time `db:"modified_at" json:"modified_at"`
	ModifiedBy *string    `db:"modified_by" json:"modified_by"`
}

type UserRepository interface {
	FindByID(id int) (*User, error)
	Create(user User) (*User, error)
	FindByUsername(username string) (*User, error)
}

type UserUsecase interface {
	Login(req dto.ReqLogin) (*dto.ResLogin, error)
	Register(req dto.ReqRegister) error
	GetByID(id int) (*dto.ResUser, error)
}