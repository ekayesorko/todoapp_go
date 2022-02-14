package impl

import (
	"todolist-assignment/app/repository"
	"todolist-assignment/app/serializer"
	"todolist-assignment/app/svc"
	"todolist-assignment/infra/errors"
)

type auth struct {
	uRep repository.IUsers
	tSvc svc.IToken
}

func NewAuthService(_urep repository.IUsers, _tsvc svc.IToken) svc.IAuth {
	return &auth{
		uRep: _urep,
		tSvc: _tsvc,
	}
}
func (auth *auth) Login(req *serializer.LoginReq) (*serializer.LoginResp, *errors.RestErr) {
	gotUsername := req.Username
	gotPassword := req.Password
	if id, err := auth.uRep.AuthAndGetID(gotUsername, gotPassword); err == nil {
		resp := &serializer.LoginResp{}
		resp.Username = gotUsername
		at, err := auth.tSvc.CreateToken(id)
		resp.AccessToken = at
		if err != nil {
			return nil, &errors.RestErr{Message: err.Error}
		}
		return resp, nil
	} else {
		return nil, &errors.RestErr{Message: "wrong username or password"}
	}
}
