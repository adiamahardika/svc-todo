package repository

import "svc-todo/entity"

type ListRepositoryInterface interface {
	CreateList(request *entity.List) ([]entity.List, error)
}

func (repo *repository) CreateList(request *entity.List) ([]entity.List, error) {
	var list []entity.List

	error := repo.db.Table("list").Create(&request).Find(&list).Error

	return list, error
}
