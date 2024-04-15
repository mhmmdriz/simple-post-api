package models

type Post struct {
	ID         int      `json:"id" gorm:"primary_key;"`
	CategoryID int      `json:"category_id"`
	Title      string   `json:"title" gorm:"type:varchar(255);not null;"`
	Content    string   `json:"content" gorm:"type:text;not null;"`
	Category   Category `json:"-" gorm:"foreignKey:CategoryID;references:ID;OnDelete:CASCADE;"`
}

type PostResponse struct {
	ID           int    `json:"id"`
	CategoryID   int    `json:"category_id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	CategoryName string `json:"category_name"`
}
