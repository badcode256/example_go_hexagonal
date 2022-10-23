package user

import (
	"database/sql"
	"fmt"

	domain "github.com/badcode256/example_go_hexagonal/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user domain.IUser) error {

	_, err := r.db.Exec(`exec create_user @username, @email, @password`,
		sql.Named("username", user.User_name),
		sql.Named("email", user.Email),
		sql.Named("password", user.Password),
	)
	if err != nil {
		return fmt.Errorf("error create user: %v", err)
	}
	return nil

}
