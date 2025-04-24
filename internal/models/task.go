package models

import "time"

// Task represents a task in the task management system.
type Task struct {
	ID          string    `json:"id"`          // Unique identifier for the task
	Title       string    `json:"title"`       // Title of the task
	Description string    `json:"description"` // Detailed description of the task
	Status      string    `json:"status"`      // Current status of the task (e.g., "pending", "completed")
	CreatedAt   time.Time `json:"created_at"`  // Timestamp when the task was created
	UpdatedAt   time.Time `json:"updated_at"`  // Timestamp when the task was last updated
}