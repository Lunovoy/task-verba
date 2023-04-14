package service

import (
	"github.com/Lunovoy/test-project-verba/internal/models"
	"github.com/Lunovoy/test-project-verba/internal/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) AddUser(user models.User) ([]string, error) {
	return s.repo.AddUser(user)
}
