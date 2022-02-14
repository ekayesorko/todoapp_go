package repository

import (
	"todolist-assignment/app/domain"
	"todolist-assignment/infra/errors"
)

type ITodo interface {
	Save(todo *domain.Todo) (*domain.Todo, *errors.RestErr)
	GetAll(userID uint) ([]*domain.IntermediateTodoResponse, *errors.RestErr)
	Get(todoID uint, userID uint) (*domain.IntermediateTodoResponse, *errors.RestErr)
	Update(todo *domain.Todo) (*domain.IntermediateTodoResponse, *errors.RestErr)
	Delete(todoID uint, userID uint) *errors.RestErr
	FirstOrCreateStatus(_status string) (*domain.TodoStatus, *errors.RestErr)
}
