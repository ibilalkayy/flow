    ./
    ├── cmd
    │   ├── budget
    │   │   ├── budget.go
    │   │   └── handler
    │   │       ├── adjust.go
    │   │       ├── alert.go
    │   │       ├── create.go
    │   │       ├── get.go
    │   │       ├── msg.go
    │   │       ├── remove.go
    │   │       └── view.go
    │   └── transaction
    │       └── transaction.go
    ├── db
    │   └── budget_db
    │       ├── connection.go
    │       └── migrations
    │           └── 001_create_budget_table.sql
    ├── docs
    │   └── structure.md
    ├── internal
    │   ├── app
    │   │   └── budget
    │   │       └── budget.go
    │   ├── middleware
    │   │   └── env.go
    │   └── structs
    │       └── common.go
    ├── tests
    │   └── app
    │       └── budget
    │           └── budget_test.go
    ├── .gitignore
    ├── .dockerignore
    ├── CODE_OF_CONDUCT.md
    ├── CONTRIBUTING.md
    ├── LICENSE
    ├── Dockerfile
    ├── README.md
    ├── docker-compose.yml
    ├── go.mod
    ├── go.sum
    ├── main.go

## **File Details:**

- **cmd/budget/budget.go:** Budget service entry point.
- **cmd/budget/handler/adjust.go:** Handler for adjusting budget.
- **cmd/budget/handler/alert.go:** Handler for alerting once the budget limit is passed.
- **cmd/budget/handler/create.go:** Handler for creating budget.
- **cmd/budget/handler/get.go:** Handler for getting budget.
- **cmd/budget/handler/msg.go:** Handler for showing the alert messages.
- **cmd/budget/handler/remove.go:** Handler for removing budget.
- **cmd/budget/handler/view.go:** Handler for viewing budget.
- **cmd/transaction/transaction.go:** Transaction service entry point.
- **db/budget_db/connection.go:** Database connection setup.
- **db/budget_db/migrations/001_create_budget_table.sql:** SQL script for creating budget table.
- **docs/structure.md:** Structure of the whole application.
- **internal/app/budget/budget.go:** Implementation of budget management functionality.
- **internal/middleware/env.go:** Environment middleware for handling environment variables.
- **tests/app/budget/budget_test.go:** Test file for handling the budget test functions.
- **.gitignore:** Specifies intentionally untracked files that Git should ignore.
- **.dockerignore:** Specifies files and directories that should be ignored when building Docker images.
- **CODE_OF_CONDUCT.md:** Code of conduct for contributors.
- **CONTRIBUTING.md:** Guidelines for contributing to the project.
- **Dockerfile:** Defines instructions to build Docker image of the application.
- **README.md:** Main documentation file providing information about the project.
- **docker-compose.yml:** Configuration file for Docker Compose.
- **go.mod:** Go module file specifying dependencies.
- **go.sum:** Go module file specifying exact versions of dependencies.
- **main.go:** Main entry point of the application.