package impl

import (
	"todolist-assignment/app/domain"
	"todolist-assignment/app/repository"
	"todolist-assignment/app/serializer"
	"todolist-assignment/app/svc"
	"todolist-assignment/infra/errors"
)

type users struct {
	ur repository.IUsers
}

func NewUsersService(_ur repository.IUsers) svc.IUsers {
	return &users{
		ur: _ur,
	}
}

func (us *users) CreateUser(req serializer.SignupReq) (*serializer.SignupResp, *errors.RestErr) {
	ud := domain.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}
	resp, saveErr := us.ur.Save(&ud)
	if saveErr != nil {
		return nil, saveErr
	}
	signupResp := &serializer.SignupResp{}
	signupResp.Username = resp.Username
	signupResp.ID = resp.ID
	return signupResp, nil
}
