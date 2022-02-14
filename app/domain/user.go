package domain

import (
	"time"
)

type User struct {
	ID        uint       `json:"id"`
	Username  string     `json:"username" gorm:"unique"`
	Password  string     `json:"password"`
	Email     string     `json:"email" gorm:"unique"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Todo      []Todo
}
