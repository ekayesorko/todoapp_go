package repository

import (
	"todolist-assignment/app/domain"
	"todolist-assignment/infra/errors"
)

type IUsers interface {
	Save(user *domain.User) (*domain.User, *errors.RestErr)
	AuthAndGetID(username string, password string) (uint, *errors.RestErr)
}
