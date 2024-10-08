// handlers/todo.go
package handlers

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "gorm.io/gorm"
    "strconv"
    "weekly-task-crud/models" // Ganti import sesuai package models
)

// Get all todos
func GetTodos(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        var todos []models.Todo // Perbarui menjadi models.Todo
        db.Find(&todos)
        return c.JSON(http.StatusOK, todos)
    }
}

// Create a new todo
func CreateTodo(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        todo := new(models.Todo) // Perbarui menjadi models.Todo
        if err := c.Bind(todo); err != nil {
            return c.JSON(http.StatusBadRequest, "Invalid data")
        }
        db.Create(todo)
        return c.JSON(http.StatusCreated, todo)
    }
}

// Update a todo by ID
func UpdateTodo(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        id, _ := strconv.Atoi(c.Param("id"))
        var todo models.Todo // Perbarui menjadi models.Todo
        if result := db.First(&todo, id); result.Error != nil {
            return c.JSON(http.StatusNotFound, "Todo not found")
        }

        if err := c.Bind(&todo); err != nil {
            return c.JSON(http.StatusBadRequest, "Invalid data")
        }

        db.Save(&todo)
        return c.JSON(http.StatusOK, todo)
    }
}

// Delete a todo by ID
func DeleteTodo(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        id, _ := strconv.Atoi(c.Param("id"))
        var todo models.Todo // Perbarui menjadi models.Todo
        if result := db.First(&todo, id); result.Error != nil {
            return c.JSON(http.StatusNotFound, "Todo not found")
        }

        db.Delete(&todo)
        return c.JSON(http.StatusOK, "Todo deleted")
    }
}
