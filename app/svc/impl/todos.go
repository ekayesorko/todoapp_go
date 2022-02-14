package impl

import (
	"todolist-assignment/app/domain"
	"todolist-assignment/app/repository"
	"todolist-assignment/app/serializer"
	"todolist-assignment/app/svc"
	"todolist-assignment/infra/errors"
	"todolist-assignment/utils"
)

type todos struct {
	tokenSvc svc.IToken
	todoRep  repository.ITodo
}

func NewTodoService(_todoRep repository.ITodo, _tokenSvc svc.IToken) svc.ITodos {
	return &todos{
		tokenSvc: _tokenSvc,
		todoRep:  _todoRep,
	}
}

func (t *todos) Create(request serializer.CreateTodoRequest) (*serializer.TodoResponse, *errors.RestErr) {
	todoDom := &domain.Todo{}
	todoStatusDom, err := t.todoRep.FirstOrCreateStatus(request.Status)
	if err != nil {
		return nil, errors.NewBadRequestError(err.Error)
	}
	err3 := utils.StructToStruct(request, todoDom)
	if err3 != nil {
		return nil, errors.NewBadRequestError(err3.Error())
	}
	todoDom.TodoStatusID = todoStatusDom.ID
	saveResp, err2 := t.todoRep.Save(todoDom)
	if err2 != nil {
		return nil, errors.NewBadRequestError(err2.Error)
	}
	resp := &serializer.TodoResponse{
		TodoId:       saveResp.ID,
		Description:  saveResp.Description,
		Priority:     saveResp.Priority,
		TodoStatusID: todoStatusDom.ID,
		Status:       todoStatusDom.Status,
	}
	return resp, nil
}

func (t *todos) GetAll(userID uint) (*serializer.GetAllTodoResponse, *errors.RestErr) {
	resp, err := t.todoRep.GetAll(userID)
	if err != nil {
		return nil, err
	}
	var todoArray []serializer.TodoResponse
	var todoResp *serializer.TodoResponse
	for _, r := range resp {
		utils.StructToStruct(r, &todoResp)
		todoArray = append(todoArray, *todoResp)
	}
	serResp := &serializer.GetAllTodoResponse{
		UserID:   userID,
		TodoList: todoArray,
	}
	return serResp, nil
}

func (t *todos) Get(todoID uint, userID uint) (*serializer.TodoResponse, *errors.RestErr) {
	tDom, err := t.todoRep.Get(todoID, userID)
	if err != nil {
		return nil, err
	}
	var tResp *serializer.TodoResponse
	if err := utils.StructToStruct(tDom, &tResp); err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	return tResp, nil
}

func (t *todos) Update(request serializer.UpdateTodoRequest) (*serializer.TodoResponse, *errors.RestErr) {
	var tDom *domain.Todo
	if err := utils.StructToStruct(request, &tDom); err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	var tsd *domain.TodoStatus
	var err2 *errors.RestErr
	if request.Status != "" {
		tsd, err2 = t.todoRep.FirstOrCreateStatus(request.Status)
		if err2 != nil {
			return nil, err2
		}
		tDom.TodoStatusID = tsd.ID
	}
	resp, err := t.todoRep.Update(tDom)
	if err != nil {
		return nil, errors.NewNotFoundError(err.Error)
	}
	var tResp *serializer.TodoResponse
	if err := utils.StructToStruct(resp, &tResp); err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	tResp.Status = tsd.Status
	return tResp, nil
}

func (t *todos) Delete(todoID uint, userID uint) *errors.RestErr {
	return t.todoRep.Delete(todoID, userID)
}
