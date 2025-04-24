package repository

import (
    "database/sql"
    "errors"
    "task-management-service/internal/models"
)

// TaskRepository struct for interacting with the database
type TaskRepository struct {
    db *sql.DB
}

// NewTaskRepository creates a new TaskRepository
func NewTaskRepository(db *sql.DB) *TaskRepository {
    return &TaskRepository{db: db}
}

// CreateTask inserts a new task into the database
func (r *TaskRepository) CreateTask(task *models.Task) error {
    query := "INSERT INTO tasks (title, description, status) VALUES (?, ?, ?)"
    result, err := r.db.Exec(query, task.Title, task.Description, task.Status)
    if err != nil {
        return err
    }
    id, err := result.LastInsertId()
    if err != nil {
        return err
    }
    task.ID = id
    return nil
}

// GetTask retrieves a task by its ID
func (r *TaskRepository) GetTask(id int64) (*models.Task, error) {
    task := &models.Task{}
    query := "SELECT id, title, description, status, created_at, updated_at FROM tasks WHERE id = ?"
    err := r.db.QueryRow(query, id).Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, nil // Task not found
        }
        return nil, err
    }
    return task, nil
}

// UpdateTask updates an existing task in the database
func (r *TaskRepository) UpdateTask(task *models.Task) error {
    query := "UPDATE tasks SET title = ?, description = ?, status = ? WHERE id = ?"
    _, err := r.db.Exec(query, task.Title, task.Description, task.Status, task.ID)
    return err
}

// DeleteTask removes a task from the database
func (r *TaskRepository) DeleteTask(id int64) error {
    query := "DELETE FROM tasks WHERE id = ?"
    _, err := r.db.Exec(query, id)
    return err
}

// GetAllTasks retrieves all tasks with pagination
func (r *TaskRepository) GetAllTasks(limit, offset int) ([]models.Task, error) {
    tasks := []models.Task{}
    query := "SELECT id, title, description, status, created_at, updated_at FROM tasks LIMIT ? OFFSET ?"
    rows, err := r.db.Query(query, limit, offset)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var task models.Task
        if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt); err != nil {
            return nil, err
        }
        tasks = append(tasks, task)
    }
    return tasks, nil
}