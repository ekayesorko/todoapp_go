package serializer

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CreateTodoRequest struct {
	UserID      uint   `json:"userID"`
	Description string `json:"description"`
	Priority    uint   `json:"priority"`
	Status      string `json:"status"`
}

type TodoResponse struct {
	TodoId      uint   `json:"ID"`
	Description string `json:"description"`
	Priority    uint   `json:"priority"`
	Status      string `json:"status"`
}

type GetAllTodoResponse struct {
	UserID   uint           `json:"userID"`
	TodoList []TodoResponse `json:"todoList"`
	Status   string         `json:"status"`
}

type UpdateTodoRequest struct {
	ID          uint   `json:"ID"`
	UserID      uint   `json:"userID"`
	Description string `json:"description"`
	Priority    uint   `json:"priority"`
	Status      string `json:"status"`
}

func (utr UpdateTodoRequest) Validate() error {
	return validation.ValidateStruct(&utr,
		validation.Field(&utr.ID, validation.Required),
		validation.Field(&utr.UserID, validation.Required),
		validation.Field(&utr.Priority, validation.Min(uint(0)), validation.Max(uint(5))),
	)
}

func (ctr CreateTodoRequest) Validate() error {
	return validation.ValidateStruct(&ctr,
		validation.Field(&ctr.Description, validation.Required),
		validation.Field(&ctr.Status, validation.Required),
		validation.Field(&ctr.Priority, validation.Required, validation.Min(uint(0)), validation.Max(uint(5))),
	)
}
