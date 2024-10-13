package main

import (
    "log"                         // Untuk logging
    "os"                          // Untuk output log ke terminal
    "time"                        // Untuk SlowThreshold
    "gorm.io/gorm/logger"         // Untuk logger Gorm
    "github.com/labstack/echo/v4" // Echo framework
    "gorm.io/driver/sqlite"       // SQLite driver untuk Gorm
    "gorm.io/gorm"                // Gorm ORM
    "weekly-task-crud/routes"     // Routes yang kamu definisikan
    "weekly-task-crud/models"     // Models yang berisi struct Todo
)

func main() {
    // Inisialisasi Echo
    e := echo.New()

    // Inisialisasi database SQLite3 dengan Logger
    db, err := gorm.Open(sqlite.Open("database/todo.db"), &gorm.Config{
        Logger: logger.New(
            log.New(os.Stdout, "\r\n", log.LstdFlags), // Output ke console
            logger.Config{
                SlowThreshold: time.Second,   // Menampilkan query yang lambat
                LogLevel:      logger.Info,   // Level log Info untuk menampilkan semua query
                Colorful:      true,          // Output warna-warni
            },
        ),
    })
    if err != nil {
        log.Fatal("Gagal menghubungkan ke database SQLite3")
    }

    // Migrate model Todo
    db.AutoMigrate(&models.Todo{}) 

    // Setup routes
    routes.SetupRoutes(e, db)

    // Jalankan server di localhost
    e.Logger.Fatal(e.Start(":8080"))
}
