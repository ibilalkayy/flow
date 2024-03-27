    ./
    ├── cmd
    │   ├── budget
    │   │   ├── budget.go
    │   │   ├── handler
    │   │   │   ├── adjust.go
    │   │   │   ├── alert.go
    │   │   │   ├── create.go
    │   │   │   ├── get.go
    │   │   │   ├── msg.go
    │   │   │   ├── remove.go
    │   │   │   ├── setup.go     
    │   │   │   └── view.go
    │   ├── init
    │   │   ├── init.go
    │   ├── transaction
    │   │   ├── transaction.go
    │   └── root.go
    ├── db
    │   ├── alert_db
    │   │   └── alert_db.go   
    │   ├── budget_db
    │   │   └── budget_db.go
    │   ├── migrations
    │   │   ├── 001_create_budget_table.sql
    │   │   └── 002_create_alert_table.sql
    │   └── connection.go
    ├── docs
    │   ├── commands.md
    │   └── structure.md
    ├── email
    │   ├── templates
    │   │   └── alert.html 
    │   └── email.go
    ├── internal
    │   ├── app
    │   │   ├── alert
    │   │   │   └── alert.go
    │   │   ├── budget
    │   │   │   └── budget.go
    │   │   └── init
    │   │       └── init.go
    │   ├── middleware
    │   │   └── env.go
    │   └── structs
    │       └── common.go
    ├── tests
    │   └── app
    │       ├── alert
    │       │   └── alert_test.go
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
    └── main.go

## **File Details:**

### Command files

- **cmd/budget/budget.go:** Budget service entry point.
- **cmd/budget/handler/adjust.go:** Handler for adjusting budget.
- **cmd/budget/handler/alert.go:** Handler for alerting.
- **cmd/budget/handler/create.go:** Handler for creating budget.
- **cmd/budget/handler/get.go:** Handler for getting budget.
- **cmd/budget/handler/msg.go:** Handler for showing the alert messages.
- **cmd/budget/handler/remove.go:** Handler for removing budget.
- **cmd/budget/handler/setup.go:** Handler for setting up the alert when the limitation is passed.
- **cmd/budget/handler/view.go:** Handler for viewing budget.
- **cmd/init/init.go:** Flow initialization functionality.
- **cmd/transaction/transaction.go:** Transaction service entry point.

### Database files
- **db/connection.go:** Database connection setup.
- **db/migrations/001_create_budget_table.sql:** SQL script for creating budget table.
- **db/migrations/002_create_alert_table.sql:** SQL script for creating alert table.
- **db/alert_db/alert_db.go:** CRUD operation for the alert functionanlity.
- **db/budget_db/budget_db.go:** CRUD operation for the budget functionanlity.

### Documentation files

- **docs/commands.md:** Commands of the whole application.
- **docs/structure.md:** Structure of the whole application.

### Email files

- **email/email.go:** Handling the email function.
- **email/templates/alert.html:** Email template for alert notification.

### App logic files

- **internal/app/alert/alert.go:** Logic for alert management functionality.
- **internal/app/budget/budget.go:** Logic for budget management functionality.
- **internal/app/init/init.go:** Logic for init functionality.
- **internal/middleware/env.go:** Environment middleware for handling environment variables.
- **tests/app/budget/alert_test.go:** Test file for handling the alert test functions.
- **tests/app/budget/budget_test.go:** Test file for handling the budget test functions.

### Root files

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