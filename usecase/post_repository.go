//go:generate mockgen -source=post_repository.go -destination=./mock/post.go
package usecase

import "api/server/domain"

type PostRepository interface {
	Store(domain.Post) (domain.Post, error)
	Related(string, string, int) (domain.Posts, error)
}
