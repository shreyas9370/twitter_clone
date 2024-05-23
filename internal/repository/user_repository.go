package repository

import (
	"database/sql"
	"twitter-clone/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	query := "INSERT INTO user (username, email, password) VALUES (?,?,?)"
	_, err := r.DB.Exec(query, user.Username, user.Email, user.Password)
	return err
}
