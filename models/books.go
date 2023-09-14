package models

type Books struct {
	ID        uint    `gorm:"primary_key;autoIncrement" json:"Id"`
	Author    *string `json:"author"`
	Title     *string `json:"title"`
	Publisher *string `json:"publisher"`
}
