package domain

import (
	"api/server/config"
	"database/sql"
	"time"
)

type Posts []Post
type Post struct {
	Id        int `json:"id" param:"id" gorm:"primaryKey"`
	UserId    int `json:"user_id" validate:"required" gorm:"column:user_id"`
	User      *User
	Message   string       `json:"message" validate:"required" gorm:"column:message"`
	DeletedAt sql.NullTime `json:"deleted_at" validate:"deleted_at" gorm:"column:deleted_at"`
	CreatedAt string       `json:"created_at" validate:"required,created_at" gorm:"column:created_at"`
}

func NewPost() Post {
	return Post{
		UserId:    0,
		User:      &User{},
		Message:   "",
		DeletedAt: sql.NullTime{},
		CreatedAt: time.Now().Format(config.TimeFormat),
	}
}
