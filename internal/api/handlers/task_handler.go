package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "task-management-service/internal/models"
    "task-management-service/internal/service"
)

// TaskHandler handles HTTP requests related to tasks.
type TaskHandler struct {
    TaskService service.TaskService
}

// NewTaskHandler creates a new TaskHandler.
func NewTaskHandler(taskService service.TaskService) *TaskHandler {
    return &TaskHandler{TaskService: taskService}
}

// CreateTask handles the creation of a new task.
func (h *TaskHandler) CreateTask(c *gin.Context) {
    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    createdTask, err := h.TaskService.CreateTask(task)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, createdTask)
}

// GetTasks handles retrieving a list of tasks with pagination and filtering.
func (h *TaskHandler) GetTasks(c *gin.Context) {
    // Implement pagination and filtering logic here
    tasks, err := h.TaskService.GetTasks()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, tasks)
}

// UpdateTask handles updating an existing task.
func (h *TaskHandler) UpdateTask(c *gin.Context) {
    id := c.Param("id")
    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    updatedTask, err := h.TaskService.UpdateTask(id, task)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, updatedTask)
}

// DeleteTask handles deleting a task.
func (h *TaskHandler) DeleteTask(c *gin.Context) {
    id := c.Param("id")
    err := h.TaskService.DeleteTask(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusNoContent, nil)
}