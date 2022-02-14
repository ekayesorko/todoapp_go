package svc

import "todolist-assignment/infra/errors"

type IToken interface {
	CreateToken(userId uint) (string, *errors.RestErr)
	VerifyToken(token string) (uint, *errors.RestErr)
}
