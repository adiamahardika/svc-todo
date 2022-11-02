package entity

type List struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Attachment  string `json:"attachment"`
}
