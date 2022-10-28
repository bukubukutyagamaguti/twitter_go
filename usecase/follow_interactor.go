package usecase

import (
	"api/server/domain"
)

type FollowInteractor struct {
	FollowRepository FollowRepository
}

func (interactor *FollowInteractor) Update(u domain.Follow) (follow domain.Follow, err error) {
	follow, err = interactor.FollowRepository.Update(u)
	return
}

func (interactor *FollowInteractor) DeleteById(u domain.Follow) (err error) {
	err = interactor.FollowRepository.DeleteById(u)
	return
}

func (interactor *FollowInteractor) SearchByFollowIdAndUserId(query string, userId int, followId int) (follow domain.Follow, err error) {
	follow, err = interactor.FollowRepository.WhereByUserIdAndFollowId(query, userId, followId)
	return
}
