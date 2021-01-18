package entities

// Task reflects tasks in DB
type Task struct {
	ID          string `json:"id" gorm:"column:id"`
	Content     string `json:"content" gorm:"column:content"`
	UserID      string `json:"user_id" gorm:"column:user_id"`
	CreatedDate string `json:"created_date" gorm:"column:created_date"`
}
