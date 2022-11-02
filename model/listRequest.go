package model

import "mime/multipart"

type CreateListRequest struct {
	Title       string                `form:"title" validate:"required,alphanum,max=100"`
	Description string                `form:"description" validate:"required,max=1000"`
	Attachment  *multipart.FileHeader `form:"attachment"`
}

type UpdateListRequest struct {
	Id          int                   `form:"id"`
	Title       string                `form:"title" validate:"required,alphanum,max=100"`
	Description string                `form:"description" validate:"required,max=1000"`
	Attachment  *multipart.FileHeader `form:"attachment"`
}

type GetListRequest struct {
	PageNo     int `query:"pageNo"`
	PageSize   int `query:"pageSize"`
	StartIndex int
}
