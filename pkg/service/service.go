package service

import (
	"todo"
	"todo/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repo *repository.Repository) *Service {
	
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}
