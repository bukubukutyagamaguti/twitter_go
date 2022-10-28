package usecase

import (
	"api/server/domain"
)

type PostInteractor struct {
	PostRepository PostRepository
}

func (interactor *PostInteractor) Add(u domain.Post) (post domain.Post, err error) {
	post, err = interactor.PostRepository.Store(u)
	return
}
