// handlers/todo.go
package handlers

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "gorm.io/gorm"
    "strconv"
    "weekly-task-crud/models"
)

// Get all todos
func GetTodos(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        var todos []models.Todo
        result := db.Find(&todos)
        if result.Error != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{
                "error": "Failed to retrieve todos",
            })
        }
        return c.JSON(http.StatusOK, todos)
    }
}

// Get a todo by ID
func GetTodoByID(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{
                "error": "Invalid todo ID",
            })
        }

        var todo models.Todo
        result := db.First(&todo, id)
        if result.Error != nil {
            if result.Error == gorm.ErrRecordNotFound {
                return c.JSON(http.StatusNotFound, map[string]string{
                    "error": "Todo not found",
                })
            }
            return c.JSON(http.StatusInternalServerError, map[string]string{
                "error": "Failed to retrieve todo",
            })
        }

        return c.JSON(http.StatusOK, todo)
    }
}

// Create a new todo
func CreateTodo(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        todo := new(models.Todo)
        if err := c.Bind(todo); err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{
                "error": "Invalid data format",
            })
        }

        // Validasi data
        if todo.Title == "" || todo.Status == "" {
            return c.JSON(http.StatusBadRequest, map[string]string{
                "error": "Title and Status are required",
            })
        }

        result := db.Create(todo)
        if result.Error != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{
                "error": "Failed to create todo",
            })
        }

        return c.JSON(http.StatusCreated, todo)
    }
}
// Update a todo by ID
func UpdateTodo(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{
                "error": "Invalid todo ID",
            })
        }

        var todo models.Todo
        result := db.First(&todo, id)
        if result.Error != nil {
            if result.Error == gorm.ErrRecordNotFound {
                return c.JSON(http.StatusNotFound, map[string]string{
                    "error": "Todo not found",
                })
            }
            return c.JSON(http.StatusInternalServerError, map[string]string{
                "error": "Failed to retrieve todo",
            })
        }

        if err := c.Bind(&todo); err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{
                "error": "Invalid data format",
            })
        }

        db.Save(&todo)
        return c.JSON(http.StatusOK, todo)
    }
}

// Delete a todo by ID
func DeleteTodo(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{
                "error": "Invalid todo ID",
            })
        }

        var todo models.Todo
        result := db.First(&todo, id)
        if result.Error != nil {
            if result.Error == gorm.ErrRecordNotFound {
                return c.JSON(http.StatusNotFound, map[string]string{
                    "error": "Todo not found",
                })
            }
            return c.JSON(http.StatusInternalServerError, map[string]string{
                "error": "Failed to retrieve todo",
            })
        }

        db.Delete(&todo)
        return c.JSON(http.StatusOK, map[string]string{
            "message": "Todo deleted",
        })
    }
}
