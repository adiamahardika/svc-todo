package repository

import (
	"svc-todo/entity"
	"svc-todo/model"
)

type ListRepositoryInterface interface {
	CreateList(request *entity.List) ([]entity.List, error)
	CountList() (int64, error)
	GetList(request *model.GetListRequest) ([]entity.List, error)
	GetListWithSub(request *model.GetListRequest) ([]entity.List, error)
	UpdateList(request *entity.List) (entity.List, error)
	DeleteList(request *int) error
}

func (repo *repository) CreateList(request *entity.List) ([]entity.List, error) {
	var list []entity.List

	error := repo.db.Table("lists").Create(&request).Find(&list).Error

	return list, error
}

func (repo *repository) CountList() (int64, error) {
	var count int64

	error := repo.db.Table("lists").Count(&count).Error

	return count, error
}

func (repo *repository) GetList(request *model.GetListRequest) ([]entity.List, error) {
	var list []entity.List

	error := repo.db.Table("lists").Where("is_active = ?", "true").Limit(request.PageSize).Offset(request.StartIndex).Order("id").Find(&list).Error

	return list, error
}

func (repo *repository) GetListWithSub(request *model.GetListRequest) ([]entity.List, error) {
	var list []entity.List

	error := repo.db.Raw("SELECT lists.*, JSON_AGG(sub_lists.*) AS sub_lists FROM lists LEFT OUTER JOIN sub_lists ON (lists.id = sub_lists.id_list) WHERE lists.is_active = 'true' GROUP BY lists.id ORDER BY lists.id LIMIT @PageSize OFFSET @StartIndex", request).Find(&list).Error

	return list, error
}

func (repo *repository) UpdateList(request *entity.List) (entity.List, error) {
	var list entity.List

	error := repo.db.Table("lists").Model(&request).Updates(request).Find(&list).Error

	return list, error
}

func (repo *repository) DeleteList(request *int) error {
	var list entity.List

	error := repo.db.Table("lists").Model(&entity.List{Id: *request}).Update("is_active", "false").Find(&list).Error

	return error
}
