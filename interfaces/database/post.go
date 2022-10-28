package database

import (
	"api/server/domain"
)

type PostRepository struct {
	SqlHandler
}

func (repo *PostRepository) Store(p domain.Post) (post domain.Post, err error) {
	if err = repo.Create(&p).Error; err != nil {
		return
	}
	return
}
