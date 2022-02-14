package app

import (
	"todolist-assignment/app/http/controller"
	rImpl "todolist-assignment/app/repository/impl"
	sImpl "todolist-assignment/app/svc/impl"
	"todolist-assignment/infra/conn"
)

func Init(grp interface{}) {
	db := conn.ConnectDB()

	userRepo := rImpl.NewMysqlUsersRepository(db)
	todoRepo := rImpl.NewMysqlTodosRepository(db)

	userService := sImpl.NewUsersService(userRepo)
	tokenService := sImpl.NewTokenService(userRepo)
	authService := sImpl.NewAuthService(userRepo, tokenService)
	todoService := sImpl.NewTodoService(todoRepo, tokenService)

	controller.NewUsersController(grp, userService, tokenService)
	controller.NewAuthController(grp, authService, userService)
	controller.NewTodosController(grp, tokenService, todoService)
}
