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
}

func (repo *repository) CreateList(request *entity.List) ([]entity.List, error) {
	var list []entity.List

	error := repo.db.Table("list").Create(&request).Find(&list).Error

	return list, error
}

func (repo *repository) CountList() (int64, error) {
	var count int64

	error := repo.db.Table("list").Count(&count).Error

	return count, error
}

func (repo *repository) GetList(request *model.GetListRequest) ([]entity.List, error) {
	var list []entity.List

	error := repo.db.Table("list").Limit(request.PageSize).Offset(request.StartIndex).Order("id").Find(&list).Error

	return list, error
}

func (repo *repository) GetListWithSub(request *model.GetListRequest) ([]entity.List, error) {
	var list []entity.List

	error := repo.db.Raw("SELECT list.*, JSON_AGG(sub_list.*) AS sub_list FROM list LEFT OUTER JOIN sub_list ON (list.id = sub_list.id_list) GROUP BY list.id ORDER BY list.id LIMIT @PageSize OFFSET @StartIndex", request).Find(&list).Error

	return list, error
}
