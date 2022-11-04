package domain

import (
	"api/server/config"
	"time"

	"gorm.io/gorm"
)

type Follows []Follow
type Follow struct {
	Id        int            `json:"id" param:"id" gorm:"primaryKey"`
	UserId    int            `json:"user_id" validate:"required" gorm:"column:user_id"`
	User      *User          `json:"user"`
	FollowId  int            `json:"follow_id" validate:"required" gorm:"column:follow_id"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" validate:"deleted_at" gorm:"column:deleted_at"`
	CreatedAt string         `json:"created_at" validate:"required,created_at" gorm:"column:created_at"`
}

func NewFollow() Follow {
	return Follow{
		UserId:    0,
		User:      &User{},
		FollowId:  0,
		DeletedAt: gorm.DeletedAt{},
		CreatedAt: time.Now().Format(config.TimeFormat),
	}
}
