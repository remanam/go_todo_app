package repository

import (
	"todo"

	"github.com/jackc/pgx/v4"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}

type Repository struct {
	Authorization
}

func NewRepository(conn *pgx.Conn) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(conn),
	}
}
