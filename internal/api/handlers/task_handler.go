package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    
    "task-management-service/internal/models"
    "task-management-service/internal/service"
    "task-management-service/pkg/utils"
    
    "github.com/gorilla/mux"
)

type TaskHandler struct {
    taskService service.TaskService
}

func NewTaskHandler(service service.TaskService) *TaskHandler {
    return &TaskHandler{taskService: service}
}

// CreateTask handles task creation
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
    var task models.Task
    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    
    if err := h.taskService.CreateTask(r.Context(), &task); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(task)
}

// GetTask handles single task retrieval
func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.ParseUint(vars["id"], 10, 64)
    if err != nil {
        http.Error(w, "Invalid task ID", http.StatusBadRequest)
        return
    }
    
    task, err := h.taskService.GetTask(r.Context(), uint(id))
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    
    json.NewEncoder(w).Encode(task)
}

// ListTasks handles task listing with pagination and filtering
func (h *TaskHandler) ListTasks(w http.ResponseWriter, r *http.Request) {
    page, _ := strconv.Atoi(r.URL.Query().Get("page"))
    if page < 1 {
        page = 1
    }
    
    pageSize, _ := strconv.Atoi(r.URL.Query().Get("page_size"))
    if pageSize < 1 {
        pageSize = 10
    }
    
    status := r.URL.Query().Get("status")
    
    tasks, total, err := h.taskService.ListTasks(r.Context(), page, pageSize, status)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    response := map[string]interface{}{
        "tasks": tasks,
        "pagination": utils.Pagination{
            Page:       page,
            PageSize:   pageSize,
            TotalItems: total,
        },
    }
    
    json.NewEncoder(w).Encode(response)
}

// UpdateTask handles task updates
func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.ParseUint(vars["id"], 10, 64)
    if err != nil {
        http.Error(w, "Invalid task ID", http.StatusBadRequest)
        return
    }
    
    var task models.Task
    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    task.ID = uint(id)
    
    if err := h.taskService.UpdateTask(r.Context(), &task); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    json.NewEncoder(w).Encode(task)
}

// DeleteTask handles task deletion
func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.ParseUint(vars["id"], 10, 64)
    if err != nil {
        http.Error(w, "Invalid task ID", http.StatusBadRequest)
        return
    }
    
    if err := h.taskService.DeleteTask(r.Context(), uint(id)); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.WriteHeader(http.StatusNoContent)
}
