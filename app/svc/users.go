package svc

import (
	"todolist-assignment/app/serializer"
	"todolist-assignment/infra/errors"
)

type IUsers interface {
	CreateUser(req serializer.SignupReq) (*serializer.SignupResp, *errors.RestErr)
}
