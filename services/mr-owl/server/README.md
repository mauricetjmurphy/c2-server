# Project Layout

## Non Shared Layout

```
my-microservice-api/
│
├── cmd/
│   └── my-microservice-api/
│       ├── main.go
│       └── config.go
│
├── internal/
│   ├── api/
│   │   ├── v1/
│   │   │   ├── handler/
│   │   │   │   ├── user_handler.go
│   │   │   │   └── order_handler.go
│   │   │   ├── router/
│   │   │   │   └── router.go
│   │   │   └── models/
│   │   │       ├── user.go
│   │   │       └── order.go
│   │   ├── config/
│   │   │   └── config.go
│   │   ├── database/
│   │   │   ├── connection.go
│   │   │   ├── migrations/
│   │   │   │   └── ... (migration files)
│   │   │   └── seeds/
│   │   │       └── ... (seed files)
│   │   ├── middleware/
│   │   │   └── auth_middleware.go
│   │   ├── repositories/
│   │   │   ├── user_repository.go
│   │   │   └── order_repository.go
│   │   ├── services/
│   │   │   ├── user_service.go
│   │   │   └── order_service.go
│   │   └── utils/
│   │       └── util.go
│   ├── auth/
│   │   ├── jwt.go
│   │   └── oauth.go
│   └── ... (other internal packages)
│
├── scripts/
│   ├── build.sh
│   └── deploy.sh
│
├── test/
│   ├── integration/
│   │   └── ... (integration tests)
│   └── unit/
│       └── ... (unit tests)
│
├── Dockerfile
├── Makefile
└── README.md
```

- `cmd/my-microservice-api`: Contains the entry point for the application (main.go) and configuration setup (config.go).

- `internal/api/v1`: Contains versioned API handlers, routers, and models. The handler directory holds the business logic for handling API requests. The router directory contains the router configuration. The models directory holds the data models.

- `internal/config`: Configuration-related code.

- `internal/database`: Contains database connection setup, migrations, and seed data.

- `internal/middleware`: Middleware functions such as authentication and logging.

- `internal/repositories`: Repository layer for interacting with the database.

- `internal/services`: Business logic layer.

- `internal/utils`: Utility functions.

- `internal/auth`: Authentication mechanisms like JWT or OAuth.

- `scripts:` Scripts for building, deploying, and other automation tasks.

- `test`: Contains integration and unit tests.

- `Dockerfile`: Dockerfile for containerizing the application.

- `Makefile`: Makefile for automating tasks.

- `README.md`: Documentation for the project.

## Shared Layout

```
my-api/
│
├── cmd/
│   └── my-enterprise-api/
│       ├── main.go
│       └── config.go
│
├── pkg/
│   ├── api/
│   │   ├── v1/
│   │   │   ├── handler/
│   │   │   │   ├── user_handler.go
│   │   │   │   └── order_handler.go
│   │   │   ├── router/
│   │   │   │   └── router.go
│   │   │   └── models/
│   │   │       ├── user.go
│   │   │       └── order.go
│   │   ├── v2/  (future versions can be added here)
│   │   │   └── ... (similar structure as v1)
│   │
│   ├── config/
│   │   └── config.go
│   │
│   ├── database/
│   │   ├── connection.go
│   │   ├── migrations/
│   │   │   └── ... (migration files)
│   │   └── seeds/
│   │       └── ... (seed files)
│   │
│   ├── middleware/
│   │   └── auth_middleware.go
│   │
│   ├── repositories/
│   │   ├── user_repository.go
│   │   └── order_repository.go
│   │
│   ├── services/
│   │   ├── user_service.go
│   │   └── order_service.go
│   │
│   └── utils/
│       └── util.go
│
├── internal/
│   ├── auth/
│   │   ├── jwt.go
│   │   └── oauth.go
│   └── ... (other internal packages)
│
├── scripts/
│   ├── build.sh
│   └── deploy.sh
│
├── test/
│   ├── integration/
│   │   └── ... (integration tests)
│   └── unit/
│       └── ... (unit tests)
│
├── Dockerfile
├── Makefile
└── README.md
```



- `cmd/my-enterprise-api`: The entry point for the application. Contains `main.go` and any configuration initialization.

- `pkg/api/v1`: Contains versioned API handlers, routers, and models. The `handler` directory holds the business logic for handling API requests. The router directory contains the `router` configuration. The `models` directory holds the data models.

- `pkg/config`: Configuration-related code.

- `pkg/database`: Contains database connection setup, migrations, and seed data.

- `pkg/middleware`: Middleware functions such as authentication and logging.

- `pkg/repositories`: Repository layer for interacting with the database.

- `pkg/services`: Business logic layer.

- `pkg/utils`: Utility functions.

- `internal`: Internal packages that are not exposed outside the module. Contains authentication mechanisms like JWT or OAuth.

- `scripts`: Scripts for building, deploying, and other automation tasks.

- `test`: Contains integration and unit tests.

- `Dockerfile`: Dockerfile for containerizing the application.

- `Makefile`: Makefile for automating tasks.

- `README.md`: Documentation for the project.