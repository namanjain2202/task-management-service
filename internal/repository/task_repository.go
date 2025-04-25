package repository

import (
    "context"
    "database/sql"
    "errors"
    "fmt"
    "time"
    
    "task-management-service/internal/models"
    customErrors "task-management-service/pkg/errors"
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
func (r *TaskRepository) CreateTask(ctx context.Context, task *models.Task) error {
    query := `
        INSERT INTO tasks (title, description, status, created_at, updated_at) 
        VALUES (?, ?, ?, ?, ?)
    `
    now := time.Now()
    task.CreatedAt = now
    task.UpdatedAt = now

    result, err := r.db.ExecContext(ctx, query, 
        task.Title, 
        task.Description, 
        task.Status,
        task.CreatedAt,
        task.UpdatedAt,
    )
    if err != nil {
        return fmt.Errorf("error creating task: %w", err)
    }

    id, err := result.LastInsertId()
    if err != nil {
        return fmt.Errorf("error getting last insert id: %w", err)
    }

    task.ID = uint(id)
    return nil
}

// GetTask retrieves a task by its ID
func (r *TaskRepository) GetTask(ctx context.Context, id uint) (*models.Task, error) {
    task := &models.Task{}
    query := `
        SELECT id, title, description, status, created_at, updated_at, deleted_at 
        FROM tasks 
        WHERE id = ? AND deleted_at IS NULL
    `
    
    err := r.db.QueryRowContext(ctx, query, id).Scan(
        &task.ID,
        &task.Title,
        &task.Description,
        &task.Status,
        &task.CreatedAt,
        &task.UpdatedAt,
        &task.DeletedAt,
    )

    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, customErrors.NewNotFoundError("task not found")
        }
        return nil, fmt.Errorf("error getting task: %w", err)
    }

    return task, nil
}

// UpdateTask updates an existing task in the database
func (r *TaskRepository) UpdateTask(ctx context.Context, task *models.Task) error {
    query := `
        UPDATE tasks 
        SET title = ?, description = ?, status = ?, updated_at = ? 
        WHERE id = ? AND deleted_at IS NULL
    `
    
    task.UpdatedAt = time.Now()
    result, err := r.db.ExecContext(ctx, query,
        task.Title,
        task.Description,
        task.Status,
        task.UpdatedAt,
        task.ID,
    )
    if err != nil {
        return fmt.Errorf("error updating task: %w", err)
    }

    rows, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error getting rows affected: %w", err)
    }

    if rows == 0 {
        return customErrors.NewNotFoundError("task not found")
    }

    return nil
}

// DeleteTask soft deletes a task from the database
func (r *TaskRepository) DeleteTask(ctx context.Context, id uint) error {
    query := `
        UPDATE tasks 
        SET deleted_at = ? 
        WHERE id = ? AND deleted_at IS NULL
    `
    
    result, err := r.db.ExecContext(ctx, query, time.Now(), id)
    if err != nil {
        return fmt.Errorf("error deleting task: %w", err)
    }

    rows, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error getting rows affected: %w", err)
    }

    if rows == 0 {
        return customErrors.NewNotFoundError("task not found")
    }

    return nil
}

// ListTasks retrieves tasks with pagination and filtering
func (r *TaskRepository) ListTasks(ctx context.Context, page, pageSize int, status string) ([]models.Task, int, error) {
    tasks := []models.Task{}
    
    // Build query with optional status filter
    whereClause := "WHERE deleted_at IS NULL"
    args := []interface{}{}
    
    if status != "" {
        whereClause += " AND status = ?"
        args = append(args, status)
    }
    
    // Get total count for pagination
    countQuery := fmt.Sprintf("SELECT COUNT(*) FROM tasks %s", whereClause)
    var total int
    err := r.db.QueryRowContext(ctx, countQuery, args...).Scan(&total)
    if err != nil {
        return nil, 0, fmt.Errorf("error counting tasks: %w", err)
    }

    // Get paginated results
    query := fmt.Sprintf(`
        SELECT id, title, description, status, created_at, updated_at 
        FROM tasks 
        %s 
        ORDER BY created_at DESC 
        LIMIT ? OFFSET ?
    `, whereClause)
    
    args = append(args, pageSize, (page-1)*pageSize)
    rows, err := r.db.QueryContext(ctx, query, args...)
    if err != nil {
        return nil, 0, fmt.Errorf("error listing tasks: %w", err)
    }
    defer rows.Close()

    for rows.Next() {
        var task models.Task
        if err := rows.Scan(
            &task.ID,
            &task.Title,
            &task.Description,
            &task.Status,
            &task.CreatedAt,
            &task.UpdatedAt,
        ); err != nil {
            return nil, 0, fmt.Errorf("error scanning task: %w", err)
        }
        tasks = append(tasks, task)
    }

    return tasks, total, nil
}
