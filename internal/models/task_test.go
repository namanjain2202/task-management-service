package models

import (
    "testing"
    "time"
)

func TestCreateTask(t *testing.T) {
    tests := []struct {
        name          string
        title         string
        description   string
        status        string
        expectedError bool
    }{
        {
            name:          "Valid task",
            title:         "Test task",
            description:   "Test description",
            status:        "pending",
            expectedError: false,
        },
        {
            name:          "Empty title",
            title:         "",
            description:   "Test description",
            status:        "pending",
            expectedError: true,
        },
        {
            name:          "Invalid status",
            title:         "Test task",
            description:   "Test description",
            status:        "invalid_status",
            expectedError: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            task := &Task{
                Title:       tt.title,
                Description: tt.description,
                Status:      tt.status,
                CreatedAt:   time.Now(),
                UpdatedAt:   time.Now(),
            }

            err := task.Validate()
            if (err != nil) != tt.expectedError {
                t.Errorf("Task.Validate() error = %v, expectedError %v", err, tt.expectedError)
            }
        })
    }
}

func TestUpdateTask(t *testing.T) {
    tests := []struct {
        name          string
        originalTask  *Task
        updates       map[string]interface{}
        expectedError bool
    }{
        {
            name: "Valid update",
            originalTask: &Task{
                Title:       "Original title",
                Description: "Original description",
                Status:      "pending",
            },
            updates: map[string]interface{}{
                "title":       "Updated title",
                "description": "Updated description",
                "status":      "completed",
            },
            expectedError: false,
        },
        {
            name: "Invalid status update",
            originalTask: &Task{
                Title:       "Original title",
                Description: "Original description",
                Status:      "pending",
            },
            updates: map[string]interface{}{
                "status": "invalid_status",
            },
            expectedError: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := tt.originalTask.Update(tt.updates)
            if (err != nil) != tt.expectedError {
                t.Errorf("Task.Update() error = %v, expectedError %v", err, tt.expectedError)
            }
        })
    }
}

func TestDeleteTask(t *testing.T) {
    tests := []struct {
        name          string
        task          *Task
        expectedError bool
    }{
        {
            name: "Valid deletion",
            task: &Task{
                ID:          1,
                Title:       "Test task",
                Description: "Test description",
                Status:      "pending",
            },
            expectedError: false,
        },
        {
            name: "Already deleted task",
            task: &Task{
                ID:        1,
                DeletedAt: &time.Time{},
            },
            expectedError: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := tt.task.Delete()
            if (err != nil) != tt.expectedError {
                t.Errorf("Task.Delete() error = %v, expectedError %v", err, tt.expectedError)
            }
        })
    }
}

func TestGetTask(t *testing.T) {
    tests := []struct {
        name          string
        task          *Task
        expectedError bool
    }{
        {
            name: "Valid task",
            task: &Task{
                ID:          1,
                Title:       "Test task",
                Description: "Test description",
                Status:      "pending",
            },
            expectedError: false,
        },
        {
            name: "Deleted task",
            task: &Task{
                ID:        1,
                DeletedAt: &time.Time{},
            },
            expectedError: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := tt.task.Get()
            if (err != nil) != tt.expectedError {
                t.Errorf("Task.Get() error = %v, expectedError %v", err, tt.expectedError)
            }

            if !tt.expectedError {
                if result.ID != tt.task.ID {
                    t.Errorf("Task.ID = %v, want %v", result.ID, tt.task.ID)
                }
                if result.Title != tt.task.Title {
                    t.Errorf("Task.Title = %v, want %v", result.Title, tt.task.Title)
                }
                if result.Description != tt.task.Description {
                    t.Errorf("Task.Description = %v, want %v", result.Description, tt.task.Description)
                }
                if result.Status != tt.task.Status {
                    t.Errorf("Task.Status = %v, want %v", result.Status, tt.task.Status)
                }
            }
        })
    }
}
