package usecase

import (
	"api/server/domain"
	"api/server/interfaces/token"
)

type UserInteractor struct {
	UserRepository UserRepository
	Tokenizer      token.Tokenizer
}

func (interactor *UserInteractor) UserById(id int) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindById(id)
	return
}

func (interactor *UserInteractor) Users() (users domain.Users, err error) {
	users, err = interactor.UserRepository.FindAll()
	return
}

func (interactor *UserInteractor) Add(u domain.User) (user domain.User, err error) {
	user, err = interactor.UserRepository.Store(u)
	return
}

func (interactor *UserInteractor) Update(u domain.User) (user domain.User, err error) {
	user, err = interactor.UserRepository.Update(u)
	return
}

func (interactor *UserInteractor) DeleteById(u domain.User) (err error) {
	err = interactor.UserRepository.DeleteById(u)
	return
}

func (interactor *UserInteractor) Login(login domain.LoginUser) (domain.User, domain.Token, error) {
	var token domain.Token
	user, err := interactor.UserRepository.WhereByEmail(login.Email)
	if err != nil {
		return user, token, err
	}

	if user.Password != login.Password || user.Email != login.Email {
		return user, token, err
	}

	token, err = interactor.Tokenizer.New(user)
	if err != nil {
		return user, token, err
	}

	return user, token, nil
}
