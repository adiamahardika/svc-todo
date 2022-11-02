package entity

type List struct {
	Id          int    `json:"id" gorm:"index:idx_name,unique,primaryKey"`
	Title       string `json:"title" gorm:"size:100"`
	Description string `json:"description"`
	Attachment  string `json:"attachment"`
	SubList     string `json:"subList" gorm:"->"`
	IsActive    string `json:"isActive"`
}
