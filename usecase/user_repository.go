//go:generate mockgen -source=user_repository.go -destination=./mock/user.go
package usecase

import "api/server/domain"

type UserRepository interface {
	FindById(id int) (domain.User, error)
	FindAll() (domain.Users, error)
	WhereByEmail(string) (domain.User, error)
	Store(domain.User) (domain.User, error)
	Update(domain.User) (domain.User, error)
	DeleteById(domain.User) error
}
