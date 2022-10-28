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

func (repo *PostRepository) PreloadById(table string, id int) (post domain.Post, err error) {
	if err = repo.Find(&post, table, id).Error; err != nil {
		return
	}
	return
}

func (repo *PostRepository) Related(table string, query string, id int) (posts domain.Posts, err error) {
	if err = repo.PreloadAndWhere(&posts, table, query, id).Error; err != nil {
		return
	}
	return
}
