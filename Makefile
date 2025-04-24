task-management-service
├── cmd
│   └── main.go
├── internal
│   ├── api
│   │   ├── handlers
│   │   │   └── task_handler.go
│   │   ├── middleware
│   │   │   └── auth.go
│   │   └── routes.go
│   ├── config
│   │   └── config.go
│   ├── models
│   │   └── task.go
│   ├── repository
│   │   └── task_repository.go
│   └── service
│       └── task_service.go
├── pkg
│   ├── errors
│   │   └── errors.go
│   └── utils
│       └── pagination.go
├── go.mod
├── go.sum
├── Makefile
└── README.md