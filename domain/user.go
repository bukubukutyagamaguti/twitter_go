package domain

import (
	"api/server/config"
	"database/sql"
	"time"
)

type Users []User
type User struct {
	Id        int            `json:"id" param:"id" gorm:"primaryKey"`
	Name      string         `json:"name" validate:"required" gorm:"column:name"`
	Email     string         `json:"email" validate:"required,email" gorm:"column:email"`
	Password  string         `json:"password" validate:"required,password"`
	Token     sql.NullString `json:"token" validate:"required,token"`
	Follows   []*Follow      `json:"follow"`
	Posts     []*Post        `json:"post"`
	DeletedAt sql.NullTime   `json:"deleted_at" validate:"deleted_at" gorm:"column:deleted_at"`
	UpdatedAt string         `json:"updated_at" validate:"required,updated_at" gorm:"column:updated_at"`
	CreatedAt string         `json:"created_at" validate:"required,created_at" gorm:"column:created_at"`
}

type LoginUser struct {
	Id       int    `json:"id" param:"id" gorm:"primaryKey"`
	Email    string `json:"email" validate:"required"`
	Name     string `json:"name"`
	Password string `json:"password" validate:"required,password"`
}

func NewUser() User {
	return User{
		Name:      "",
		Email:     "",
		Password:  "",
		Token:     sql.NullString{},
		Follows:   []*Follow{nil},
		Posts:     []*Post{nil},
		DeletedAt: sql.NullTime{},
		UpdatedAt: time.Now().Format(config.TimeFormat),
		CreatedAt: "",
	}
}
