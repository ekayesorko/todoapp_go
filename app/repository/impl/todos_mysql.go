package impl

import (
	"gorm.io/gorm"
	"todolist-assignment/app/domain"
	"todolist-assignment/app/repository"
	"todolist-assignment/infra/errors"
	"todolist-assignment/utils"
)

type Todos struct {
	DB *gorm.DB
}

func NewMysqlTodosRepository(db *gorm.DB) repository.ITodo {
	return &Todos{
		DB: db,
	}
}

func (tRep *Todos) FirstOrCreateStatus(_status string) (*domain.TodoStatus, *errors.RestErr) {
	var tStatusDom *domain.TodoStatus
	res := tRep.DB.Model(&domain.TodoStatus{}).
		Where(&domain.TodoStatus{Status: _status}).
		FirstOrCreate(&tStatusDom)
	if res.Error != nil {
		return nil, errors.NewInternalServerError(res.Error.Error())
	}
	return tStatusDom, nil
}

func (tRep *Todos) Save(todo *domain.Todo) (*domain.Todo, *errors.RestErr) {
	res := tRep.DB.Model(&domain.Todo{}).Create(&todo)
	if res.Error != nil {
		return nil, errors.NewBadRequestError("todos rep error")
	}
	return todo, nil
}

func (tRep *Todos) GetAll(userID uint) ([]*domain.IntermediateTodoResponse, *errors.RestErr) {
	var todoList []*domain.IntermediateTodoResponse
	res := tRep.DB.Model(&domain.Todo{}).
		Select("todos.*, todo_statuses.*").
		Joins("join todo_statuses on todo_statuses.id = todos.todo_status_id").
		Where("user_id = ?", userID).
		Scan(&todoList)
	if res.Error != nil {
		return nil, errors.NewNotFoundError("not found")
	}
	return todoList, nil
}

func (tRep *Todos) Get(todoID uint, userID uint) (*domain.IntermediateTodoResponse, *errors.RestErr) {
	var todo *domain.IntermediateTodoResponse
	res := tRep.DB.Model(&domain.Todo{}).
		Select("todos.*, todo_statuses.*").
		Joins("join todo_statuses on todo_statuses.id = todos.todo_status_id").
		Where("user_id = ? and todos.id = ?", userID, todoID).
		First(&todo)
	if res.Error != nil {
		return nil, errors.NewNotFoundError(res.Error.Error())
	}
	return todo, nil
}

func (tRep *Todos) Update(todo *domain.Todo) (*domain.IntermediateTodoResponse, *errors.RestErr) {
	var res *gorm.DB

	var itr *domain.IntermediateTodoResponse
	res = tRep.DB.Model(&domain.Todo{}).
		Where("user_id = ? and id = ?", todo.UserID, todo.ID).
		Updates(&todo)
	if res.RowsAffected == 0 {
		return nil, errors.NewBadRequestError(errors.ErrRecordNotFound)
	}
	if err := utils.StructToStruct(&todo, &itr); err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	return itr, nil
}

func (tRep *Todos) Delete(todoID uint, userID uint) *errors.RestErr {
	var todo *domain.Todo
	res := tRep.DB.Model(&domain.Todo{}).
		Where("user_id = ? and id = ?", userID, todoID).
		Delete(&todo)
	if res.Error != nil {
		return errors.NewNotFoundError(res.Error.Error())
	}
	return nil
}
