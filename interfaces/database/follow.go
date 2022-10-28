package database

import "api/server/domain"

type FollowRepository struct {
	SqlHandler
}

func (repo *FollowRepository) Update(f domain.Follow) (follow domain.Follow, err error) {
	if err = repo.Save(&f).Error; err != nil {
		return
	}
	return
}
func (repo *FollowRepository) DeleteById(f domain.Follow) (err error) {
	if err = repo.Delete(&f).Error; err != nil {
		return
	}
	return
}

func (repo *FollowRepository) WhereByUserIdAndFollowId(query string, id int, followId int) (follow domain.Follow, err error) {
	if err = repo.Where(&follow, query, id, followId).Error; err != nil {
		return
	}
	return
}
