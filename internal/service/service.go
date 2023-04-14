package service

import (
	"github.com/Lunovoy/test-project-verba/internal/models"
	"github.com/Lunovoy/test-project-verba/internal/repository"
)

type User interface {
	AddUser(user models.User) ([]string, error)
}

type Service struct {
	User
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repo.User),
	}
}
