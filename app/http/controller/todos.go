package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"todolist-assignment/app/http/middleware"
	"todolist-assignment/app/serializer"
	"todolist-assignment/app/svc"
	"todolist-assignment/utils/const"
)

type todos struct {
	tokenSvc svc.IToken
	todoSvc  svc.ITodos
}

func NewTodosController(_grp interface{}, _tokenSvc svc.IToken, _todoSvc svc.ITodos) {
	g := _grp.(*echo.Group)
	todo := &todos{
		tokenSvc: _tokenSvc,
		todoSvc:  _todoSvc,
	}
	g.POST("/todo/create", todo.Create, middleware.JWTParser)
	g.GET("/todo/get", todo.GetAll, middleware.JWTParser)
	g.GET("/todo/get/:id", todo.Get, middleware.JWTParser)
	g.PUT("/todo/update/:id", todo.Update, middleware.JWTParser)
	g.DELETE("/todo/delete/:id", todo.Delete, middleware.JWTParser)
}

func (todos *todos) Create(ctx echo.Context) error {
	userID := ctx.Get(_const.LoggedInUser).(*serializer.LoggedInUser).UserID

	var tcs *serializer.CreateTodoRequest
	tcs.Status = "Pending"

	if err := ctx.Bind(&tcs); err != nil {
		return ctx.JSON(http.StatusUnauthorized, "bad json")
	}
	if err := tcs.Validate(); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	tcs.UserID = userID
	resp, err := todos.todoSvc.Create(*tcs)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusCreated, resp)
}

func (todos *todos) GetAll(ctx echo.Context) error {
	userID := ctx.Get(_const.LoggedInUser).(*serializer.LoggedInUser).UserID

	serResp, err := todos.todoSvc.GetAll(userID)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err)
	}
	return ctx.JSON(http.StatusOK, serResp)
}

func (todos *todos) Get(ctx echo.Context) error {
	userID := ctx.Get(_const.LoggedInUser).(*serializer.LoggedInUser).UserID

	todoIDstring := ctx.Param("id")
	todoID, err3 := strconv.ParseUint(todoIDstring, 10, 32)
	if err3 != nil {
		return ctx.JSON(http.StatusBadRequest, err3)
	}
	sResp, err4 := todos.todoSvc.Get(uint(todoID), userID)
	if err4 != nil {
		return ctx.JSON(http.StatusBadRequest, err4)
	}
	return ctx.JSON(http.StatusOK, sResp)
}

func (todos *todos) Update(ctx echo.Context) error {
	userID := ctx.Get(_const.LoggedInUser).(*serializer.LoggedInUser).UserID

	todoIDString := ctx.Param("id")
	todoID, err3 := strconv.ParseUint(todoIDString, 10, 32)
	if err3 != nil {
		return ctx.JSON(http.StatusBadRequest, err3)
	}
	var upSer serializer.UpdateTodoRequest

	if err := ctx.Bind(&upSer); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	upSer.ID = uint(todoID)
	upSer.UserID = userID
	if err := upSer.Validate(); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	todoResp, err4 := todos.todoSvc.Update(upSer)
	if err4 != nil {
		return ctx.JSON(http.StatusBadRequest, err4)
	}
	return ctx.JSON(http.StatusOK, todoResp)
}

func (todos *todos) Delete(ctx echo.Context) error {
	userID := ctx.Get(_const.LoggedInUser).(*serializer.LoggedInUser).UserID

	todoIDString := ctx.Param("id")
	todoID, err3 := strconv.ParseUint(todoIDString, 10, 32)
	if err3 != nil {
		return ctx.JSON(http.StatusBadRequest, err3)
	}
	err4 := todos.todoSvc.Delete(uint(todoID), userID)
	if err4 != nil {
		return ctx.JSON(http.StatusNotAcceptable, err4)
	}
	return ctx.JSON(http.StatusOK, nil)
}
