# Task Management System Microservice

## Overview
This project is a Task Management System microservice that allows users to create, read, update, and delete tasks. It includes features for pagination and filtering, making it easy to manage tasks efficiently.

## Features
- Create, Read, Update, Delete (CRUD) operations for tasks
- Pagination support for task lists
- Filtering capabilities to find tasks based on various criteria
- Authentication and authorization middleware

## Design Decisions
- The application is structured as a microservice to allow for scalability and maintainability.
- The use of Go for its performance and concurrency features.
- The separation of concerns through the use of different packages for API, models, repository, and services.

## Getting Started

### Prerequisites
- Go 1.16 or later
- A database (e.g., PostgreSQL, MySQL)

### Installation
1. Clone the repository:
   ```
   git clone <repository-url>
   cd task-management-service
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

### Running the Service
1. Set up your database and update the configuration in `internal/config/config.go`.
2. Run the application:
   ```
   go run cmd/main.go
   ```

### API Documentation
- **POST /tasks**: Create a new task
- **GET /tasks**: Retrieve a list of tasks with pagination and filtering
- **GET /tasks/{id}**: Retrieve a specific task by ID
- **PUT /tasks/{id}**: Update a specific task by ID
- **DELETE /tasks/{id}**: Delete a specific task by ID

## Conclusion
This microservice provides a robust solution for managing tasks and can be integrated into larger systems as needed. Further enhancements can include additional features such as user management and notifications.

A microservice for managing tasks with CRUD operations, pagination, and status filtering.

## Features
- Create, Read, Update, Delete tasks
- Pagination support
- Status filtering
- Soft delete
- RESTful API

## API Endpoints
- `POST /tasks` - Create a new task
- `GET /tasks` - List tasks with pagination and filtering
- `GET /tasks/{id}` - Get a specific task
- `PUT /tasks/{id}` - Update a task
- `DELETE /tasks/{id}` - Delete a task