package repository

import (
	"database/sql"
	"time"
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

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	var createdAt []byte
	query := "select id, username, email, password, created_at from user where email = ?"
	err := r.DB.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &createdAt)
	if err != nil {
		return nil, err
	}
	createdStr := string(createdAt)
	user.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(createdStr))
	if err != nil {
		return nil, err
	}
	return user, nil
}
