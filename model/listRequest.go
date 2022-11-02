package model

import "mime/multipart"

type ListRequest struct {
	Title       string                `json:"title" form:"title" validate:"required,alphanum,max=100"`
	Description string                `json:"description" form:"description" validate:"required,max=1000"`
	Attachment  *multipart.FileHeader `json:"attachment" form:"attachment"`
}
