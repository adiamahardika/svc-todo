package model

import (
	"svc-todo/entity"
)

type GetListResponse struct {
	Id          int               `json:"id" gorm:"primaryKey"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Attachment  string            `json:"attachment"`
	SubList     []*entity.SubList `json:"subList"`
}
