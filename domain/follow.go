package domain

import "gorm.io/gorm"

type Follows []Follow
type Follow struct {
	Id        int            `json:"id" param:"id" gorm:"primaryKey"`
	UserId    int            `json:"user_id" validate:"required" gorm:"column:user_id"`
	User      *User          `json:"user"`
	FollowId  int            `json:"follow_id" validate:"required" gorm:"column:follow_id"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" validate:"deleted_at" gorm:"column:deleted_at"`
	CreatedAt string         `json:"created_at" validate:"required,created_at" gorm:"column:created_at"`
}
