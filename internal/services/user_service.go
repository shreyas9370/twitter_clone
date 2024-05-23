package services

import (
	"twitter-clone/internal/models"
	"twitter-clone/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(user *models.User) error {
	return s.repo.CreateUser(user)
}
