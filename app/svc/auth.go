package svc

import (
	"todolist-assignment/app/serializer"
	"todolist-assignment/infra/errors"
)

type IAuth interface {
	Login(req *serializer.LoginReq) (*serializer.LoginResp, *errors.RestErr)
}
