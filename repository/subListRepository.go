package repository

import (
	"svc-todo/entity"
)

type SubListRepositoryInterface interface {
	CreateSubList(request *entity.SubList) ([]entity.SubList, error)
	GetSubList(request *entity.SubList) ([]entity.SubList, error)
	UpdateSubList(request *entity.SubList) (entity.SubList, error)
	DeleteSubList(request *int) error
}

func (repo *repository) CreateSubList(request *entity.SubList) ([]entity.SubList, error) {
	var sub_list []entity.SubList

	error := repo.db.Table("sub_lists").Create(&request).Find(&sub_list).Error

	return sub_list, error
}

func (repo *repository) GetSubList(request *entity.SubList) ([]entity.SubList, error) {
	var sub_list []entity.SubList

	error := repo.db.Table("sub_lists").Where("is_active = 'true' AND id_list = ?", request.IdList).Find(&sub_list).Error

	return sub_list, error
}

func (repo *repository) UpdateSubList(request *entity.SubList) (entity.SubList, error) {
	var sub_list entity.SubList

	error := repo.db.Table("sub_lists").Model(&request).Updates(request).Find(&sub_list).Error

	return sub_list, error
}

func (repo *repository) DeleteSubList(request *int) error {
	var sub_list entity.SubList

	error := repo.db.Table("sub_lists").Model(&entity.SubList{Id: *request}).Update("is_active", "false").Find(&sub_list).Error

	return error
}
