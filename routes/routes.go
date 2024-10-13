// routes/routes.go
package routes

import (
    "github.com/labstack/echo/v4"
    "gorm.io/gorm"
    "weekly-task-crud/handlers"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
    e.GET("/todos", handlers.GetTodos(db))
    e.GET("/todos/:id", handlers.GetTodoByID(db))
    e.POST("/todos", handlers.CreateTodo(db))
    e.PUT("/todos/:id", handlers.UpdateTodo(db))
    e.DELETE("/todos/:id", handlers.DeleteTodo(db))
}
