package usecase

import "api/server/domain"

type PostRepository interface {
	Store(domain.Post) (domain.Post, error)
}
