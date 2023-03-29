package todo

type TodoDTO struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Title string `json:"title"`
}
