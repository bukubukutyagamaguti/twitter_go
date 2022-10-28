package domain

type Users []User
type User struct {
	Id        int       `json:"id" param:"id" gorm:"primaryKey"`
	Name      string    `json:"name" validate:"required" gorm:"column:name"`
	Email     string    `json:"email" validate:"required,email" gorm:"column:email"`
	Password  string    `json:"password" validate:"required,password"`
	Token     string    `json:"token" validate:"required,token"`\
	DeletedAt string    `json:"deleted_at" validate:"deleted_at" gorm:"column:deleted_at"`
	UpdatedAt string    `json:"updated_at" validate:"required,updated_at" gorm:"column:updated_at"`
	CreatedAt string    `json:"created_at" validate:"required,created_at" gorm:"column:created_at"`
}