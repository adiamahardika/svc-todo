package model

import "mime/multipart"

type CreateSubListRequest struct {
	IdList      int                   ` form:"idList" validate:"required"`
	Title       string                ` form:"title" validate:"required,alphanum,max=100"`
	Description string                ` form:"description" validate:"required,max=1000"`
	Attachment  *multipart.FileHeader ` form:"attachment"`
}
