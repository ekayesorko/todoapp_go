package impl

import (
	"gorm.io/gorm"
	"todolist-assignment/app/domain"
	"todolist-assignment/app/repository"
	"todolist-assignment/infra/errors"
)

type Todos struct {
	DB *gorm.DB
}

func NewMysqlTodosRepository(db *gorm.DB) repository.ITodo {
	return &Todos{
		DB: db,
	}
}

func (tRep *Todos) Save(todo *domain.Todo) (*domain.Todo, *errors.RestErr) {
	res := tRep.DB.Model(&domain.Todo{}).Create(&todo)
	if res.Error != nil {
		return nil, errors.NewBadRequestError("todos rep error")
	}
	return todo, nil
}

func (tRep *Todos) GetAll(userID uint) ([]*domain.Todo, *errors.RestErr) {
	var todoList []*domain.Todo
	res := tRep.DB.Model(&domain.Todo{}).Where("user_id = ?", userID).Find(&todoList)
	if res.Error != nil {
		return nil, errors.NewNotFoundError("not found")
	}
	return todoList, nil
}

func (tRep *Todos) Get(todoID uint, userID uint) (*domain.Todo, *errors.RestErr) {
	var todo *domain.Todo
	res := tRep.DB.Model(&domain.Todo{}).Where("user_id = ? and id = ?", userID, todoID).First(&todo)
	if res.Error != nil {
		return nil, errors.NewNotFoundError(res.Error.Error())
	}
	return todo, nil
}

func (tRep *Todos) Update(todo *domain.Todo) (*domain.Todo, *errors.RestErr) {
	var res *gorm.DB
	//if todo.Priority == 0 {
	//	err := tRep.DB.Model(&domain.Todo{}).
	//		Select("priority").
	//		Where("user_id = ? and id = ?", todo.UserID, todo.ID).
	//		Find(&todo.Priority)
	//	if err.Error != nil {
	//		return nil, errors.NewNotFoundError(err.Error.Error())
	//	}
	//}
	//if todo.Description == "" {
	//	err := tRep.DB.Model(&domain.Todo{}).
	//		Select("description").
	//		Where("user_id = ? and id = ?", todo.UserID, todo.ID).
	//		Find(&todo.Description)
	//	if err.Error != nil {
	//		return nil, errors.NewNotFoundError(err.Error.Error())
	//	}
	//}
	res = tRep.DB.Model(&domain.Todo{}).
		Where("user_id = ? and id = ?", todo.UserID, todo.ID).
		Updates(&todo)
	if res.RowsAffected == 0 {
		return nil, errors.NewBadRequestError(errors.ErrRecordNotFound)
	}
	return todo, nil
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
