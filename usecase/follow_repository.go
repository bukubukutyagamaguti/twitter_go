package usecase

import "api/server/domain"

type FollowRepository interface {
	Update(domain.Follow) (domain.Follow, error)
	DeleteById(domain.Follow) error
	WhereByUserIdAndFollowId(string, int, int) (domain.Follow, error)
}