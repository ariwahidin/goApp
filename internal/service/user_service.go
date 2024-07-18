package service

import (
	"goapp/internal/model"
	"goapp/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	return s.repo.FetchAll()
}
