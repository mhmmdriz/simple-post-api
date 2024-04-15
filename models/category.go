package models

type Category struct {
	ID    int    `json:"id" gorm:"primary_key;"`
	Name  string `json:"name" gorm:"type:varchar(255);not null;"`
	Posts []Post `json:"posts" gorm:"foreignKey:CategoryID;references:ID;OnDelete:CASCADE;"`
}

type CategoryOnly struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
