package repository

import (
	"todolist-assignment/app/domain"
	"todolist-assignment/infra/errors"
)

type ITodo interface {
	Save(todo *domain.Todo) (*domain.Todo, *errors.RestErr)
	GetAll(userID uint) ([]*domain.Todo, *errors.RestErr)
	Get(todoID uint, userID uint) (*domain.Todo, *errors.RestErr)
	Update(todo *domain.Todo) (*domain.Todo, *errors.RestErr)
	Delete(todoID uint, userID uint) *errors.RestErr
}
