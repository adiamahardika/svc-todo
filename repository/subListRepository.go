package repository

import (
	"svc-todo/entity"
)

type SubListRepositoryInterface interface {
	CreateSubList(request *entity.SubList) ([]entity.SubList, error)
	GetSubList(request *entity.SubList) ([]entity.SubList, error)
	UpdateSubList(request *entity.SubList) (entity.SubList, error)
}

func (repo *repository) CreateSubList(request *entity.SubList) ([]entity.SubList, error) {
	var sub_list []entity.SubList

	error := repo.db.Table("sub_list").Create(&request).Find(&sub_list).Error

	return sub_list, error
}

func (repo *repository) GetSubList(request *entity.SubList) ([]entity.SubList, error) {
	var sub_list []entity.SubList

	error := repo.db.Table("sub_list").Where("id_list = ?", request.IdList).Find(&sub_list).Error

	return sub_list, error
}

func (repo *repository) UpdateSubList(request *entity.SubList) (entity.SubList, error) {
	var sub_list entity.SubList

	error := repo.db.Table("sub_list").Model(&request).Updates(request).Find(&sub_list).Error

	return sub_list, error
}
