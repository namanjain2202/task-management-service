package main

import (
    "database/sql"
    "log"
    "net/http"
    
    "task-management-service/internal/api/handlers"
    "task-management-service/internal/api/routes"
    "task-management-service/internal/config"
    "task-management-service/internal/repository"
    "task-management-service/internal/service"
    
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // Load configuration settings
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("could not load config: %v", err)
    }

    // Initialize database connection
    db, err := sql.Open("mysql", cfg.DatabaseURL)
    if err != nil {
        log.Fatalf("could not connect to database: %v", err)
    }
    defer db.Close()

    // Initialize repository
    taskRepo := repository.NewTaskRepository(db)

    // Initialize service
    taskService := service.NewTaskService(taskRepo)

    // Initialize handler
    taskHandler := handlers.NewTaskHandler(taskService)

    // Set up routes
    router := routes.SetupRouter(taskHandler)

    // Start the server
    log.Printf("Starting server on port %s...", cfg.ServerPort)
    if err := http.ListenAndServe(cfg.ServerPort, router); err != nil {
        log.Fatalf("could not start server: %v", err)
    }
}
