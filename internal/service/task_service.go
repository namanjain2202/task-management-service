package service

import (
    "errors"
    "time"

    "task-management-service/internal/models"
    "task-management-service/internal/repository"
)

// TaskService provides methods for managing tasks.
type TaskService struct {
    repo repository.TaskRepository
}

// NewTaskService creates a new TaskService.
func NewTaskService(repo repository.TaskRepository) *TaskService {
    return &TaskService{repo: repo}
}

// CreateTask creates a new task in the system.
func (s *TaskService) CreateTask(task *models.Task) error {
    if task.Title == "" {
        return errors.New("task title cannot be empty")
    }
    task.CreatedAt = time.Now()
    return s.repo.Create(task)
}

// GetTask retrieves a task by its ID.
func (s *TaskService) GetTask(id string) (*models.Task, error) {
    return s.repo.FindByID(id)
}

// UpdateTask updates an existing task.
func (s *TaskService) UpdateTask(task *models.Task) error {
    if task.ID == "" {
        return errors.New("task ID cannot be empty")
    }
    return s.repo.Update(task)
}

// DeleteTask removes a task from the system.
func (s *TaskService) DeleteTask(id string) error {
    return s.repo.Delete(id)
}

// ListTasks retrieves a list of tasks with pagination and filtering.
func (s *TaskService) ListTasks(page, limit int, filter string) ([]models.Task, error) {
    return s.repo.FindAll(page, limit, filter)
}