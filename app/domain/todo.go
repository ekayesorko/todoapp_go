package domain

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Description string `json:"description"`
	Priority    uint   `json:"priority"`
	UserID      uint   `json:"userID"`
	StatusID    uint   `json:"statusID"`
}
