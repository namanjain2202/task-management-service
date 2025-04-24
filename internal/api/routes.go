package api

import (
    "github.com/gorilla/mux"
    "net/http"
)

// RegisterRoutes sets up the API routes and associates them with their respective handler methods.
func RegisterRoutes(r *mux.Router) {
    r.HandleFunc("/tasks", CreateTask).Methods(http.MethodPost)
    r.HandleFunc("/tasks", GetTasks).Methods(http.MethodGet)
    r.HandleFunc("/tasks/{id}", GetTask).Methods(http.MethodGet)
    r.HandleFunc("/tasks/{id}", UpdateTask).Methods(http.MethodPut)
    r.HandleFunc("/tasks/{id}", DeleteTask).Methods(http.MethodDelete)
}