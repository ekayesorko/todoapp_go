package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"todolist-assignment/app/http/middleware"
	"todolist-assignment/app/serializer"
	"todolist-assignment/app/svc"
	_const "todolist-assignment/utils/const"
)

type users struct {
	us svc.IUsers
	ts svc.IToken
}

func NewUsersController(_grp interface{}, _us svc.IUsers, _ts svc.IToken) {
	uc := &users{
		us: _us,
		ts: _ts,
	}
	g := _grp.(*echo.Group)
	g.POST("/user/create", uc.Create)
	g.GET("/user/check", uc.CheckLogin, middleware.JWTParser)
}

func (uc *users) Create(ctx echo.Context) error {
	//var ud *domain.User
	var signupReq *serializer.SignupReq
	fmt.Println("controller started")

	if err := ctx.Bind(&signupReq); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	if err := signupReq.Validate(); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	signupResp, err := uc.us.CreateUser(*signupReq)
	if err != nil {
		return ctx.JSON(http.StatusNotAcceptable, err)
	}
	return ctx.JSON(http.StatusCreated, signupResp)
}

func (uc *users) CheckLogin(ctx echo.Context) error {
	userID := ctx.Get(_const.LoggedInUser)
	return ctx.JSON(http.StatusOK, userID)
}
