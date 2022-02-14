package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"todolist-assignment/app/serializer"
	"todolist-assignment/app/svc"
)

type auth struct {
	authSvc svc.IAuth
	userSvc svc.IUsers
}

func NewAuthController(_grp interface{}, as svc.IAuth, us svc.IUsers) {
	g := _grp.(*echo.Group)
	ac := &auth{
		authSvc: as,
		userSvc: us,
	}
	g.POST("/auth/login", ac.Login)
}

func (ac *auth) Login(ctx echo.Context) error {
	fmt.Println("login controller")
	var loginReq *serializer.LoginReq
	if err := ctx.Bind(&loginReq); err != nil {
		fmt.Println("bad json")
	}
	resp, err := ac.authSvc.Login(loginReq)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, nil)
	}
	return ctx.JSON(http.StatusOK, resp)
}
