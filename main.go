// main.go
package main

import (
    "log"
    "github.com/labstack/echo/v4"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "weekly-task-crud/routes"
    
)

func main() {
    // Inisialisasi Echo
    e := echo.New()

    // Inisialisasi database
    db, err := gorm.Open(sqlite.Open("database/todo.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database")
    }

    // Setup routes
    routes.SetupRoutes(e, db)

    // Jalankan server di localhost
    e.Logger.Fatal(e.Start(":8080"))
}
