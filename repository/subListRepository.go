package repository

import "svc-todo/entity"

type SubListRepositoryInterface interface {
	CreateSubList(request *entity.SubList) ([]entity.SubList, error)
}

func (repo *repository) CreateSubList(request *entity.SubList) ([]entity.SubList, error) {
	var sub_list []entity.SubList

	error := repo.db.Table("sub_list").Create(&request).Find(&sub_list).Error

	return sub_list, error
}
