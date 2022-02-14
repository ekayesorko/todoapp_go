package domain

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Description  string `json:"description"`
	Priority     uint   `json:"priority"`
	UserID       uint   `json:"userID"`
	TodoStatusID uint   `json:"todoStatusID"`
}

type IntermediateTodoResponse struct {
	ID           uint   `json:"ID"`
	Description  string `json:"description"`
	Priority     uint   `json:"priority"`
	UserID       uint   `json:"userID"`
	Status       string `json:"status"`
	TodoStatusID uint   `json:"todoStatusID"`
}
