package models

import (
    "errors"
    "time"
)

type Task struct {
    ID          uint       `json:"id"`
    Title       string     `json:"title"`
    Description string     `json:"description"`
    Status      string     `json:"status"`
    CreatedAt   time.Time  `json:"created_at"`
    UpdatedAt   time.Time  `json:"updated_at"`
    DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

func (t *Task) Validate() error {
    if t.Title == "" {
        return errors.New("title is required")
    }
    if !isValidStatus(t.Status) {
        return errors.New("invalid status")
    }
    return nil
}

func (t *Task) Update(updates map[string]interface{}) error {
    if status, ok := updates["status"].(string); ok {
        if !isValidStatus(status) {
            return errors.New("invalid status")
        }
        t.Status = status
    }
    if title, ok := updates["title"].(string); ok {
        t.Title = title
    }
    if description, ok := updates["description"].(string); ok {
        t.Description = description
    }
    t.UpdatedAt = time.Now()
    return nil
}

func (t *Task) Delete() error {
    if t.DeletedAt != nil {
        return errors.New("task already deleted")
    }
    now := time.Now()
    t.DeletedAt = &now
    return nil
}

func (t *Task) Get() (*Task, error) {
    if t.DeletedAt != nil {
        return nil, errors.New("task not found")
    }
    return t, nil
}

func isValidStatus(status string) bool {
    validStatuses := []string{"pending", "completed", "in_progress"}
    for _, s := range validStatuses {
        if s == status {
            return true
        }
    }
    return false
}
