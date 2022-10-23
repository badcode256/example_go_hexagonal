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

func (r *UserRepository) UpdateUser(user domain.UUser) error {

	_, err := r.db.Exec(`exec edit_user @id, @username, @email, @password`,
		sql.Named("id", user.Id),
		sql.Named("username", user.User_name),
		sql.Named("email", user.Email),
		sql.Named("password", user.Password),
	)
	if err != nil {
		return fmt.Errorf("error update user: %v", err)
	}
	return nil

}

func (r *UserRepository) DeleteUser(id int64) error {

	_, err := r.db.Exec(`exec delete_user @id`,
		sql.Named("id", id),
	)
	if err != nil {
		return fmt.Errorf("error delete user: %v", err)
	}
	return nil

}

func (r *UserRepository) ListUsers() (*[]domain.Users, error) {
	var users []domain.Users
	rows, err := r.db.Query(
		`exec list_user`,
	)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user domain.Users
		err2 := rows.Scan(

			&user.Id,
			&user.User_name,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err2 != nil {
			return nil, err2
		}
		users = append(users, user)
	}

	return &users, nil
}
