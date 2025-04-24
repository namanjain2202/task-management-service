package main

import (
    "log"
    "net/http"
    "task-management-service/internal/api/routes"
    "task-management-service/internal/config"
)

func main() {
    // Load configuration settings
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("could not load config: %v", err)
    }

    // Set up routes
    router := routes.SetupRouter()

    // Start the server
    log.Printf("Starting server on port %s...", cfg.ServerPort)
    if err := http.ListenAndServe(cfg.ServerPort, router); err != nil {
        log.Fatalf("could not start server: %v", err)
    }
}