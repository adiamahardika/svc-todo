package entity

type SubList struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	IdList      int    `json:"idList" query:"idList"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Attachment  string `json:"attachment"`
}
