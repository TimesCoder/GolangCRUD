// models/todo.go
package models

type Todo struct {
    ID     uint   `json:"id" gorm:"primaryKey"`
    Title  string `json:"title"`
    Status string `json:"status"`
}
