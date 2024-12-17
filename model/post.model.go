package model

type Post struct {
	ID      int    `json:"id" gorm:"type:uuid;primary_key;autoIncrement"`
	Title   string `json:"title" gorm:"type:text"`
	Content string `json:"content" gorm:"type:text"`
	Image   string `json:"image" gorm:"type:text"`
}
