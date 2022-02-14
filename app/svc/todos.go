package svc

import (
	"todolist-assignment/app/serializer"
	"todolist-assignment/infra/errors"
)

type ITodos interface {
	Create(request serializer.CreateTodoRequest) (*serializer.TodoResponse, *errors.RestErr)
	GetAll(userID uint) (*serializer.GetAllTodoResponse, *errors.RestErr)
	Get(todoID uint, userID uint) (*serializer.TodoResponse, *errors.RestErr)
	Update(request serializer.UpdateTodoRequest) (*serializer.TodoResponse, *errors.RestErr)
	Delete(todoID uint, userID uint) *errors.RestErr
}
