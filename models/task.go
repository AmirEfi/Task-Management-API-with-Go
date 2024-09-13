package controllers

type Task struct {
    ID          uint   `gorm:"primaryKey" json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Completed   bool   `json:"completed"`
    DueDate     string `json:"due_date"`
}
