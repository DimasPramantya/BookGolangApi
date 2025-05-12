package pgrepository

import (
	"database/sql"

	"github.com/DimasPramantya/goMiniProject/internal/domain"
)

type userPgRepository struct {
	db *sql.DB
}

func NewUserPgRepository(db *sql.DB) domain.UserRepository {
	return &userPgRepository{db: db}
}

func (u *userPgRepository) Create(user domain.User) (*domain.User, error) {
	err := u.db.QueryRow(`INSERT INTO users (username, password, created_by) VALUES ($1, $2, $3) Returning id`,
		user.Username, user.Password, user.CreatedBy).Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userPgRepository) FindByID(id int) (*domain.User, error) {
	row := u.db.QueryRow(`SELECT id, username, password, created_at, created_by, modified_at, modified_by FROM users WHERE id = $1`, id)
	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt, &user.CreatedBy, &user.ModifiedAt, &user.ModifiedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (u *userPgRepository) FindByUsername(username string) (*domain.User, error) {
	row := u.db.QueryRow(`SELECT id, username, password, created_at, created_by FROM users WHERE username = $1`, username)
	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt, &user.CreatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}
