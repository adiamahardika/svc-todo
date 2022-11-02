package model

import "mime/multipart"

type CreateListRequest struct {
	Title       string                `json:"title" form:"title" validate:"required,alphanum,max=100"`
	Description string                `json:"description" form:"description" validate:"required,max=1000"`
	Attachment  *multipart.FileHeader `json:"attachment" form:"attachment"`
}

type UpdateListRequest struct {
	Id          int                   `json:"id" form:"id"`
	Title       string                `json:"title" form:"title" validate:"required,alphanum,max=100"`
	Description string                `json:"description" form:"description" validate:"required,max=1000"`
	Attachment  *multipart.FileHeader `json:"attachment" form:"attachment"`
}

type GetListRequest struct {
	PageNo     int `query:"pageNo"`
	PageSize   int `query:"pageSize"`
	StartIndex int
}
