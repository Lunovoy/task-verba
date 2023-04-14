package repository

import (
	"github.com/Lunovoy/test-project-verba/internal/models"
	"github.com/jmoiron/sqlx"
)

type User interface {
	AddUser(user models.User) ([]string, error)
}

type Repository struct {
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserPostgres(db),
	}
}
