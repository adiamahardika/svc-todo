package entity

type SubList struct {
	Id          int    `json:"id" gorm:"index:idx_name,unique,primaryKey"`
	IdList      int    `json:"idList" query:"idList"`
	Title       string `json:"title" gorm:"size:255"`
	Description string `json:"description"`
	Attachment  string `json:"attachment"`
	IsActive    string `json:"isActive"`
}
