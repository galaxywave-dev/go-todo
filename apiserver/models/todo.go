package models

type Todo struct {
	ID    uint32 `json:"id" gorm:"primary_key"`
	Title string `json:"title"`
}
