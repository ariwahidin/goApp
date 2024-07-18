package repository

import (
	"database/sql"
	"goapp/internal/model"
)

type UserRepository interface {
	FetchAll() ([]model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FetchAll() ([]model.User, error) {
	rows, err := r.db.Query("SELECT id, username, email FROM users WHERE email = 'Blair.Parisian94@gmail.com'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
